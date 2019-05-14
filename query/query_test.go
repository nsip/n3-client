package query

import (
	"testing"

	c "../config"
)

func TestN3LoadConfig(t *testing.T) {
	InitClient(c.FromFile("./config.toml", "../config/config.toml"))
}

func TestQueryMeta(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)

	s, p, o, v := Meta("D3E34F41-9D75-101A-8C3D-00AA001A1652", "V") //         *** n3node thinks it is claiming ticket ***
	for i := range s {
		fPln("----------------------------------------------------")
		fPf("%d # %d: Reply: %s\n%s\n%s \n", i, v[i], s[i], p[i], o[i])
	}
	fPln()
}

func TestQuery1(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)

	ObjectID := "738F4DF5-949F-4380-8186-8252440A6F6F"
	Object := ""
	s, p, o, _ := Data(ObjectID, "") //      ** root **
	fPln("Object:")
	for i := range s {
		Object = o[i]
		fPf("%-50s%-10s%s\n", s[i], p[i], o[i])
	}

	fPln("\nArray Info:")
	s, p, o, _ = Data(ObjectID, "[]") //     ** array **
	for i := range s {
		fPf("%-40s%-65s%10s\n", s[i], p[i], o[i])
	}

	fPln("\nStructure:")
	s, p, o, _ = Data(Object, "::") //       ** struct **
	for i := range s {
		fPf("%-65s%-10s%s\n", s[i], p[i], o[i])
	}

	fPln("\nValues:")
	s, p, o, _ = Data(ObjectID, Object) //   ** values **
	for i := range s {
		fPf("%-40s%-85s%s\n", s[i], p[i], o[i])
	}
}

func TestQuery2(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)

	s1, _, _, _ := Data("", "xapi ~ actor ~ name", "Samuel Busse")
	if len(s1) == 0 {
		fPln("nothing found 1")
		return
	}
	s2, _, _, _ := Data("", "xapi ~ object ~ id", "http://example.com/assignments/Geography-8-1-B:4")
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
		s, p, o, _ := Data(ObjectID, "") //     ** root **
		fPln("Object:")
		for i := range s {
			Object = o[i]
			fPf("%-50s%-10s%s\n", s[i], p[i], o[i])
		}

		fPln("\nArray Info:")
		s, p, o, _ = Data(ObjectID, "[]") //    ** array ** /* context must end with '-xapi' */
		for i := range s {
			fPf("%-50s%-50s%s\n", s[i], p[i], o[i])
		}

		fPln("\nStructure:")
		s, p, o, _ = Data(Object, "::") //      ** struct **
		for i := range s {
			fPf("%-50s%-10s%s\n", s[i], p[i], o[i])
		}

		fPln("\nValues:")
		s, p, o, _ = Data(ObjectID, Object) //  ** values **
		for i := range s {
			fPf("%-50s%-50s%s\n", s[i], p[i], o[i])
		}

	} else {
		fPln("nothing found 1 & 2")
		return
	}
}
