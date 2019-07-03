package global

import (
	"fmt"
	"os"

	c "../config"

	u "github.com/cdutwhu/go-util"
	lru "github.com/hashicorp/golang-lru"
	"github.com/nsip/n3-messages/n3grpc"
)

type (
	SQDType int

	QryIDs struct {
		Ctx string
		Qry map[string]interface{}
		IDs []string
	}
)

const (
	XML SQDType = iota
	JSON
	NLRU         = 8192
	NQryIDsCache = 524288
)

var (
	fPln = fmt.Println
	fSf  = fmt.Sprintf
	must = u.Must

	Cfg *c.Config
	N3clt *n3grpc.Client

	CurCtx        = "demo"
	OriExePath, _ = os.Getwd()

	// LCRoot *** ID : Root *** ID query cache
	LCRoot = must(lru.NewWithEvict(NLRU, func(k, v interface{}) {
		fPln("Query Root onEvicted:", k, v)
	})).(*lru.Cache)

	// LCSchema *** ID : Schema *** Schema query cache
	LCSchema = must(lru.NewWithEvict(NLRU, func(k, v interface{}) {
		fPln("Query Schema onEvicted:", k, v)
	})).(*lru.Cache)

	// LCJSON *** ID : JSON *** JSON query cache
	LCJSON = must(lru.NewWithEvict(NLRU, func(k, v interface{}) {
		fPln("Query JSON onEvicted:", k, v)
	})).(*lru.Cache)

	CacheQryIDs    = make([]QryIDs, NQryIDsCache)
	CacheQryIDsPtr = -1

	MpQryRstRplc = map[string]string{
		`en-US`:      `en_US`,
		`": "true"`:  `": true`,
		`": "false"`: `": false`,
		`": "null"`:  `": null`,
	}
)
