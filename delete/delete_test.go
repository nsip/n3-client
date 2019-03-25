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
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)
	Sif("E6271E2D-700E-4286-AA81-58FEF49060C1") /* context must end with '-sif' */
	time.Sleep(1 * time.Second)
}
