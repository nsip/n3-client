package rest

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

	c "../config"
	d "../delete"
	g "../global"
	"../gql"
	q "../query"
	"../send"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// **************************** controllers **************************** //

// OriExePathChk :
func OriExePathChk() {
AGAIN:
	if path, _ := os.Getwd(); path != g.OriExePath {
		time.Sleep(100 * time.Millisecond)
		goto AGAIN
	}
}

// InitClient :
func InitClient(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	CFG = config
}

// getIDList :
func getIDList(c echo.Context) error {
	defer func() {
		mtxID.Unlock()
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxID.Lock()

	params := c.QueryParams()
	mPP, mPV, mKV := map[string]string{}, map[string]interface{}{}, map[string]string{}
	if object, ok := params["object"]; ok {
		files := Must(ioutil.ReadDir(CFG.Query.ParamPathDir)).([]os.FileInfo)
		for _, f := range files {
			if f.Name() == object[0] {
				data := string(Must(ioutil.ReadFile(CFG.Query.ParamPathDir + f.Name())).([]byte))
				mKV = Str(data).KeyValueMap('\n', ':', '#')
				break
			}
		}
		for k, v := range params {
			if _, ok := mKV[k]; ok {
				mPP[k], mPV[k] = mKV[k], Str(v[0]).T(BLANK).V()
			}
		}

		if all, ok := params["all"]; ok && all[0] == "true" {
			return c.JSON(http.StatusAccepted, IDsByPO(mPP, mPV, true))
		}
		return c.JSON(http.StatusAccepted, IDsByPO(mPP, mPV, false))
	}
	return c.JSON(http.StatusBadRequest, "<object> must be provided")
}

// delFromNode :
func delFromNode(c echo.Context) error {
	defer func() {
		mtxDel.Unlock()
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxDel.Lock()

	idList := c.QueryParams()["id"]
	d.DelBat(idList...)
	g.RmIDsInLRU(idList...)
	return c.JSON(http.StatusAccepted, fSf("%d objects have been deleted", len(idList)))
}

// postToNode : Publish Data to N3-Transport
func postToNode(c echo.Context) error {
	defer func() {
		mtxPub.Unlock()
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxPub.Lock()

	idmark, dfltRoot := c.QueryParam("idmark"), c.QueryParam("dfltRoot")
	// fPln(idmark, ":", dfltRoot)
	if idmark == "" {
		return c.JSON(http.StatusBadRequest, "<idmark> must be provided")
	}
	if dfltRoot == "" {
		return c.JSON(http.StatusBadRequest, "<dfltRoot> must be provided")
	}

	data := string(Must(ioutil.ReadAll(c.Request().Body)).([]byte))
	if data == "" {
		return c.JSON(http.StatusBadRequest, "Nothing to be sent as POST BODY is empty")
	}

	IDs, nV, nS, nA := send.Pub2Node(data, idmark, dfltRoot)
	g.RmIDsInLRU(IDs...)
	return c.JSON(http.StatusAccepted, fSf("<%d> v-tuples, <%d> s-tuples, <%d> a-tuples have been sent", nV, nS, nA))
}

// Request : wrapper type to capture GQL input
type Request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// postQueryGQL :
func postQueryGQL(c echo.Context) error {
	defer func() {
		mtxQry.Unlock()
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxQry.Lock()

	// ********************* POSTMAN client ********************* //
	// fname, gname := c.QueryParam("fname"), c.QueryParam("gname")
	// qTxt := string(Must(ioutil.ReadAll(c.Request().Body)).([]byte))

	fPln(":::queryGQL:::", c.QueryParam("objid"))

	// ********************* GRAPHIQL client ********************* //
	req := new(Request) //
	PE(c.Bind(req))     //                   *** ONLY <POST> echo can Bind OK ***
	qTxt := req.Query
	mPV := map[string]interface{}{}
	for k, v := range req.Variables {
		mPV[k] = v.(string)
	}

	IDs, rmStructs, root := []string{}, []string{}, ""
	mReplace := map[string]string{
		"en-US": "en_US",
	}

	// IDs = append(IDs, "ca669951-9511-4e53-ae92-50845d3bdcd6") // *** if param is hard-coded here, GraphiQL can show Schema-Doc ***
	if id, ok := mPV["objid"]; ok { //                              *** if param is given at runtime, GraphiQL cannot show Schema-Doc ***
		IDs = append(IDs, id.(string))
		_, _, o, _ := q.Data(id.(string), "")
		if len(o) > 0 && o[0] != "" {
			root = o[0]
		} else {
			return c.JSON(http.StatusAccepted, "id provided is not in db")
		}
	} else {
		return c.JSON(http.StatusAccepted, "<objid> is missing")
	}

	qSchema := string(Must(ioutil.ReadFile(CFG.Query.SchemaDir + root + ".gql")).([]byte)) //  *** content must be related to resolver path ***
	if len(IDs) >= 1 {
		rst := gql.Query(IDs, qSchema, CFG.Query.SchemaDir, qTxt, mPV, rmStructs, mReplace) // *** rst is already JSON string, so use String to return ***
		rst = ASCIIToOri(rst)                                                               // *** ascii rst back to original rst ***
		return c.String(http.StatusAccepted, rst)
	}

	return c.JSON(http.StatusAccepted, "Nothing Found")
}

// getObject :
func getObject(c echo.Context) error {
	defer func() {
		mtxObj.Unlock()
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxObj.Lock()

	id := c.QueryParam("id")
	return c.JSON(http.StatusAccepted, id)
}

// getSchema :
func getSchema(c echo.Context) error {
	defer func() {
		mtxScm.Unlock()
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxScm.Lock()

	id := c.QueryParam("id")
	return c.JSON(http.StatusAccepted, id)
}

// getQueryText :
func getQueryText(c echo.Context) error {
	defer func() {
		mtxQryTxt.Unlock()
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxQryTxt.Lock()

	id := c.QueryParam("id")
	return c.JSON(http.StatusAccepted, id)
}

// ************************************************ HOST ************************************************ //

// HostHTTPAsync : Host a HTTP Server for publishing xml json string(request body) to <n3-transport> grpc Server
func HostHTTPAsync() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE},
	}))

	// Route
	e.GET(CFG.Rest.PathTest, func(c echo.Context) error { return c.String(http.StatusOK, "n3client is running\n") })
	e.GET(CFG.Rest.PathID, getIDList)
	e.GET(CFG.Rest.PathObj, getObject)
	e.GET(CFG.Rest.PathScm, getSchema)
	e.GET(CFG.Rest.PathGQLTxt, getQueryText)
	e.POST(CFG.Rest.PathPub, postToNode)
	e.POST(CFG.Rest.PathGQL, postQueryGQL)
	e.DELETE(CFG.Rest.PathDel, delFromNode)

	// Server
	e.Start(fSf(":%d", CFG.Rest.Port))
}
