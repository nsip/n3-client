package publish

import (
	"errors"
	"time"

	g "github.com/nsip/n3-client/global"
	q "github.com/nsip/n3-client/query"
	xjy "github.com/nsip/n3-client/xjy"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/messages/pb"
)

// Junk :
func Junk(ctx string, n int) {
	for i := 0; i < n; i++ {
		tuple := must(messages.NewTuple("subject", "predicate", "object")).(*pb.SPOTuple)
		tuple.Version = int64(i)
		pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
		pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx+"-meta"))
	}
}

// Send :
func Send(ctx, subject, predicate, object string, ver int64) {
	tuple := must(messages.NewTuple(subject, predicate, object)).(*pb.SPOTuple)
	tuple.Version = ver
	if !g.Cfg.Debug.TrialPub {
		pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
	}
}

// Terminate :
func Terminate(ctx, objID, termID string, ver int64) {
	tuple := must(messages.NewTuple(termID, g.MARKTerm, objID)).(*pb.SPOTuple)
	tuple.Version = ver
	if !g.Cfg.Debug.TrialPub {
		pe(g.N3clt.Publish(tuple, g.Cfg.RPC.Namespace, ctx))
	}
}

// RequireVer : verType ( "V" / "A" / "S" )
func RequireVer(ctx, objID, verType string) (ver int64, termID string, err error) {
	_, p, o, _ := q.Meta(ctx, objID, verType)
	if len(o) == 0 {
		return -1, "", fEf("RequireVer Error: Are you authorised to n3node? OR using a dead object ID <%s> ?", objID)
	}
	ver, termID = S(o[0]).ToInt64()+1, p[0]
	return
}

/************************************************************/

