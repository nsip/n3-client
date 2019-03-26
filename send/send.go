package send

import (
	"time"

	c "../config"
	g "../global"
	q "../query"
	"../xjy"
	u "github.com/cdutwhu/go-util"
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
		PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxSif))
	}
}

// Terminate :
func Terminate(t g.SQDType, objID, termID string) {
	defer func() { ver++ }()
	if CFG == nil || g.N3clt == nil {
		InitFrom(c.FromFile("./config.toml", "../config/config.toml"))
	}

	tuple := Must(messages.NewTuple(termID, TERMMARK, objID)).(*pb.SPOTuple)
	tuple.Version = ver
	ctx := u.CaseAssign(t, g.SIF, g.XAPI, CFG.RPC.CtxSif, CFG.RPC.CtxXapi).(string)
	PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, ctx))
}

// RequireVer :
func RequireVer(t g.SQDType, objID string) (ver int64, termID string) {
	if CFG == nil || g.N3clt == nil {
		InitFrom(c.FromFile("./config.toml", "../config/config.toml"))
	}
	_, p, o, _ := q.Meta(t, objID, "V")
	PC(len(p) == 0, fEf("Got Version Error, Dead ObjectID is used"))
	// fPln(p, o)
	ver, termID = u.Str(o[0]).ToInt64()+1, p[0]
	return
}

/************************************************************/

// InitFrom :
func InitFrom(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	CFG = config
	g.N3clt = u.TerOp(g.N3clt == nil, n3grpc.NewClient(CFG.RPC.Server, CFG.RPC.Port), g.N3clt).(*n3grpc.Client)
}

// Sif : SendSif
func Sif(str string) (cntV, cntS, cntA int, termID string) {
	PC(CFG == nil || g.N3clt == nil, fEf("Missing Send Init, do 'Init(&config) before sending'\n"))

	sif, sqType := u.Str(str), g.SIF
	sif.SetEnC()
	PC(sif.L() == 0, fEf("Incoming string is invalid xml segment\n"))

	prevID, prevTermID := "", ""

	mapStructRecord := map[string]string{}
	n := sif.XMLSegsCount()
	prevEnd := 0
	for iObj := 1; iObj <= n; iObj++ {
		nextStart := u.TerOp(iObj == 1, 0, prevEnd+1).(int)
		_, xml, _, end := sif.S(nextStart, u.ALL).XMLSegPos(1, 1)
		prevEnd = end + nextStart

		// fPf("%d SIF *****************************************\n", iObj)

		// ************* Structure & Array ************* //

		xjy.XMLModelInfo(xml, "RefId", pathDel, childDel,
			func(p, v string) {
				if prevV, ok := mapStructRecord[p]; !ok || (ok && v != prevV && u.Str(v).FieldsSeqContain(prevV, childDel)) {
					defer func() { ver, cntS = ver+1, cntS+1 }()
					tuple := Must(messages.NewTuple(u.Str(p).RmPrefix(HEADTRIM).V(), "::", v)).(*pb.SPOTuple)
					tuple.Version = ver
					PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxSif))
					mapStructRecord[p] = v
				}
			},
			func(p, objID string, arrCnt int) {
				defer func() { ver, cntA = ver+1, cntA+1 }()
				tuple := Must(messages.NewTuple(u.Str(p).RmPrefix(HEADTRIM).V(), objID, u.I32(arrCnt).ToStr())).(*pb.SPOTuple)
				tuple.Version = ver
				PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxSif))
			},
		)

		// ************* VALUES ************* //

		yaml := xjy.Xstr2Y(xml) //                               *** must use current object piece ***
		// ioutil.WriteFile("sif.yaml", []byte(yaml), 0666)
		xjy.YAMLScan(yaml, "RefId", pathDel, xjy.XML, true, //   *** skipDir must be <true>, otherwise dir version might be incorrect number ***
			func(p, v, id string) {
				// fPln(p, v, id)
				defer func() { ver, cntV, prevID, prevTermID = ver+1, cntV+1, id, termID }()
				if id != prevID {
					ver, termID = RequireVer(sqType, id)
					// fPln("Got:", ver, termID)
				}
				tuple := Must(messages.NewTuple(id, u.Str(p).RmPrefix(HEADTRIM).V(), v)).(*pb.SPOTuple)
				tuple.Version = ver
				PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxSif))
			})
		Terminate(sqType, prevID, prevTermID) // *** object terminator ***
	}

