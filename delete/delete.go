package delete

import (
	g "github.com/nsip/n3-client/global"
	"github.com/nsip/n3-messages/messages/pb"
)

// Del :
func Del(ctx, subject string) {
	dTuple := &pb.SPOTuple{Subject: subject, Predicate: g.MARKDead}
	if !g.Cfg.Debug.TrialPub {
		pe(g.N3clt.Publish(dTuple, g.Cfg.RPC.Namespace, ctx))
	}
}

// DelBat :
func DelBat(ctx string, subjects ...string) {
	for _, s := range subjects {
		Del(ctx, s)
	}
}
