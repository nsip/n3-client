package delete

import (
	c "../config"
	g "../global"
	u "github.com/cdutwhu/go-util"
	"github.com/nsip/n3-messages/messages/pb"
	"github.com/nsip/n3-messages/n3grpc"
)

// Init :
func Init(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	Cfg = config
	g.N3clt = u.TerOp(g.N3clt == nil, n3grpc.NewClient(Cfg.RPC.Server, Cfg.RPC.Port), g.N3clt).(*n3grpc.Client)
}

func del(t g.SQDType, sub string) {
	if Cfg == nil || g.N3clt == nil {
		Init(c.GetConfig("./config.toml", "../config/config.toml"))
	}

	dTuple := &pb.SPOTuple{Subject: sub, Predicate: DEADMARK}
	ctx := u.CaseAssign(t, g.SIF, g.XAPI, g.META_SIF, g.META_XAPI, Cfg.RPC.CtxSif, Cfg.RPC.CtxXapi, Cfg.RPC.CtxMetaSif, Cfg.RPC.CtxMetaXapi).(string)
	g.N3clt.Publish(dTuple, Cfg.RPC.Namespace, ctx)	
}

// Sif :
func Sif(subject string) {
	del(g.SIF, subject)
}

// Xapi :
func Xapi(subject string) {
	del(g.XAPI, subject)
}

// Meta :
func Meta(t g.SQDType, subject string) {
	switch t {
	case g.SIF:
		del(g.META_SIF, subject)
	case g.XAPI:
		del(g.META_XAPI, subject)
	default:
		panic("Meta: SQDType is not supported!")
	}
}
