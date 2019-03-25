package query

import (
	c "../config"
	g "../global"
	u "github.com/cdutwhu/go-util"
	"github.com/nsip/n3-messages/messages/pb"
	"github.com/nsip/n3-messages/n3grpc"
)

// InitFrom :
func InitFrom(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	CFG = config
	g.N3clt = u.TerOp(g.N3clt == nil, n3grpc.NewClient(CFG.RPC.Server, CFG.RPC.Port), g.N3clt).(*n3grpc.Client)
}

func query(t g.SQDType, spo []string) (s, p, o []string, v []int64) {
	if CFG == nil || g.N3clt == nil {
		InitFrom(c.FromFile("./config.toml", "../config/config.toml"))
	}

	qTuple := &pb.SPOTuple{}
	if len(spo) == 2 {
		qTuple = &pb.SPOTuple{Subject: spo[0], Predicate: spo[1]}
	} else if len(spo) == 3 {
		qTuple = &pb.SPOTuple{Subject: spo[0], Predicate: spo[1], Object: spo[2]}
	} else {
		panic("subject & predicate must be provided to general query, empty-subject & predicate & object must be provided to id query")
	}

	ctx := u.CaseAssign(t, g.SIF, g.XAPI, g.META_SIF, g.META_XAPI, CFG.RPC.CtxSif, CFG.RPC.CtxXapi, CFG.RPC.CtxMetaSif, CFG.RPC.CtxMetaXapi).(string)
	for _, t := range g.N3clt.Query(qTuple, CFG.RPC.Namespace, ctx) {
		s, p, o, v = append(s, t.Subject), append(p, t.Predicate), append(o, t.Object), append(v, t.Version)
	}
	return
}

// Sif :
func Sif(spo ...string) (s, p, o []string, v []int64) {
	return query(g.SIF, spo)
}

// Xapi :
func Xapi(spo ...string) (s, p, o []string, v []int64) {
	return query(g.XAPI, spo)
}

// Meta :
func Meta(t g.SQDType, spo ...string) (s, p, o []string, v []int64) {
	switch t {
	case g.SIF:
		return query(g.META_SIF, spo)
	case g.XAPI:
		return query(g.META_XAPI, spo)
	default:
		panic("Meta: qType is not supported!")
	}
}
