package delete

import (
	c "../config"
	g "../global"
)

// Init :
func Init(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	Cfg = config
	if g.N3pub == nil {
		g.N3pub = Must(n3grpc.NewPublisher(Cfg.RPC.Server, Cfg.RPC.Port)).(*n3grpc.Publisher)
	}
}
