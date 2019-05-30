package rest

import (
	"io/ioutil"
	"net/http"

	c "../config"
	g "../global"
	"../gql"
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

// SendToNode : Send XAPI / SIF to N3-Transport
func SendToNode(c echo.Context) error {
	defer func() {
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	body := Must(ioutil.ReadAll(c.Request().Body)).([]byte)
	data := string(body)
	IDs, nV, nS, nA := send.ToNode(data, "id", "xapi")
	for _, id := range IDs {
		g.LCSchema.Remove(id)
		g.LCJSON.Remove(id)
	}

	return c.JSON(http.StatusAccepted, fSf("<%d> v-tuples, <%d> s-tuples, <%d> a-tuples have been sent", nV, nS, nA))
}

// GQLRequest : wrapper type to capture GQL input
type GQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func queryGQL(c echo.Context) error {
	defer func() {
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	// mPP := map[string]string{
	// 	"fname":           "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ Name ~ FamilyName",
	// 	"gname":           "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ Name ~ GivenName",
	// 	"staffid":         "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ StaffPersonalRefId",
	// 	"teachinggroupid": "TeachingGroup ~ -RefId",
	// 	"tgid":            "GradingAssignment ~ TeachingGroupRefId",
	// 	"studentid":       "StudentAttendanceTimeList ~ StudentPersonalRefId",
	// 	"objid":           "xapi ~ object ~ id",
	// }
	mPV := map[string]interface{}{}

	// ********************* POSTMAN client *********************
	// fname, gname := c.QueryParam("fname"), c.QueryParam("gname")
	// queryTxt := string(Must(ioutil.ReadAll(c.Request().Body)).([]byte))

	// ********************* GRAPHIQL client *********************
	gqlreq := new(GQLRequest) //
	PE(c.Bind(gqlreq))        //                            *** ONLY <POST> echo can Bind OK ***
	queryTxt := gqlreq.Query
	for k, v := range gqlreq.Variables {
		mPV[k] = v.(string)
	}

	IDs, rmStructs := []string{}, []string{}
	IDs = append(IDs, "ca669951-9511-4e53-ae92-50845d3bdcd6") // *** if param is hard-coded here, GraphiQL can show Schema-Doc ***
	// if id, ok := mPV["objid"]; ok { //                              *** if param is given runtime, GraphiQL cannot show Schema-Doc ***
	// 	IDs = append(IDs, id.(string))
	// }

	// switch {
	// case sCtn(queryTxt, "TeachingGroupByName(") || sCtn(queryTxt, "TeachingGroupByStaffID("):
	// 	IDs, rmStructs = IDsByPO(mPP, mPV), []string{"StudentList"}
	// case sCtn(queryTxt, "TeachingGroup("):
	// 	IDs = IDsByPO(mPP, mPV)
	// case sCtn(queryTxt, "GradingAssignment("):
	// 	IDs = IDsByPO(mPP, mPV)
	// case sCtn(queryTxt, "StudentAttendance("):
	// 	IDs = IDsByPO(mPP, mPV)
	// case sCtn(queryTxt, "QueryXAPI("):
	// 	IDs = IDsByPO(mPP, mPV)
	// }

	qSchema := string(Must(ioutil.ReadFile("./gql/qSchema/qRoot.gql")).([]byte)) //  *** content should be related to resolver path ***

	if len(IDs) >= 1 {
		rst := gql.Query(IDs, qSchema, queryTxt, mPV, rmStructs) //*** rst is already JSON string, so use String to return ***
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
	e.POST(CFG.Rest.PathSend, SendToNode)
	e.POST(CFG.Rest.PathGQL, queryGQL)

	// Server
	e.Start(fSf(":%d", CFG.Rest.Port))
}
