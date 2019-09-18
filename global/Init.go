package global

import (
	c "github.com/nsip/n3-client/config"
	"github.com/nsip/n3-messages/n3grpc"
)

// Init :
func Init() {
	Cfg = c.FromFile("./config.toml", "../config.toml") // exe uses the first, unit tests use the second.
	pc(Cfg == nil, fEf("Init Config @ Cfg"))

	N3clt = n3grpc.NewClient(Cfg.RPC.Server, Cfg.RPC.Port)
	pc(N3clt == nil, fEf("Init Config @ N3clt"))
}
