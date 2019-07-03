package query

import (
	"testing"

	c "../config"
	g "../global"
)

func TestN3LoadConfig(t *testing.T) {
	InitClient(c.FromFile("./config.toml", "../config/config.toml"))
}

func TestQueryMeta(t *testing.T) {
	defer func() { ph(recover(), g.Cfg.ErrLog) }()
	TestN3LoadConfig(t)

	ctx := g.CurCtx

	s, p, o, v := Meta(ctx, "738F4DF5-949F-4380-8186-8252440A6F6F", "V") //         *** n3node thinks it is claiming ticket ***
	for i := range s {
		fPln("----------------------------------------------------")
		fPf("%d # %d: Reply: %s\n%s\n%s \n", i, v[i], s[i], p[i], o[i])
	}
	fPln()
}

func TestQuery1(t *testing.T) {
	defer func() { ph(recover(), g.Cfg.ErrLog) }()
	TestN3LoadConfig(t)

	ctx := g.CurCtx

	// ObjectID := "40efdee5-df52-4cd7-b70c-d457629ae32d"
	ObjectID := "5758757f-b34b-4385-9d14-0290bc9afd07"
	Object := ""
	s, p, o, _ := Data(ctx, ObjectID, "") //      ** root **
	fPln("Object:")
	for i := range s {
		Object = o[i]
		fPf("%-40s%-10s%s\n", s[i], p[i], o[i])
	}

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
}

func TestQuery2(t *testing.T) {
	defer func() { ph(recover(), g.Cfg.ErrLog) }()
	TestN3LoadConfig(t)

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

	if rst := IArrIntersect(Strs(s1), Strs(s2)); rst != nil {

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
