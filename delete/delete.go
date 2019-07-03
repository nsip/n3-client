package delete

import (
	c "../config"
	g "../global"
	"github.com/nsip/n3-messages/messages/pb"
	"github.com/nsip/n3-messages/n3grpc"
)

// Init :
func Init(config *c.Config) {
	pc(config == nil, fEf("Init Config"))
	g.Cfg = config
	g.N3clt = IF(g.N3clt == nil, n3grpc.NewClient(g.Cfg.RPC.Server, g.Cfg.RPC.Port), g.N3clt).(*n3grpc.Client)
}

// Del :
func Del(ctx, subject string) {
	if g.Cfg == nil || g.N3clt == nil {
		Init(c.FromFile("../build/config.toml"))
	}
	dTuple := &pb.SPOTuple{Subject: subject, Predicate: DEADMARK}
	pe(g.N3clt.Publish(dTuple, g.Cfg.RPC.Namespace, ctx))
}

// DelBat :
func DelBat(ctx string, subjects ...string) {
	for _, s := range subjects {
		Del(ctx, s)
	}
}
