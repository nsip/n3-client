package query

import (
	"testing"

	g "github.com/nsip/n3-client/global"
)

func TestQueryMeta(t *testing.T) {
	g.Init()
	defer func() { ph(recover(), g.Cfg.ErrLog) }()

	g.CurCtx = g.Cfg.RPC.CtxList[0] //

	// s, p, o, v := Data(g.CurCtx, "", "xapi ~ verb ~ display ~ en-US", "completed") //   *** n3node thinks it is claiming ticket ***
	s, p, o, v := Data(g.CurCtx, "", g.MARKTerm, "4947ED1F-1E94-4850-8B8F-35C653F51E8C") //   *** n3node thinks it is claiming ticket ***
	for i := range s {
		fPln("----------------------------------------------------")
		fPf("%d # %d: Reply:\n%s\n%s\n%s\n", i, v[i], s[i], p[i], o[i])
	}
	fPln(" **************************** ")

	s, p, o, v = Data(g.CurCtx, "644153cf-02c2-4670-810b-534ac148c011", g.MARKTerm)
	for i := range s {
		fPln("----------------------------------------------------")
		fPf("%d # %d: Reply:\n%s\n%s\n%s\n", i, v[i], s[i], p[i], o[i])
	}
	fPln()
}

func TestQuery1(t *testing.T) {
	g.Init()
	defer func() { ph(recover(), g.Cfg.ErrLog) }()

	// lCtxList := len(g.Cfg.RPC.CtxList)
	g.CurCtx = g.Cfg.RPC.CtxList[0] //
	ctx := g.CurCtx

	ObjectID := "5e8d383b-13d2-481f-8db5-c16376279566"
	Object := ""
	s, p, o, _ := Data(ctx, ObjectID, "") //      ** root **
	fPln("Object:")
	for i := range s {
		Object = o[i]
		fPf("%-40s%-10s%s\n", s[i], p[i], o[i])
	}

	// ************************* //

	fPln("\nArray:")
	s, p, o, _ = Data(ctx, ObjectID, "[]") //     ** array **
	for i := range s {
		fPf("%-40s%-70s%-10s\n", s[i], p[i], o[i])
	}

	fPln("\nStructure:")
	s, p, o, _ = Data(ctx, ObjectID, "::") //     ** struct **
	for i := range s {
		fPf("%-40s%-60s%s\n", s[i], p[i], o[i])
	}

	fPln("\nValues:")
	s, p, o, _ = Data(ctx, ObjectID, Object) //   ** values **
	for i := range s {
		fPf("%-40s%-85s%s\n", s[i], p[i], o[i])
	}

	// ************************* //

	fPln("\nContext-prictrlID:")
	s, p, o, _ = Data("ctxid", "xapi", "comment 2") //   ** get prictrl ID **
	for i := range s {
		fPf("%-40s%-85s%s\n", s[i], p[i], o[i])
	}

}

func TestQuery2(t *testing.T) {
	g.Init()
	defer func() { ph(recover(), g.Cfg.ErrLog) }()

	ctx := g.CurCtx

	s1, _, _, _ := Data(ctx, "", "xapi ~ actor ~ name", "Samuel Busse")
	if len(s1) == 0 {
		fPln("nothing found 1")
		return
	}
	s2, _, _, _ := Data(ctx, "", "xapi ~ object ~ id", "http://example.com/assignments/Geography-8-1-B:4")
	if len(s2) == 0 {
		fPln("nothing found 2")
		return
	}

	if rst := IArrIntersect(Ss(s1), Ss(s2)); rst != nil {

		ids := rst.([]string)
		for _, id := range ids {
			fPln(id)
		}

		fPln(" ----------------------------------------- ")

		ObjectID := ids[0] // "478CE5FA-0BCE-424C-9F11-E57A15E941CF"
		Object := ""
		s, p, o, _ := Data(ctx, ObjectID, "") //     ** root **
		fPln("Object:")
		for i := range s {
			Object = o[i]
			fPf("%-50s%-10s%s\n", s[i], p[i], o[i])
		}

		fPln("\nArray Info:")
		s, p, o, _ = Data(ctx, ObjectID, "[]") //    ** array ** /* context must end with '-xapi' */
		for i := range s {
			fPf("%-50s%-50s%s\n", s[i], p[i], o[i])
		}

		fPln("\nStructure:")
		s, p, o, _ = Data(ctx, Object, "::") //      ** struct **
		for i := range s {
			fPf("%-50s%-10s%s\n", s[i], p[i], o[i])
		}

		fPln("\nValues:")
		s, p, o, _ = Data(ctx, ObjectID, Object) //  ** values **
		for i := range s {
			fPf("%-50s%-50s%s\n", s[i], p[i], o[i])
		}

	} else {
		fPln("nothing found 1 & 2")
		return
	}
}