// Pub2Node :
func Pub2Node(ctx, str, dfltRoot string) (IDs, Objs []string, nV, nS, nA int, err error) {
	prevIDs, termIDs := "", ""
	prevIDa, termIDa := "", ""
	prevIDv, termIDv, prevTermIDv := "", "", ""
	verS, verA, verV := int64(1), int64(1), int64(1)
	termIDvList := []string{}
	e := errors.New("")

	switch IF(IsJSON(str), g.JSON, g.XML).(g.DataType) {
	case g.XML:
		{
			strMod := prepXML(str)

			IDs, Objs, err = xjy.XMLInfoScan(strMod, g.DELIPath,
				func(p, id string, v []string, lastObjTuple bool) error {
					// fPln("S ---> ", p, "::", v)
					id = "::" + id
					defer func() { verS, nS, prevIDs = verS+1, nS+1, id }()
					if id != prevIDs {
						if verS, termIDs, e = RequireVer(ctx, id, "S"); e != nil {
							return e
						}
						// fPln("Got Ver S:", verS, termIDs)
					}
					Send(ctx, id, p, sJ(v, g.DELIChild), verS)
					if lastObjTuple {
						Terminate(ctx, id, termIDs, verS+1) //                 *** STRUCT terminator ***
					}
					return nil
				},
				func(p, id string, n int, lastObjTuple bool) error {
					// fPln("A ---> ", p, id, n)
					id = "[]" + id
					defer func() { verA, nA, prevIDa = verA+1, nA+1, id }()
					if id != prevIDa {
						if verA, termIDa, e = RequireVer(ctx, id, "A"); e != nil {
							return e
						}
						// fPln("Got Ver A:", verA, termIDa)
					}
					Send(ctx, id, p, fSf("%d", n), verA)
					if lastObjTuple {
						Terminate(ctx, id, termIDa, verA+1) //                 *** ARRAY terminator ***
					}
					return nil
				},
			)

			if err != nil {
				return
			}

			err = xjy.YAMLScan(strMod, dfltRoot, g.DELIPath, IDs, g.XML,
				func(p, v, id string) error {
					defer func() { verV, nV, prevIDv, prevTermIDv = verV+1, nV+1, id, termIDv }()
					// fPf("V ---> %-70s : %-36s : %-36s\n", p, v, id)
					if id != prevIDv {
						if prevIDv != "" {
							Terminate(ctx, prevIDv, prevTermIDv, verV)
						}
						if verV, termIDv, e = RequireVer(ctx, id, "V"); e != nil {
							return e
						}
						// fPln("Got Ver V:", verV, termIDv)
						termIDvList = append(termIDvList, termIDv)
					}
					if l := len(v); l > 2 && v[0] == '\'' && v[l-1] == '\'' {
						v = v[1 : l-1]
					}
					Send(ctx, id, p, v, verV)
					return nil
				})
			Terminate(ctx, prevIDv, prevTermIDv, verV) //                      *** VALUES terminator ***
			termIDvList = append(termIDvList, prevTermIDv)

			if err != nil {
				return
			}

			postpXML(ctx, str, IDs, Objs)

		} // XML

	case g.JSON:
		{
			strMod := prepJSON(str)

			IDs, Objs, err = xjy.JSONObjScan(strMod, dfltRoot,
				func(p, id string, v []string, lastObjTuple bool) error {
					id = "::" + id
					defer func() { verS, nS, prevIDs = verS+1, nS+1, id }()
					if id != prevIDs {
						if verS, termIDs, e = RequireVer(ctx, id, "S"); e != nil {
							return e
						}
						// fPln("Got Ver S:", verS, termIDs)
					}
					Send(ctx, id, p, sJ(v, g.DELIChild), verS)
					if lastObjTuple {
						Terminate(ctx, id, termIDs, verS+1)
					}
					return nil
				},
				func(p, id string, n int, lastObjTuple bool) error {
					id = "[]" + id
					defer func() { verA, nA, prevIDa = verA+1, nA+1, id }()
					if id != prevIDa {
						if verA, termIDa, e = RequireVer(ctx, id, "A"); e != nil {
							return e
						}
						// fPln("Got Ver A:", verA, termIDa)
					}
					Send(ctx, id, p, fSf("%d", n), verA)
					if lastObjTuple {
						Terminate(ctx, id, termIDa, verA+1)
					}
					return nil
				},
			)

			if err != nil {
				return
			}

			err = xjy.YAMLScan(strMod, dfltRoot, g.DELIPath, IDs, g.JSON,
				func(p, v, id string) error {
					defer func() { verV, nV, prevIDv, prevTermIDv = verV+1, nV+1, id, termIDv }()
					// fPf("V ---> %-70s : %-36s : %-36s\n", p, v, id)
					if id != prevIDv {
						if prevIDv != "" {
							Terminate(ctx, prevIDv, prevTermIDv, verV)
						}
						if verV, termIDv, e = RequireVer(ctx, id, "V"); e != nil {
							return e
						}
						// fPln("Got Ver V:", verV, termIDv)
						termIDvList = append(termIDvList, termIDv)
					}
					if l := len(v); l > 2 && v[0] == '\'' && v[l-1] == '\'' {
						v = v[1 : l-1]
					}
					Send(ctx, id, p, v, verV)
					return nil
				})

			Terminate(ctx, prevIDv, prevTermIDv, verV) //                      *** OBJECT terminator ***
			termIDvList = append(termIDvList, prevTermIDv)

			if err != nil {
				return
			}

			postpJSON(ctx, str, IDs, Objs)

		} // JSON

	} // case

	return

	// DOING: DB Storing Check
	otstdTermIDvList := []string{}
AGAIN:
	fPln("checking...")
	for _, termID := range termIDvList {
		if objIDList, _, _, _ := q.Data(ctx, termID, g.MARKTerm); objIDList == nil || len(objIDList) == 0 {
			otstdTermIDvList = append(otstdTermIDvList, termID)
		}
	}
	if len(otstdTermIDvList) > 0 {
		termIDvList = otstdTermIDvList
		otstdTermIDvList = []string{}
		time.Sleep(200 * time.Millisecond)
		goto AGAIN
	}

	return
}

// Pub2NodeAsyn :
func Pub2NodeAsyn(ctx, str, dfltRoot string, IDs, Objs chan []string, nV, nS, nA chan int) error {
	ids, objs, nv, ns, na, e := Pub2Node(ctx, str, dfltRoot)
	IDs <- ids
	Objs <- objs
	nV <- nv
	nS <- ns
	nA <- na
	return e
}
