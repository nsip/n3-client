package query

import (
	"testing"

	c "../config"
	g "../global"
)

func TestN3LoadConfig(t *testing.T) {
	cfg := c.GetConfig("./config.toml", "../config/config.toml")
	// fPln(cfg.RPC)
	// fPln(cfg.Filewatcher)
	Init(cfg)
}

func TestQueryMetaSif(t *testing.T) {
	defer func() { PH(recover(), Cfg.Global.ErrLog, true) }()
	TestN3LoadConfig(t)

	s, p, o, v := Meta(g.SIF, "D3E34F41-9D75-101A-8C3D-00AA001A1652", "V")	
	for i := range s {
		fPln("----------------------------------------------------")
		fPf("%d # %d: Reply: %s\n%s\n%s \n", i, v[i], s[i], p[i], o[i])
	}
	fPln()
}

func TestQuerySif(t *testing.T) {
	defer func() { PH(recover(), Cfg.Global.ErrLog, true) }()
	TestN3LoadConfig(t)
	s, p, o, v := Sif("D3E34F41-9D75-101A-8C3D-00AA001A1652", "StaffPersonal") //Sif("0E11C01D-54A2-4E9F-8C67-4FE2540BA6C8", "StaffPersonal") //
	// s, p, o, v := Sif("9269671A-BB89-4281-B20D-668C1D7FFD05", "TeachingGroup") /* context must end with '-sif' */
	fPln(len(s))
	for i := range s {
		fPln("----------------------------------------------------")
		fPf("%d # %d: Reply: %s\n%s\n%s \n", i, v[i], s[i], p[i], o[i])
	}
	fPln()
}

func TestQueryXapi(t *testing.T) {
	defer func() { PH(recover(), Cfg.Global.ErrLog, true) }()
	TestN3LoadConfig(t)
	s, p, o, v := Xapi("D3E34F41-9D75-101A-8C3D-00AA001A1652", "actor") /* context must end with '-xapi' */
	fPln(len(s))
	for i := range s {
		fPln("----------------------------------------------------")
		fPf("%d # %d: Reply: %s\n%s\n%s \n", i, v[i], s[i], p[i], o[i])
	}
	fPln()
}
