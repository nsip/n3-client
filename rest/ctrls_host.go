package rest

import (
	"fmt"
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

func cdUL(dir string) string {
	return S(dir).RmTailFromLast("/").V()
}

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
	defer func() { mtxID.Unlock() }()
	OriExePathChk()
	mtxID.Lock()

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

// delFromNode : this func can only delete normal data. IF delete privacy control, use cli-privacy
func delFromNode(c echo.Context) error {
	defer func() { mtxDel.Unlock() }()
	OriExePathChk()
	mtxDel.Lock()

	IDs := c.QueryParams()["id"]
	d.DelBat(g.CurCtx, IDs...)
	return c.JSON(http.StatusAccepted, fSf("%d objects have been deleted", len(IDs)))
}

// postToNode : Publish Data to N3-Transport
func postToNode(c echo.Context) error {
	defer func() { mtxPub.Unlock() }()
	OriExePathChk()
	mtxPub.Lock()

	root := c.QueryParam("dfltRoot")
	// fPln(dfltRoot)
	if root == "" {
		return c.String(http.StatusBadRequest, "<dfltRoot> must be provided")
	}

	data := string(must(ioutil.ReadAll(c.Request().Body)).([]byte))
	if data == "" {
		return c.String(http.StatusBadRequest, "Nothing to be sent as BODY is empty")
	}

	if _, _, nV, nS, nA, e := pub.Pub2Node(g.CurCtx, data, root); e != nil { //    *** preprocess, postprocess included ***
		return e
	} else {
		return c.JSON(http.StatusAccepted, fSf("<%d> v-tuples, <%d> s-tuples, <%d> a-tuples have been sent", nV, nS, nA))
	}
}

// postFileToNode :
func postFileToNode(c echo.Context) error {
	defer func() { mtxPub.Unlock() }()
	OriExePathChk()
	mtxPub.Lock()

	name, pwd, root := c.FormValue("username"), c.FormValue("password"), c.FormValue("root")
	fPln(name, pwd, root)

	if g.CurCtx = ctxFromCredential(name, pwd); g.CurCtx == "" {
		return c.String(http.StatusUnauthorized, "wrong username or password")
	}

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	buffer := make([]byte, file.Size)
	src.Read(buffer)
	data := string(buffer)
	if !IsJSON(data) {
		ioutil.WriteFile("not acceptable file.txt", buffer, 0666)
		return c.String(http.StatusBadRequest, "NOT JSON, CANNOT SEND")
	}

	if _, _, _, _, _, e := pub.Pub2Node(g.CurCtx, data, root); e != nil { //             *** preprocess, postprocess included ***
		return e
	}

	return c.String(http.StatusOK, fmt.Sprintf("%s uploaded successfully", file.Filename))
}

// Request : wrapper type to capture GQL input
type Request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// postQueryGQL :
func postQueryGQL(c echo.Context) error {
	defer func() { mtxQry.Unlock() }()

	OriExePathChk()
	mtxQry.Lock()

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
	defer func() { mtxObj.Unlock() }()

	OriExePathChk()
	mtxObj.Lock()

	id := c.QueryParam("id")
	return c.JSON(http.StatusAccepted, id)
}

// getSchema :
func getSchema(c echo.Context) error {
	defer func() { mtxScm.Unlock() }()

	OriExePathChk()
	mtxScm.Lock()

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
	e.Use(middleware.BodyLimit("2G"))

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	webloc := g.Cfg.Group.APP + g.Cfg.Route.Pub
	e.File(webloc, "../www/service.html")
	e.Static(cdUL(webloc), "../www/") //             "/" is html - ele - <src>'s path

	// Maybe Auth middleware dislike long request body ? manually check
	e.POST(g.Cfg.Route.FilePub, postFileToNode)

	// Group
	api := e.Group(g.Cfg.Group.API)
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	api.Use(middleware.BodyLimit("2G"))

	uname := ""
	// BasicAuth ( Big Body has ERR_CONNECTION_RESET in this )
	api.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		fPln("---------------------in basicAuth-----------------------------", username, password)
		if g.CurCtx = ctxFromCredential(username, password); g.CurCtx == "" {
			return false, c.String(http.StatusUnauthorized, "wrong username or password")
		}
		uname = username
		return true, nil
	}))

	// api Route
	// api.GET("/filetest", func(c echo.Context) error { return c.File("/home/qing/Desktop/index.html") })
	api.GET(g.Cfg.Route.Greeting, func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, "+uname+". n3client is running @ "+time.Now().Format("2006-01-02 15:04:05.000"))
	})
	api.GET(g.Cfg.Route.ID, getIDList)
	api.GET(g.Cfg.Route.Obj, getObject)
	api.GET(g.Cfg.Route.Scm, getSchema)
	api.POST(g.Cfg.Route.Pub, postToNode)
	api.POST(g.Cfg.Route.GQL, postQueryGQL)
	api.DELETE(g.Cfg.Route.Del, delFromNode)

	// Server
	e.Start(fSf(":%d", g.Cfg.WebService.Port))
}

func ctxFromCredential(uname, pwd string) string {
	switch {
	case uname == "admin" && pwd == "admin":
		return g.Cfg.RPC.CtxPrivDef
	case uname == "user" && pwd == "user":
		return g.Cfg.RPC.CtxList[0]
	case uname == "user1" && pwd == "user1":
		return g.Cfg.RPC.CtxList[1]
	default:
		return ""
	}
}
