package main

import (
	"fmt"

	g "../global"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	S = w.Str
)

var (
	fPln = fmt.Println

	lCtx  = len(g.Cfg.RPC.CtxList)
	ctxid = g.Cfg.RPC.CtxList[lCtx-1] //"ctxid"
)
