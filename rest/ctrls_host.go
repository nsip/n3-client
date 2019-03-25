package rest

import (
	"io/ioutil"
	"net/http"

	c "../config"
	"../gql"
	"../send"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// **************************** controllers **************************** //

// InitFrom :
func InitFrom(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	CFG = config
}

func sendSIF(c echo.Context) error {
	defer func() {
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	body := Must(ioutil.ReadAll(c.Request().Body)).([]byte)
	nV, nS, nA, termID := send.Sif(string(body))
	return c.JSON(http.StatusAccepted, fSf("<%d> v-tuples, <%d> s-tuples, <%d> a-tuples have been sent, @ %s", nV, nS, nA, termID))
}

func sendXAPI(c echo.Context) error {
	defer func() {
		PHE(recover(), CFG.Global.ErrLog, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	body := Must(ioutil.ReadAll(c.Request().Body)).([]byte)
	cntV, cntS, cntA, termID := send.Xapi(string(body))
	return c.JSON(http.StatusAccepted, fSf("%d tuples has been sent, @ %s", cntV+cntS+cntA, termID))
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

	mapPP := map[string]string{
		"fname":           "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ Name ~ FamilyName",
		"gname":           "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ Name ~ GivenName",
		"staffid":         "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ StaffPersonalRefId",
		"teachinggroupid": "TeachingGroup ~ -RefId",
		"tgid":            "GradingAssignment ~ TeachingGroupRefId",
		"studentid":       "StudentAttendanceTimeList ~ StudentPersonalRefId",
		"objid":           "XAPI ~ object ~ id",
	}
	mapPV := map[string]interface{}{}

	// *** postman client ***
	// fname, gname := c.QueryParam("fname"), c.QueryParam("gname")
	// queryTxt := string(Must(ioutil.ReadAll(c.Request().Body)).([]byte))

	// *** graphiql client ***
	gqlreq := new(GQLRequest) //
	PE(c.Bind(gqlreq))        //                            *** ONLY <POST> echo can Bind OK ***
	queryTxt := gqlreq.Query
	for k, v := range gqlreq.Variables {
		mapPV[k] = v.(string)
	}

	schemaQuery := string(Must(ioutil.ReadFile("./gql/query.gql")).([]byte))
	IDs, fType, rmStructs := []string{}, "", []string{}

	switch {
	case sCtn(queryTxt, "TeachingGroupByName(") || sCtn(queryTxt, "TeachingGroupByStaffID("):
		IDs, fType, rmStructs = IDsByPOFromSIF(mapPP, mapPV), "sif", []string{"StudentList"}
	case sCtn(queryTxt, "TeachingGroup("):
		IDs, fType = IDsByPOFromSIF(mapPP, mapPV), "sif"
	case sCtn(queryTxt, "GradingAssignment("):
		IDs, fType = IDsByPOFromSIF(mapPP, mapPV), "sif"
	case sCtn(queryTxt, "StudentAttendance("):
		IDs, fType = IDsByPOFromSIF(mapPP, mapPV), "sif"
	case sCtn(queryTxt, "QueryXAPI("):
		IDs, fType = IDsByPOFromXAPI(mapPP, mapPV), "xapi"
	}

	if len(IDs) >= 1 {
		rst := gql.GQuery(IDs, fType, schemaQuery, queryTxt, mapPV, rmStructs) //*** rst is already JSON string, so use String to return ***
		return c.String(http.StatusAccepted, rst)
	}

	return c.JSON(http.StatusAccepted, "Nothing Found")
}

// **************************** HOST **************************** //

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
	e.POST(CFG.Rest.SifPathS, sendSIF)
	e.POST(CFG.Rest.XapiPathS, sendXAPI)
	e.POST(CFG.Rest.PathGQL, queryGQL)

	// Server
	e.Start(fSf(":%d", CFG.Rest.Port))
}
