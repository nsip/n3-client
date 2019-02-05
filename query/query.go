package query

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
	if g.N3pub == nil {
		g.N3pub = Must(n3grpc.NewPublisher(Cfg.RPC.Server, Cfg.RPC.Port)).(*n3grpc.Publisher)
	}
}

func query(t g.SQType, sp []string) (s, p, o []string, v []int64) {
	if Cfg == nil || g.N3pub == nil {
		Cfg = c.GetConfig("./config.toml", "../config/config.toml")
		Init(Cfg)
	}

	qTuple := &pb.SPOTuple{Subject: sp[0], Predicate: sp[1], Object: ""}
	ctx := u.CaseAssign(t, g.SIF, g.XAPI, g.META_SIF, g.META_XAPI, Cfg.RPC.CtxSif, Cfg.RPC.CtxXapi, Cfg.RPC.CtxMetaSif, Cfg.RPC.CtxMetaXapi).(string)
	for _, t := range g.N3pub.Query(qTuple, Cfg.RPC.Namespace, ctx) {
		s, p, o, v = append(s, t.Subject), append(p, t.Predicate), append(o, t.Object), append(v, t.Version)
	}
	return
}

// Sif :
func Sif(sp ...string) (s, p, o []string, v []int64) {
	return query(g.SIF, sp)
}

// Xapi :
func Xapi(sp ...string) (s, p, o []string, v []int64) {
	return query(g.XAPI, sp)
}

// Meta :
func Meta(t g.SQType, sp ...string) (s, p, o []string, v []int64) {
	switch t {
	case g.SIF:
		return query(g.META_SIF, sp)
	case g.XAPI:
		return query(g.META_XAPI, sp)
	default:
		panic("Meta: qType is not supported!")
	}
}