CHECKOVER:
	if _, _, _, v := q.Sif(prevTermID, TERMMARK); v == nil || len(v) == 0 {
		time.Sleep(DELAY * time.Millisecond)
		goto CHECKOVER
	}

	lPln(fSf("<%06d> data tuples sent, <%06d> struct tuples sent, <%06d> array tuples sent\n", cntV, cntS, cntA))	

	/********************************************************************/

	// xjy.XMLModelInfo(sif.V(), "RefId", pathDel, childDel,
	// 	func(p, v string) {
	// 		defer func() { ver, cntS = ver+1, cntS+1 }()
	// 		tuple := Must(messages.NewTuple(u.Str(p).RmPrefix(HEADTRIM).V(), "::", v)).(*pb.SPOTuple)
	// 		tuple.Version = ver
	// 		PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxSif))
	// 	},
	// 	func(p, objID string, arrCnt int) {
	// 		defer func() { ver, cntA = ver+1, cntA+1 }()
	// 		tuple := Must(messages.NewTuple(u.Str(p).RmPrefix(HEADTRIM).V(), objID, u.I32(arrCnt).ToStr())).(*pb.SPOTuple)
	// 		tuple.Version = ver
	// 		PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxSif))
	// 	},
	// )

	/********************************************************************/

	// doneV, prevID, prevTermID := make(chan int), "", ""
	// yaml := xjy.Xstr2Y(sif.V())
	// // ioutil.WriteFile("sif.yaml", []byte(yaml), 0666)
	// go xjy.YAMLScanAsync(yaml, "RefId", pathDel, xjy.XML, true, // *** skipDir must be <true>, otherwise dir version might be incorrect number ***
	// 	func(p, v, id string) {
	// 		// fPln(p, v, id)
	// 		defer func() { ver, cntV, prevID, prevTermID = ver+1, cntV+1, id, termID }()
	// 		if id != prevID {
	// 			ver, termID = RequireVer(sqType, id)
	// 			fPln("Got:", ver, termID)
	// 			if prevID != "" {
	// 				Terminate(sqType, prevID, prevTermID)
	// 			}
	// 		}
	// 		tuple := Must(messages.NewTuple(id, u.Str(p).RmPrefix(HEADTRIM).V(), v)).(*pb.SPOTuple)
	// 		tuple.Version = ver
	// 		PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxSif))
	// 	},
	// 	doneV)
	// <-doneV

	/********************************************************************/

	// 	lPln(fSf("<%06d> data tuples sent, <%06d> struct tuples sent, <%06d> array tuples sent\n", cntV, cntS, cntA))

	// 	Terminate(sqType, prevID, prevTermID) // *** last object terminator ***
	// CHECK:
	// 	if _, _, _, v := q.Sif(prevTermID, TERMMARK); v == nil || len(v) == 0 {
	// 		time.Sleep(DELAY * time.Millisecond)
	// 		goto CHECK
	// 	}

	return
}

