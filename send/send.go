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
func Terminate(objID, termID string) {
	defer func() { ver++ }()
	if CFG == nil || g.N3clt == nil {
		InitClientFrom(c.FromFile("./config.toml", "../config/config.toml"))
	}
	tuple := Must(messages.NewTuple(termID, TERMMARK, objID)).(*pb.SPOTuple)
	tuple.Version = ver
	PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
}

// RequireVer :
func RequireVer(objID string) (ver int64, termID string) {
	if CFG == nil || g.N3clt == nil {
		InitClientFrom(c.FromFile("./config.toml", "../config/config.toml"))
	}
	_, p, o, _ := q.Meta(objID, "V")
	PC(len(p) == 0, fEf("Got Version Error, Dead ObjectID is used"))
	// fPln(p, o)
	ver, termID = Str(o[0]).ToInt64()+1, p[0]
	return
}

/************************************************************/

// InitClientFrom :
func InitClientFrom(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	CFG = config
	g.N3clt = IF(g.N3clt == nil, n3grpc.NewClient(CFG.RPC.Server, CFG.RPC.Port), g.N3clt).(*n3grpc.Client)
}

// ToNode :
func ToNode(str string, dt g.SQDType) (cntV, cntS, cntA int) {
	PC(CFG == nil || g.N3clt == nil, fEf("Missing Sending Init, do 'Init(&config) before sending'\n"))
	data := Str(str)
	data.SetEnC()

	switch dt {
	case g.SIF:
		{
			mStructRecord := map[string][]string{}
			IDs := xjy.XMLInfoScan(data.V(), "RefId", PATH_DEL,
				func(p string, v []string) {
					if _, ok := mStructRecord[p]; !ok {
						defer func() { ver, cntS, mStructRecord[p] = ver+1, cntS+1, v }()
						vstr := sJ(v, CHILD_DEL)
						fPln("S ---> ", p, " : ", vstr)
						tuple := Must(messages.NewTuple(p, "::", vstr)).(*pb.SPOTuple)
						tuple.Version = ver
						PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
					}
				},
				func(p, v string, n int) {
					// fPln("A ---> ", p, v, n)
					if n > 1 {
						defer func() { ver, cntA = ver+1, cntA+1 }()
						fPln("A ---> ", p, v, n)
						tuple := Must(messages.NewTuple(p, v, fSf("%d", n))).(*pb.SPOTuple)
						tuple.Version = ver
						PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
					}
				},
			)

			prevID, termID, prevTermID := "", "", ""
			xjy.YAMLScan(data.V(), "RefId", IDs, DT_XML, func(p, v, id string) {
				defer func() { ver, cntV, prevID, prevTermID = ver+1, cntV+1, id, termID }()
				fPf("V ---> %-70s : %-36s : %-36s\n", p, v, id)
				if id != prevID {
					if prevID != "" {
						Terminate(prevID, prevTermID)
					}
					ver, termID = RequireVer(id)
					fPln("Got:", ver, termID)
				}
				tuple := Must(messages.NewTuple(id, p, v)).(*pb.SPOTuple)
				tuple.Version = ver
				PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
			})
			Terminate(prevID, prevTermID) //                              *** object terminator ***
		}

	case g.XAPI:
		{
			mStructRecord := map[string][]string{}
			IDs := xjy.JSONObjScan(data.V(), "id", "xapi",
				func(p string, v []string) {
					if _, ok := mStructRecord[p]; !ok {
						defer func() { ver, cntS, mStructRecord[p] = ver+1, cntS+1, v }()
						vstr := sJ(v, CHILD_DEL)
						fPf("S ---> %-70s:: %s\n", p, vstr)
						tuple := Must(messages.NewTuple(p, "::", vstr)).(*pb.SPOTuple)
						tuple.Version = ver
						PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
					}
				},
				func(p, v string, n int) {
					fPf("A ---> %-70s[]%s -- [%d]\n", p, v, n)
					defer func() { ver, cntA = ver+1, cntA+1 }()
					tuple := Must(messages.NewTuple(p, v, fSf("%d", n))).(*pb.SPOTuple)
					tuple.Version = ver
					PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
				},
			)

			prevID, termID, prevTermID := "", "", ""
			xjy.YAMLScan(data.V(), "id", IDs, DT_JSON, func(p, v, id string) {
				defer func() { ver, cntV, prevID, prevTermID = ver+1, cntV+1, id, termID }()
				fPf("V ---> %-70s : %-36s : %-36s\n", p, v, id)
				if id != prevID {
					if prevID != "" {
						Terminate(prevID, prevTermID)
					}
					ver, termID = RequireVer(id)
					fPln("Got:", ver, termID)
				}
				tuple := Must(messages.NewTuple(id, p, v)).(*pb.SPOTuple)
				tuple.Version = ver
				PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.Ctx))
			})
			Terminate(prevID, prevTermID) //                              *** object terminator ***
		}
	}

	return
}
