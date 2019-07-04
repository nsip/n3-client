package publish

import (
	c "../config"
	g "../global"
	q "../query"
	xjy "../xjy"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/messages/pb"
	"github.com/nsip/n3-messages/n3grpc"
)

// Junk :
func Junk(ctx string, n int) {
	pc(g.Cfg == nil || g.N3clt == nil, fEf("Missing Init, do 'Init(&config) before sending'\n"))
	for i := 0; i < n; i++ {
		tuple := must(messages.NewTuple("ab", "pre", "obj")).(*pb.SPOTuple)
		tuple.Version = int64(i)
		pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
	}
}

// Terminate :
func Terminate(ctx, objID, termID string, ver int64) {
	if g.Cfg == nil || g.N3clt == nil {
		InitClient(c.FromFile("../build/config.toml"))
	}
	tuple := must(messages.NewTuple(termID, g.MARKTerm, objID)).(*pb.SPOTuple)
	tuple.Version = ver
	pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
}

// RequireVer : verType ( "V" / "A" / "S" )
func RequireVer(ctx, objID, verType string) (ver int64, termID string) {
	if g.Cfg == nil || g.N3clt == nil {
		InitClient(c.FromFile("../build/config.toml"))
	}
	_, p, o, _ := q.Meta(ctx, objID, verType)
	pc(len(o) == 0, fEf("Got Version Error, Dead ObjectID: %s", objID))
	ver, termID = S(o[0]).ToInt64()+1, p[0]
	return
}

/************************************************************/

// InitClient :
func InitClient(config *c.Config) {
	pc(config == nil, fEf("Init Config"))
	g.Cfg = config
	g.N3clt = IF(g.N3clt == nil, n3grpc.NewClient(g.Cfg.RPC.Server, g.Cfg.RPC.Port), g.N3clt).(*n3grpc.Client)
}

// Pub2Node :
func Pub2Node(ctx, str, idmark, dfltRoot string) (IDs, Objs []string, nV, nS, nA int) {
	pc(g.Cfg == nil || g.N3clt == nil, fEf("Missing Sending Init, do 'Init(&config) before sending'\n"))

	prevIDs, termIDs := "", ""
	prevIDa, termIDa := "", ""
	prevIDv, termIDv, prevTermIDv := "", "", ""
	verS, verA, verV := int64(1), int64(1), int64(1)

	switch IF(IsJSON(str), g.JSON, g.XML).(g.DataType) {
	case g.XML:
		{
			strMod := prepXML(str)

			IDs, Objs = xjy.XMLInfoScan(strMod, idmark, g.DELIPath,
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

			xjy.YAMLScan(strMod, idmark, dfltRoot, IDs, g.XML,
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

			postpXML(str, IDs, Objs)

		} // XML

	case g.JSON:
		{
			strMod := prepJSON(str)

			IDs, Objs = xjy.JSONObjScan(strMod, idmark, dfltRoot,
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

			xjy.YAMLScan(strMod, idmark, dfltRoot, IDs, g.JSON,
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

			postpJSON(str, IDs, Objs)

		} // JSON

	} // case

	return
}
