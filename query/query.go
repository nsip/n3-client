package query

import (
	c "../config"
	g "../global"

	"github.com/nsip/n3-messages/messages/pb"
	"github.com/nsip/n3-messages/n3grpc"
)

// InitClient :
func InitClient(config *c.Config) {
	pc(config == nil, fEf("Init Config"))
	g.Cfg = config
	g.N3clt = IF(g.N3clt == nil, n3grpc.NewClient(g.Cfg.RPC.Server, g.Cfg.RPC.Port), g.N3clt).(*n3grpc.Client)
}

func query(ctx string, metaQry bool, spo []string) (s, p, o []string, v []int64) {
	if g.Cfg == nil || g.N3clt == nil {
		InitClient(c.FromFile("../build/config.toml"))
	}

	qTuple := &pb.SPOTuple{}
	switch len(spo) {
	case 2:
		qTuple = &pb.SPOTuple{Subject: spo[0], Predicate: spo[1]}
	case 3:
		qTuple = &pb.SPOTuple{Subject: spo[0], Predicate: spo[1], Object: spo[2]}
	default:
		panic("subject & predicate must be provided to general query. <empty string>-subject & predicate & object must be provided to id query")
	}

	for _, t := range g.N3clt.Query(qTuple, g.Cfg.RPC.Namespace, IF(!metaQry, ctx, ctx+"-meta").(string)) {
		s, p, o, v = append(s, t.Subject), append(p, t.Predicate), append(o, t.Object), append(v, t.Version)
	}
	return
}

// Data :
func Data(ctx string, spo ...string) (s, p, o []string, v []int64) {
	return query(ctx, false, spo)
}

// Meta :
func Meta(ctx string, spo ...string) (s, p, o []string, v []int64) {
	return query(ctx, true, spo)
}
