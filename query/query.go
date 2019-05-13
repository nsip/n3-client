package query

import (
	c "../config"
	g "../global"

	"github.com/nsip/n3-messages/messages/pb"
	"github.com/nsip/n3-messages/n3grpc"
)

// InitFrom :
func InitFrom(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	CFG = config
	g.N3clt = IF(g.N3clt == nil, n3grpc.NewClient(CFG.RPC.Server, CFG.RPC.Port), g.N3clt).(*n3grpc.Client)
}

func query(DataOrMeta int, spo []string) (s, p, o []string, v []int64) {
	if CFG == nil || g.N3clt == nil {
		InitFrom(c.FromFile("./config.toml", "../config/config.toml"))
	}

	qTuple := &pb.SPOTuple{}
	switch len(spo) {
	case 2:
		qTuple = &pb.SPOTuple{Subject: spo[0], Predicate: spo[1]}
	case 3:
		qTuple = &pb.SPOTuple{Subject: spo[0], Predicate: spo[1], Object: spo[2]}
	default:
		panic("subject & predicate must be provided to general query, empty-subject & predicate & object must be provided to id query")
	}

	ctx := IF(DataOrMeta == 1, CFG.RPC.Ctx, CFG.RPC.CtxMeta).(string)
	for _, t := range g.N3clt.Query(qTuple, CFG.RPC.Namespace, ctx) {
		s, p, o, v = append(s, t.Subject), append(p, t.Predicate), append(o, t.Object), append(v, t.Version)
	}
	return
}

// Data :
func Data(spo ...string) (s, p, o []string, v []int64) {
	return query(1, spo)
}

// Meta :
func Meta(spo ...string) (s, p, o []string, v []int64) {
	return query(0, spo)
}
