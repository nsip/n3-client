package publish

import (
	"time"

	g "../global"
	q "../query"
	xjy "../xjy"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/messages/pb"
)

// Send :
func Send(ctx, subject, predicate, object string) {
	tuple := must(messages.NewTuple(subject, predicate, object)).(*pb.SPOTuple)
	tuple.Version = 0
	pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
}

// Junk :
func Junk(ctx string, n int) {
	for i := 0; i < n; i++ {
		tuple := must(messages.NewTuple("subject", "predicate", "object")).(*pb.SPOTuple)
		tuple.Version = int64(i)
		pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
		pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx+"-meta"))
	}
}

// Terminate :
func Terminate(ctx, objID, termID string, ver int64) {
	tuple := must(messages.NewTuple(termID, g.MARKTerm, objID)).(*pb.SPOTuple)
	tuple.Version = ver
	pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
}

// RequireVer : verType ( "V" / "A" / "S" )
func RequireVer(ctx, objID, verType string) (ver int64, termID string) {
	_, p, o, _ := q.Meta(ctx, objID, verType)
	pc(len(o) == 0, fEf("Got Version Error, Dead ObjectID: %s", objID))
	ver, termID = S(o[0]).ToInt64()+1, p[0]
	return
}

/************************************************************/

// Pub2Node :
func Pub2Node(ctx, str, dfltRoot string) (IDs, Objs []string, nV, nS, nA int) {
	prevIDs, termIDs := "", ""
	prevIDa, termIDa := "", ""
	prevIDv, termIDv, prevTermIDv := "", "", ""
	verS, verA, verV := int64(1), int64(1), int64(1)

	switch IF(IsJSON(str), g.JSON, g.XML).(g.DataType) {
	case g.XML:
		{
			strMod := prepXML(str)

			IDs, Objs = xjy.XMLInfoScan(strMod, g.DELIPath,
				func(p, id string, v []string, lastObjTuple bool) {
					// fPln("S ---> ", p, "::", v)
					id = "::" + id
					defer func() { verS, nS, prevIDs = verS+1, nS+1, id }()
					if id != prevIDs {
						verS, termIDs = RequireVer(ctx, id, "S")
						// fPln("Got Ver S:", verS, termIDs)
					}
					tuple := must(messages.NewTuple(p, id, sJ(v, g.DELIChild))).(*pb.SPOTuple)
					tuple.Version = verS
					pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
					if lastObjTuple {
						Terminate(ctx, id, termIDs, verS+1) //                 *** object struct terminator ***
					}
				},
				func(p, id string, n int, lastObjTuple bool) {
					// fPln("A ---> ", p, id, n)
					id = "[]" + id
					defer func() { verA, nA, prevIDa = verA+1, nA+1, id }()
					if id != prevIDa {
						verA, termIDa = RequireVer(ctx, id, "A")
						// fPln("Got Ver A:", verA, termIDa)
					}
					tuple := must(messages.NewTuple(p, id, fSf("%d", n))).(*pb.SPOTuple)
					tuple.Version = verA
					pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
					if lastObjTuple {
						Terminate(ctx, id, termIDa, verA+1) //                 *** object array terminator ***
					}
				},
			)

			xjy.YAMLScan(strMod, dfltRoot, g.DELIPath, IDs, g.XML,
				func(p, v, id string) {
					defer func() { verV, nV, prevIDv, prevTermIDv = verV+1, nV+1, id, termIDv }()
					// fPf("V ---> %-70s : %-36s : %-36s\n", p, v, id)
					if id != prevIDv {
						if prevIDv != "" {
							Terminate(ctx, prevIDv, prevTermIDv, verV)
						}
						verV, termIDv = RequireVer(ctx, id, "V")
						// fPln("Got Ver V:", verV, termIDv)
					}
					if l := len(v); l > 2 && v[0] == '\'' && v[l-1] == '\'' {
						v = v[1 : l-1]
					}
					tuple := must(messages.NewTuple(id, p, v)).(*pb.SPOTuple)
					tuple.Version = verV
					pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
				})
			Terminate(ctx, prevIDv, prevTermIDv, verV) //                      *** object values terminator ***

			postpXML(ctx, str, IDs, Objs)

		} // XML

	case g.JSON:
		{
			strMod := prepJSON(str)

			IDs, Objs = xjy.JSONObjScan(strMod, dfltRoot,
				func(p, id string, v []string, lastObjTuple bool) {
					id = "::" + id
					defer func() { verS, nS, prevIDs = verS+1, nS+1, id }()
					if id != prevIDs {
						verS, termIDs = RequireVer(ctx, id, "S")
						// fPln("Got Ver S:", verS, termIDs)
					}
					tuple := must(messages.NewTuple(p, id, sJ(v, g.DELIChild))).(*pb.SPOTuple)
					tuple.Version = verS
					pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
					if lastObjTuple {
						Terminate(ctx, id, termIDs, verS+1)
					}
				},
				func(p, id string, n int, lastObjTuple bool) {
					id = "[]" + id
					defer func() { verA, nA, prevIDa = verA+1, nA+1, id }()
					if id != prevIDa {
						verA, termIDa = RequireVer(ctx, id, "A")
						// fPln("Got Ver A:", verA, termIDa)
					}
					tuple := must(messages.NewTuple(p, id, fSf("%d", n))).(*pb.SPOTuple)
					tuple.Version = verA
					pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
					if lastObjTuple {
						Terminate(ctx, id, termIDa, verA+1)
					}
				},
			)

			xjy.YAMLScan(strMod, dfltRoot, g.DELIPath, IDs, g.JSON,
				func(p, v, id string) {
					defer func() { verV, nV, prevIDv, prevTermIDv = verV+1, nV+1, id, termIDv }()
					// fPf("V ---> %-70s : %-36s : %-36s\n", p, v, id)
					if id != prevIDv {
						if prevIDv != "" {
							Terminate(ctx, prevIDv, prevTermIDv, verV)
						}
						verV, termIDv = RequireVer(ctx, id, "V")
						// fPln("Got Ver V:", verV, termIDv)
					}
					if l := len(v); l > 2 && v[0] == '\'' && v[l-1] == '\'' {
						v = v[1 : l-1]
					}
					tuple := must(messages.NewTuple(id, p, v)).(*pb.SPOTuple)
					tuple.Version = verV
					pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
				})
			Terminate(ctx, prevIDv, prevTermIDv, verV) //                      *** object terminator ***

			postpJSON(ctx, str, IDs, Objs)

		} // JSON

	} // case

	// TODO: DB Storing Check
	nStored := 0
	ticker := time.NewTicker(2000 * time.Millisecond)
	for range ticker.C {
		for _, ID := range IDs {
			if termIDList, _, _, _ := q.Data(ctx, "", g.MARKTerm, ID); termIDList != nil && len(termIDList) > 0 {
				nStored++
			}
		}
		if nStored == len(IDs) {
			fPln("All sent and stored")
			goto STORED
		}
	}

STORED:
	ticker.Stop()
	return
}
