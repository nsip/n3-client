package rest

import (
	"io/ioutil"
	"net/http"

	c "../config"
	g "../global"
	"../gql"
	q "../query"
	"../send"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// **************************** controllers **************************** //

// InitClient :
func InitClient(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	CFG = config
}

// getIDList :
func getIDList(c echo.Context) error {
	defer func() {
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	object := c.QueryParam("object")
	if object == "" {
		return c.JSON(http.StatusBadRequest, "<object> must be provided")
	}

	mPP := map[string]string{}
	mPV := map[string]interface{}{}
	switch object {
	case "StaffPersonal":
		{
			mPP["fname"] = "StaffPersonal ~ PersonInfo ~ Name ~ FamilyName"
			mPP["gname"] = "StaffPersonal ~ PersonInfo ~ Name ~ GivenName"
			mPV["fname"] = c.QueryParam("fname")
			mPV["gname"] = c.QueryParam("gname")
		}
	case "TeachingGroup":
		{
			mPP["fname"] = "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ Name ~ FamilyName"
			mPP["gname"] = "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ Name ~ GivenName"
			mPV["fname"] = c.QueryParam("fname")
			mPV["gname"] = c.QueryParam("gname")
		}
	default:
		return c.JSON(http.StatusBadRequest, fSf("<%s>'s id query is not implemented", object))
	}

	return c.JSON(http.StatusAccepted, IDsByPO(mPP, mPV))
}

// sendToNode : Send Data to N3-Transport
func sendToNode(c echo.Context) error {
	defer func() {
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

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

	IDs, nV, nS, nA := send.ToNode(data, idmark, dfltRoot)
	for _, id := range IDs {
		g.LCSchema.Remove(id)
		g.LCJSON.Remove(id)
	}
	return c.JSON(http.StatusAccepted, fSf("<%d> v-tuples, <%d> s-tuples, <%d> a-tuples have been sent", nV, nS, nA))
}

// Request : wrapper type to capture GQL input
type Request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func queryGQL(c echo.Context) error {
	defer func() {
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	// ********************* POSTMAN client *********************
	// fname, gname := c.QueryParam("fname"), c.QueryParam("gname")
	// qTxt := string(Must(ioutil.ReadAll(c.Request().Body)).([]byte))

	// ********************* GRAPHIQL client *********************
	req := new(Request) //
	PE(c.Bind(req))     //                            *** ONLY <POST> echo can Bind OK ***
	qTxt := req.Query
	mPV := map[string]interface{}{}
	for k, v := range req.Variables {
		mPV[k] = v.(string)
	}

	IDs, rmStructs, root, mReplace := []string{}, []string{}, "", map[string]string{"en-US": "en_US"}
	// IDs = append(IDs, "ca669951-9511-4e53-ae92-50845d3bdcd6") // *** if param is hard-coded here, GraphiQL can show Schema-Doc ***
	if id, ok := mPV["objid"]; ok { //                              *** if param is given at runtime, GraphiQL cannot show Schema-Doc ***
		IDs = append(IDs, id.(string))
		_, _, o, _ := q.Data(id.(string), "")
		if len(o) > 0 {
			root = o[0]
		} else {
			return c.JSON(http.StatusAccepted, "id provided is not in db")
		}
	} else {
		return c.JSON(http.StatusAccepted, "<objid> is missing")
	}

	// switch {
	// case sCtn(qTxt, "TeachingGroupByName(") || sCtn(qTxt, "TeachingGroupByStaffID("):
	// 	IDs, rmStructs = IDsByPO(mPP, mPV), []string{"StudentList"}
	// case sCtn(qTxt, "TeachingGroup("):
	// 	IDs = IDsByPO(mPP, mPV)
	// case sCtn(qTxt, "GradingAssignment("):
	// 	IDs = IDsByPO(mPP, mPV)
	// case sCtn(qTxt, "StudentAttendance("):
	// 	IDs = IDsByPO(mPP, mPV)
	// case sCtn(qTxt, "QueryXAPI("):
	// 	IDs = IDsByPO(mPP, mPV)
	// }

	qSchemaDir := "./gql/qSchema/"
	qSchema := string(Must(ioutil.ReadFile(qSchemaDir + root + ".gql")).([]byte)) //  *** content should be related to resolver path ***

	if len(IDs) >= 1 {
		rst := gql.Query(IDs, qSchema, qSchemaDir, qTxt, mPV, rmStructs, mReplace) // *** rst is already JSON string, so use String to return ***
		return c.String(http.StatusAccepted, rst)
	}

	return c.JSON(http.StatusAccepted, "Nothing Found")
}

// ************************************************ HOST ************************************************ //

// HostHTTPAsync : Host a HTTP Server for publishing SIF(xml) XAPI(json) string(request body) to <n3-transport> grpc Server
func HostHTTPAsync() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	// Route
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "n3client is running\n") })
	e.POST(CFG.Rest.PathSend, sendToNode)
	e.POST(CFG.Rest.PathGQL, queryGQL)
	e.GET("/id/", getIDList)

	// Server
	e.Start(fSf(":%d", CFG.Rest.Port))
}
