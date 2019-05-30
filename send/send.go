package send

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
func Junk(n int) {
	PC(CFG == nil || g.N3clt == nil, fEf("Missing Init, do 'Init(&config) before sending'\n"))
	for i := 0; i < n; i++ {
		tuple := Must(messages.NewTuple("ab", "pre", "obj")).(*pb.SPOTuple)
		tuple.Version = int64(i)
		PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
	}
}

// Terminate :
func Terminate(objID, termID string, ver int64) {
	if CFG == nil || g.N3clt == nil {
		InitClient(c.FromFile("./config.toml", "../config/config.toml"))
	}
	tuple := Must(messages.NewTuple(termID, TERMMARK, objID)).(*pb.SPOTuple)
	tuple.Version = ver
	PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
}

// RequireVer : verType ( "V" / "A" / "S" )
func RequireVer(objID, verType string) (ver int64, termID string) {
	if CFG == nil || g.N3clt == nil {
		InitClient(c.FromFile("./config.toml", "../config/config.toml"))
	}
	_, p, o, _ := q.Meta(objID, verType)
	PC(len(o) == 0, fEf("Got Version Error, Dead ObjectID: %s", objID))
	ver, termID = Str(o[0]).ToInt64()+1, p[0]
	return
}

/************************************************************/

// InitClient :
func InitClient(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	CFG = config
	g.N3clt = IF(g.N3clt == nil, n3grpc.NewClient(CFG.RPC.Server, CFG.RPC.Port), g.N3clt).(*n3grpc.Client)
}

// ToNode :
func ToNode(str, idmark, dfltRoot string) (IDs []string, cntV, cntS, cntA int) {

	PC(CFG == nil || g.N3clt == nil, fEf("Missing Sending Init, do 'Init(&config) before sending'\n"))
	data := Str(str)
	data.SetEnC()

	prevIDs, termIDs := "", ""
	prevIDa, termIDa := "", ""
	prevIDv, termIDv, prevTermIDv := "", "", ""
	verS, verA, verV := int64(1), int64(1), int64(1)

	switch IF(IsJSON(str), g.JSON, g.XML).(g.SQDType) {
	case g.XML:
		{
			IDs = xjy.XMLInfoScan(data.V(), idmark, PATH_DEL,
				func(p, id string, v []string, lastObjTuple bool) {
					// fPln("S ---> ", p, "::", v)
					id = "::" + id
					defer func() { verS, cntS, prevIDs = verS+1, cntS+1, id }()
					if id != prevIDs {
						verS, termIDs = RequireVer(id, "S")
						fPln("Got Ver S:", verS, termIDs)
					}
					tuple := Must(messages.NewTuple(p, id, sJ(v, CHILD_DEL))).(*pb.SPOTuple)
					tuple.Version = verS
					PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
					if lastObjTuple {
						Terminate(id, termIDs, verS+1) //                              *** object struct terminator ***
					}
				},
				func(p, id string, n int, lastObjTuple bool) {
					// fPln("A ---> ", p, id, n)
					id = "[]" + id
					defer func() { verA, cntA, prevIDa = verA+1, cntA+1, id }()
					if id != prevIDa {
						verA, termIDa = RequireVer(id, "A")
						fPln("Got Ver A:", verA, termIDa)
					}
					tuple := Must(messages.NewTuple(p, id, fSf("%d", n))).(*pb.SPOTuple)
					tuple.Version = verA
					PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
					if lastObjTuple {
						Terminate(id, termIDa, verA+1) //                          *** object array terminator ***
					}
				},
			)

			xjy.YAMLScan(data.V(), idmark, dfltRoot, IDs, DT_XML,
				func(p, v, id string) {
					defer func() { verV, cntV, prevIDv, prevTermIDv = verV+1, cntV+1, id, termIDv }()
					// fPf("V ---> %-70s : %-36s : %-36s\n", p, v, id)
					if id != prevIDv {
						if prevIDv != "" {
							Terminate(prevIDv, prevTermIDv, verV)
						}
						verV, termIDv = RequireVer(id, "V")
						fPln("Got Ver V:", verV, termIDv)
					}
					tuple := Must(messages.NewTuple(id, p, v)).(*pb.SPOTuple)
					tuple.Version = verV
					PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
				})
			Terminate(prevIDv, prevTermIDv, verV) //                              *** object values terminator ***
		}

	case g.JSON:
		{
			IDs = xjy.JSONObjScan(data.V(), idmark, dfltRoot,
				func(p, id string, v []string, lastObjTuple bool) {
					id = "::" + id
					defer func() { verS, cntS, prevIDs = verS+1, cntS+1, id }()
					if id != prevIDs {
						verS, termIDs = RequireVer(id, "S")
						fPln("Got Ver S:", verS, termIDs)
					}
					tuple := Must(messages.NewTuple(p, id, sJ(v, CHILD_DEL))).(*pb.SPOTuple)
					tuple.Version = verS
					PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
					if lastObjTuple {
						Terminate(id, termIDs, verS+1)
					}
				},
				func(p, id string, n int, lastObjTuple bool) {
					id = "[]" + id
					defer func() { verA, cntA, prevIDa = verA+1, cntA+1, id }()
					if id != prevIDa {
						verA, termIDa = RequireVer(id, "A")
						fPln("Got Ver A:", verA, termIDa)
					}
					tuple := Must(messages.NewTuple(p, id, fSf("%d", n))).(*pb.SPOTuple)
					tuple.Version = verA
					PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
					if lastObjTuple {
						Terminate(id, termIDa, verA+1)
					}
				},
			)

			xjy.YAMLScan(data.V(), idmark, dfltRoot, IDs, DT_JSON,
				func(p, v, id string) {
					defer func() { verV, cntV, prevIDv, prevTermIDv = verV+1, cntV+1, id, termIDv }()
					// fPf("V ---> %-70s : %-36s : %-36s\n", p, v, id)
					if id != prevIDv {
						if prevIDv != "" {
							Terminate(prevIDv, prevTermIDv, verV)
						}
						verV, termIDv = RequireVer(id, "V")
						fPln("Got Ver V:", verV, termIDv)
					}
					tuple := Must(messages.NewTuple(id, p, v)).(*pb.SPOTuple)
					tuple.Version = verV
					PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
				})
			Terminate(prevIDv, prevTermIDv, verV) //                              *** object terminator ***
		}
	}

	return
}
