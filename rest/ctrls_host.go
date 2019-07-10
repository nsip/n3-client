package rest

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

	d "../delete"
	g "../global"
	"../gql"
	pub "../publish"
	q "../query"
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

// getIDList :
func getIDList(c echo.Context) error {
	defer func() {
		mtxID.Unlock()
		phe(recover(), g.Cfg.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxID.Lock()

	g.CurCtx = g.Cfg.RPC.CtxList[0]

	params := c.QueryParams()
	if object, ok := params["object"]; ok { //                             *** object Value only indicates the file to get parampath ***
		mPP, mPV := map[string]string{}, map[string]interface{}{}
		ppDir, foundPPFile := g.Cfg.Query.ParamPathDir, false
		for _, f := range must(ioutil.ReadDir(ppDir)).([]os.FileInfo) {
			if f.Name() == object[0] {
				data := string(must(ioutil.ReadFile(ppDir + f.Name())).([]byte))
				mPP = S(data).KeyValueMap('\n', ':', '#')
				foundPPFile = true
				break
			}
		}
		if !foundPPFile {
			return c.JSON(http.StatusForbidden, "<object>'s params-path file was not found, contact n3-client admin to solve it")
		}
		for k, v := range params {
			if _, ok := mPP[k]; ok {
				mPV[k] = S(v[0]).T(BLANK).V()
			}
		}
		all, ok := params["all"]
		getall := IF(ok && all[0] == "true", true, false).(bool)
		return c.JSON(http.StatusAccepted, GetIDs(g.CurCtx, mPP, mPV, object[0], getall))
	}
	return c.JSON(http.StatusBadRequest, "<object> must be provided")
}

// delFromNode :
func delFromNode(c echo.Context) error {
	defer func() {
		mtxDel.Unlock()
		phe(recover(), g.Cfg.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxDel.Lock()

	g.CurCtx = g.Cfg.RPC.CtxList[0]

	idList := c.QueryParams()["id"]
	d.DelBat(g.CurCtx, idList...)
	g.RmIDsInLRU(idList...)
	g.RmQryIDsCache(idList...)
	return c.JSON(http.StatusAccepted, fSf("%d objects have been deleted", len(idList)))
}

// postToNode : Publish Data to N3-Transport
func postToNode(c echo.Context) error {
	defer func() {
		mtxPub.Unlock()
		phe(recover(), g.Cfg.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxPub.Lock()

	g.CurCtx = g.Cfg.RPC.CtxList[0]

	idmark, dfltRoot := c.QueryParam("idmark"), c.QueryParam("dfltRoot")
	// fPln(idmark, ":", dfltRoot)
	if idmark == "" {
		return c.JSON(http.StatusBadRequest, "<idmark> must be provided")
	}
	if dfltRoot == "" {
		return c.JSON(http.StatusBadRequest, "<dfltRoot> must be provided")
	}

	data := string(must(ioutil.ReadAll(c.Request().Body)).([]byte))
	if data == "" {
		return c.JSON(http.StatusBadRequest, "Nothing to be sent as POST BODY is empty")
	}

	_, _, nV, nS, nA := pub.Pub2Node(g.CurCtx, data, idmark, dfltRoot) //             *** preprocess, postprocess included ***
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
		phe(recover(), g.Cfg.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxQry.Lock()

	g.CurCtx = g.Cfg.RPC.CtxList[0]

	// ********************* POSTMAN client ********************* //
	// fname, gname := c.QueryParam("fname"), c.QueryParam("gname")
	// qTxt := string(Must(ioutil.ReadAll(c.Request().Body)).([]byte))

	fPln(":::queryGQL:::", c.QueryParam("objid"))

	// ********************* GRAPHIQL client ********************* //
	req := new(Request) //
	pe(c.Bind(req))     //          *** ONLY <POST> echo can Bind OK ***
	mPV := map[string]interface{}{}
	for k, v := range req.Variables {
		mPV[k] = v.(string)
	}

	IDs := []string{}
	// IDs = append(IDs, "ca669951-9511-4e53-ae92-50845d3bdcd6") // *** if param is hard-coded here, GraphiQL can show Schema-Doc ***
	if id, ok := mPV["objid"]; ok { //                              *** if param is given at runtime, GraphiQL cannot show Schema-Doc ***
		IDs = append(IDs, id.(string))
		if _, _, o, _ := q.Data(g.CurCtx, id.(string), ""); len(o) == 0 || o[0] == "" {
			return c.JSON(http.StatusAccepted, "id provided is not in db")
		}
	} else {
		return c.JSON(http.StatusAccepted, "<objid> is missing")
	}

	if len(IDs) >= 1 {
		gqlrst := gql.Query(g.CurCtx, IDs, g.Cfg.Query.SchemaDir, req.Query, mPV, g.MpQryRstRplc) // *** gqlrst is already JSON string, use String to return ***
		return c.String(http.StatusAccepted, gqlrst)
	}

	return c.JSON(http.StatusAccepted, "Nothing Found")
}

// getObject :
func getObject(c echo.Context) error {
	defer func() {
		mtxObj.Unlock()
		phe(recover(), g.Cfg.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxObj.Lock()

	g.CurCtx = g.Cfg.RPC.CtxList[0]

	id := c.QueryParam("id")
	return c.JSON(http.StatusAccepted, id)
}

// getSchema :
func getSchema(c echo.Context) error {
	defer func() {
		mtxScm.Unlock()
		phe(recover(), g.Cfg.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	OriExePathChk()
	mtxScm.Lock()

	g.CurCtx = g.Cfg.RPC.CtxList[0]

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
	e.GET(g.Cfg.Rest.PathTest, func(c echo.Context) error { return c.String(http.StatusOK, "n3client is running\n") })
	e.GET(g.Cfg.Rest.PathID, getIDList)
	e.GET(g.Cfg.Rest.PathObj, getObject)
	e.GET(g.Cfg.Rest.PathScm, getSchema)
	e.POST(g.Cfg.Rest.PathPub, postToNode)
	e.POST(g.Cfg.Rest.PathGQL, postQueryGQL)
	e.DELETE(g.Cfg.Rest.PathDel, delFromNode)

	// Server
	e.Start(fSf(":%d", g.Cfg.Rest.Port))
}
