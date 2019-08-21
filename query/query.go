package query

import (
	g "github.com/nsip/n3-client/global"
	"github.com/nsip/n3-messages/messages/pb"
)

func query(ctx string, metaQry bool, spo []string) (s, p, o []string, v []int64) {
	qTuple := &pb.SPOTuple{}
	switch len(spo) {
	case 2:
		qTuple = &pb.SPOTuple{Subject: spo[0], Predicate: spo[1]}
	case 3:
		qTuple = &pb.SPOTuple{Subject: spo[0], Predicate: spo[1], Object: spo[2]}
	default:
		panic("subject & predicate must be provided to general query. <empty string>-subject & predicate & object must be provided to id query")
	}

	if !g.Cfg.Debug.TrialQry {
		for _, t := range g.N3clt.Query(qTuple, g.Cfg.RPC.Namespace, IF(!metaQry, ctx, ctx+"-meta").(string)) {
			s, p, o, v = append(s, t.Subject), append(p, t.Predicate), append(o, t.Object), append(v, t.Version)
		}
	} else {
		s, p, o, v = append(s, spo[0]), append(p, spo[1]), append(o, "-999"), append(v, -1)
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