// Xapi : SendXapi
func Xapi(str string) (cntV, cntS, cntA int, termID string) {
	PC(CFG == nil, fEf("Missing Send Init, do 'Init(&config) before sending'\n"))

	xapi, sqType := u.Str(str), g.XAPI
	xapi.SetEnC()
	PC(xapi.L() == 0 || !xapi.IsJSON(), fEf("Incoming string is invalid json\n"))

	defaultRoot, addedRoot, prevID, prevTermID := "XAPI", false, "", ""
	mapStructRecord := map[string]string{}

	if xapi.T(u.BLANK).C(0) != '[' {
		xapi = u.Str("[ " + xapi.V() + " ]") //                                   *** wrap single json object into array ***
		PC(!xapi.IsJSON(), fEf("wrapped xapi array string is invalid json\n"))
	}

	if ok, jsonType, n := xapi.IsJSONRootArray(); ok {
		if jsonType == "Object" {
			prevEnd := 0
			for i := 1; i <= n; i++ {
				nextStart := u.TerOp(i == 1, 0, prevEnd+1).(int)
				json, _, end := xapi.S(nextStart, u.ALL).BracketsPos(u.BCurly, 1, 1) // ***
				prevEnd = end + nextStart

				addedRoot = xjy.JSONModelInfo(json.V(), "id", defaultRoot, pathDel, childDel,
					func(p, v string) {
						if i == 1 { //                                             *** only use one as sample to get struct info ***
							if prevV, ok := mapStructRecord[p]; !ok || (ok && v != prevV && u.Str(v).FieldsSeqContain(prevV, childDel)) {
								defer func() { ver, cntS = ver+1, cntS+1 }()
								tuple := Must(messages.NewTuple(p, "::", v)).(*pb.SPOTuple)
								tuple.Version = ver
								PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxXapi))
								mapStructRecord[p] = v
							}
						}
					},
					func(p, objID string, arrCnt int) {
						defer func() { ver, cntA = ver+1, cntA+1 }()
						tuple := Must(messages.NewTuple(p, objID, u.I32(arrCnt).ToStr())).(*pb.SPOTuple)
						tuple.Version = ver
						PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxXapi))
					},
				)

				/********************************************************************/

				yaml := xjy.Jstr2Y(json.V())
				// ioutil.WriteFile("xapi.yaml", []byte(yaml), 0666)
				xjy.YAMLScan(yaml, "id", pathDel, xjy.JSON, true, // *** skipDir must be <true>, otherwise dir version might be incorrect number ***
					func(p, v, id string) {
						defer func() { ver, cntV, prevID, prevTermID = ver+1, cntV+1, id, termID }()
						// fPln(p, v, id) // *** DEBUG
						if id != prevID {
							ver, termID = RequireVer(sqType, id)
							// fPln("Got:", ver, termID) // *** DEBUG
						}
						p = u.TerOp(addedRoot, defaultRoot+pathDel+p, p).(string)
						tuple := Must(messages.NewTuple(id, p, v)).(*pb.SPOTuple)
						tuple.Version = ver
						PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxXapi))
					})
				Terminate(sqType, prevID, prevTermID) //            *** object terminator ***
			}

		CHECKOVER:
			if _, _, _, v := q.Xapi(prevTermID, TERMMARK); v == nil || len(v) == 0 {
				time.Sleep(DELAY * time.Millisecond)
				goto CHECKOVER
			}

			lPln(fSf("<%06d> data tuples sent, <%06d> struct tuples sent, <%06d> array tuples sent\n", cntV, cntS, cntA))
		}
	}

	/********************************************************************/

	// defaultRoot := fSf("T%d", time.Now().Unix())
	// addedRoot := xjy.JSONModelInfo(xapi.V(), "id", defaultRoot, pathDel, childDel,
	// 	func(p, v string) {
	// 		defer func() { ver, cntS = ver+1, cntS+1 }()
	// 		tuple := Must(messages.NewTuple(p, "::", v)).(*pb.SPOTuple)
	// 		tuple.Version = ver
	// 		PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxXapi))
	// 	},
	// 	func(p, objID string, arrCnt int) {
	// 		defer func() { ver, cntA = ver+1, cntA+1 }()
	// 		tuple := Must(messages.NewTuple(p, objID, u.I32(arrCnt).ToStr())).(*pb.SPOTuple)
	// 		tuple.Version = ver
	// 		PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxXapi))
	// 	},
	// )

	/********************************************************************/

	// 	doneV, prevID, prevTermID := make(chan int), "", ""
	// 	yaml := xjy.Jstr2Y(xapi.V())
	// 	// ioutil.WriteFile("xapi.yaml", []byte(yaml), 0666)
	// 	go xjy.YAMLScanAsync(yaml, "id", pathDel, xjy.JSON, true, // *** skipDir must be <true>, otherwise dir version might be incorrect number ***
	// 		func(p, v, id string) {
	// 			defer func() { ver, cntV, prevID, prevTermID = ver+1, cntV+1, id, termID }()

	// 			fPln(p, v, id) // *** DEBUG

	// 			if id != prevID {
	// 				ver, termID = RequireVer(sqType, id)
	// 				fPln("Got:", ver, termID)
	// 				if prevID != "" {
	// 					Terminate(sqType, prevID, prevTermID)
	// 				}
	// 			}

	// 			p = u.TerOp(addedRoot, defaultRoot+pathDel+p, p).(string)
	// 			tuple := Must(messages.NewTuple(id, p, v)).(*pb.SPOTuple)
	// 			tuple.Version = ver
	// 			PE(g.N3clt.Publish(tuple, CFG.RPC.Namespace, CFG.RPC.CtxXapi))
	// 		}, doneV)
	// 	<-doneV

	// 	lPln(fSf("<%06d> data tuples sent, <%06d> struct tuples sent, <%06d> array tuples sent\n", cntV, cntS, cntA))

	// 	Terminate(sqType, prevID, prevTermID) // *** last object terminator ***
	// CHECK:
	// 	if _, _, _, v := q.Xapi(prevTermID, TERMMARK); v == nil || len(v) == 0 {
	// 		time.Sleep(DELAY * time.Millisecond)
	// 		goto CHECK
	// 	}

	return
}
