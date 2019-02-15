package delete

import (
	"testing"
	"time"

	c "../config"
)

func TestN3LoadConfig(t *testing.T) {
	Init(c.GetConfig("./config.toml", "../config/config.toml"))
}

func TestDelete(t *testing.T) {
	defer func() { PH(recover(), Cfg.Global.ErrLog) }()
	TestN3LoadConfig(t)
	Sif("D3E34F41-9D75-101A-8C3D-00AA001A1655") //"9269671A-BB89-4281-B20D-668C1D7FFD05") /* context must end with '-sif' */
	time.Sleep(1 * time.Second)
}
