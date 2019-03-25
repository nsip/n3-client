package query

import (
	"testing"

	c "../config"
	g "../global"
	u "github.com/cdutwhu/go-util"
)

func TestN3LoadConfig(t *testing.T) {
	InitFrom(c.FromFile("./config.toml", "../config/config.toml"))
}

func TestQueryMetaSif(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)

	s, p, o, v := Meta(g.SIF, "D3E34F41-9D75-101A-8C3D-00AA001A1652", "V")
	for i := range s {
		fPln("----------------------------------------------------")
		fPf("%d # %d: Reply: %s\n%s\n%s \n", i, v[i], s[i], p[i], o[i])
	}
	fPln()
}

func TestQuerySif(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)

	Fn := Sif

	ObjectID := "1BC3EAB7-3E48-4371-8C14-6D1E67BEBD6D"
	Object := ""
	s, p, o, _ := Fn(ObjectID, "") //      ** root **
	fPln("Object:")
	for i := range s {
		Object = o[i]
		fPf("%-50s%-10s%s\n", s[i], p[i], o[i])
	}

	fPln("\nArray Info:")
	s, p, o, _ = Fn(ObjectID, "[]") //     ** array **
	for i := range s {
		fPf("%-40s%-65s%10s\n", s[i], p[i], o[i])
	}

	fPln("\nStructure:")
	s, p, o, _ = Fn(Object, "::") //       ** struct **
	for i := range s {
		fPf("%-65s%-10s%s\n", s[i], p[i], o[i])
	}

	fPln("\nValues:")
	s, p, o, _ = Fn(ObjectID, Object) //   ** values **
	for i := range s {
		fPf("%-40s%-85s%s\n", s[i], p[i], o[i])
	}
}

func TestQueryXapi(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)

	Fn := Xapi

	ids1, ids2 := []string{}, []string{}
	s, p, o, _ := Fn("", "XAPI ~ actor ~ name", "Brittany Baxter")
	for i := range s {
		ids1 = append(ids1, s[i])
	}
	s, p, o, _ = Fn("", "XAPI ~ object ~ id", "http://example.com/assignments/History-7-1-B:5")
	for i := range s {
		ids2 = append(ids2, s[i])
	}
	ids := u.Strs(ids1).ToG().InterSec(u.Strs(ids2).ToG()...)
	for _, id := range ids {
		fPln(id)
	}

	if len(ids) == 0 {
		return
	}

	fPln(" ----------------------------------------- ")

	ObjectID := ids[0].(string) //"478CE5FA-0BCE-424C-9F11-E57A15E941CF"
	Object := ""
	s, p, o, _ = Fn(ObjectID, "") //     ** root **
	fPln("Object:")
	for i := range s {
		Object = o[i]
		fPf("%-50s%-10s%s\n", s[i], p[i], o[i])
	}

	fPln("\nArray Info:")
	s, p, o, _ = Fn(ObjectID, "[]") //    ** array ** /* context must end with '-xapi' */
	for i := range s {
		fPf("%-50s%-50s%s\n", s[i], p[i], o[i])
	}

	fPln("\nStructure:")
	s, p, o, _ = Fn(Object, "::") //      ** struct **
	for i := range s {
		fPf("%-50s%-10s%s\n", s[i], p[i], o[i])
	}

	fPln("\nValues:")
	s, p, o, _ = Fn(ObjectID, Object) //  ** values **
	for i := range s {
		fPf("%-50s%-50s%s\n", s[i], p[i], o[i])
	}
}
