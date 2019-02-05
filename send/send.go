package send

import (
	"time"

	"github.com/nsip/n3-messages/messages/pb"

	c "../config"
	g "../global"
	q "../query"
	"../xjy"
	u "github.com/cdutwhu/go-util"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/n3grpc"
)

// Junk :
func Junk(n int) {
	PC(Cfg == nil || g.N3pub == nil, fEf("Missing Init, do 'Init(&config) before sending'\n"))
	for i := 0; i < n; i++ {
		tuple := Must(messages.NewTuple("ab", "pre", "obj")).(*pb.SPOTuple)
		tuple.Version = int64(i)
		PE(g.N3pub.Publish(tuple, Cfg.RPC.Namespace, Cfg.RPC.CtxSif))
	}
}

// Terminate :
func Terminate(t g.SQType, objID, termID string) {
	defer func() { ver++ }()
	if Cfg == nil || g.N3pub == nil {
		Cfg = c.GetConfig("./config.toml", "../config/config.toml")
		Init(Cfg)
	}

	tuple := Must(messages.NewTuple(termID, TERMMARK, objID)).(*pb.SPOTuple)
	tuple.Version = ver
	ctx := u.CaseAssign(t, g.SIF, g.XAPI, Cfg.RPC.CtxSif, Cfg.RPC.CtxXapi).(string)
	PE(g.N3pub.Publish(tuple, Cfg.RPC.Namespace, ctx))
}

// RequireVer :
func RequireVer(t g.SQType, objID string) (ver int64, termID string) {
	if Cfg == nil || g.N3pub == nil {
		Cfg = c.GetConfig("./config.toml", "../config/config.toml")
		Init(Cfg)
	}
	_, p, o, _ := q.Meta(t, objID, "V")
	PC(len(p) == 0, fEf("Got Version Error"))
	ver, termID = u.Str(o[0]).ToInt64()+1, p[0]
	return
}

/************************************************************/

// Init :
func Init(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	Cfg = config
	if g.N3pub == nil {
		g.N3pub = Must(n3grpc.NewPublisher(Cfg.RPC.Server, Cfg.RPC.Port)).(*n3grpc.Publisher)
	}
}

// Sif is
func Sif(str string) (cntV, cntS, cntA int, termID string) {
	PC(Cfg == nil || g.N3pub == nil, fEf("Missing Send Init, do 'Init(&config) before sending'\n"))

	content, sqType := u.Str(str), g.SIF
	PC(content.L() == 0 || !content.IsXMLSegSimple(), fEf("Incoming string is invalid xml segment\n"))

	xjy.XMLModelInfo(content.V(), "RefId", true,
		func(p, v string) {
			defer func() { ver, cntS = ver+1, cntS+1 }()
			tuple := Must(messages.NewTuple(u.Str(p).RmPrefix(HEADTRIM), "::", v)).(*pb.SPOTuple)
			tuple.Version = ver
			PE(g.N3pub.Publish(tuple, Cfg.RPC.Namespace, Cfg.RPC.CtxSif))
		},
		func(p, objID string, arrCnt int) {
			defer func() { ver, cntA = ver+1, cntA+1 }()
			tuple := Must(messages.NewTuple(u.Str(p).RmPrefix(HEADTRIM), objID, u.I32(arrCnt).ToStr())).(*pb.SPOTuple)
			tuple.Version = ver
			PE(g.N3pub.Publish(tuple, Cfg.RPC.Namespace, Cfg.RPC.CtxSif))
		},
	)

	doneV, prevID, prevTermID := make(chan int), "", ""
	go xjy.YAMLScanAsync(xjy.Xstr2Y(content.V()), "RefId", xjy.XML, true,
		func(p, v, id string) {
			defer func() { ver, cntV, prevID, prevTermID = ver+1, cntV+1, id, termID }()
			if id != prevID {
				ver, termID = RequireVer(sqType, id)
				fPln("Got:", ver, termID)
				if prevID != "" {
					Terminate(sqType, prevID, prevTermID)
				}
			}
			tuple := Must(messages.NewTuple(id, u.Str(p).RmPrefix(HEADTRIM), v)).(*pb.SPOTuple)
			tuple.Version = ver
			PE(g.N3pub.Publish(tuple, Cfg.RPC.Namespace, Cfg.RPC.CtxSif))
		},
		doneV)
	<-doneV

	lPln(fSf("<%06d> data tuples sent, <%06d> struct tuples sent, <%06d> array tuples sent\n", cntV, cntS, cntA))

	Terminate(sqType, prevID, prevTermID) // *** last object terminator ***
CHECK:
	if _, _, _, v := q.Sif(prevTermID, TERMMARK); v == nil || len(v) == 0 {
		time.Sleep(DELAY * time.Millisecond)
		goto CHECK
	}
	return
}

// Xapi is
func Xapi(str string) (cnt int, termID string) {
	PC(Cfg == nil, fEf("Missing Send Init, do 'Init(&config) before sending'\n"))

	content, sqType := u.Str(str), g.XAPI
	PC(content.L() == 0 || !content.IsJSON(), fEf("Incoming string is invalid json\n"))

	done, prevID, prevTermID := make(chan int), "", ""
	go xjy.YAMLScanAsync(xjy.Jstr2Y(content.V()), "id", xjy.JSON, true, func(p, v, id string) {
		defer func() { ver, cnt, prevID, prevTermID = ver+1, cnt+1, id, termID }()
		if id != prevID {
			ver, termID = RequireVer(sqType, id)
			fPln("Got:", ver, termID)
			if prevID != "" {
				Terminate(sqType, prevID, prevTermID)
			}
		}
		tuple := Must(messages.NewTuple(id, p, v)).(*pb.SPOTuple)
		tuple.Version = ver
		PE(g.N3pub.Publish(tuple, Cfg.RPC.Namespace, Cfg.RPC.CtxXapi))
	}, done)
	<-done

	lPln(fSf("<%06d> tuples sent\n", cnt))

	Terminate(sqType, prevID, prevTermID) // *** last object terminator ***
CHECK:
	if _, _, _, v := q.Xapi(prevTermID, TERMMARK); v == nil || len(v) == 0 {
		time.Sleep(DELAY * time.Millisecond)
		goto CHECK
	}
	return
}
