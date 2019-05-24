package delete

import (
	c "../config"
	g "../global"
	"github.com/nsip/n3-messages/messages/pb"
	"github.com/nsip/n3-messages/n3grpc"
)

// Init :
func Init(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	CFG = config
	g.N3clt = IF(g.N3clt == nil, n3grpc.NewClient(CFG.RPC.Server, CFG.RPC.Port), g.N3clt).(*n3grpc.Client)
}

// Del :
func Del(subject string) {
	if CFG == nil || g.N3clt == nil {
		Init(c.FromFile("./config.toml", "../config/config.toml"))
	}
	dTuple := &pb.SPOTuple{Subject: subject, Predicate: DEADMARK}
	PE(g.N3clt.Publish(dTuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
}

// DelBat :
func DelBat(subjects ...string) {
	for _, s := range subjects {
		Del(s)
	}
}
