package global

import (
	"fmt"
	"os"

	u "github.com/cdutwhu/go-util"
	lru "github.com/hashicorp/golang-lru"
	"github.com/nsip/n3-messages/n3grpc"
)

type (
	SQDType int

	QryIDs struct {
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
	N3clt *n3grpc.Client

	fPln = fmt.Println
	fSf  = fmt.Sprintf
	Must = u.Must

	OriExePath, _ = os.Getwd()

	// LCRoot *** ID : Root *** ID query cache
	LCRoot = Must(lru.NewWithEvict(NLRU, func(k, v interface{}) {
		fPln("Query Root onEvicted:", k, v)
	})).(*lru.Cache)

	// LCSchema *** ID : Schema *** Schema query cache
	LCSchema = Must(lru.NewWithEvict(NLRU, func(k, v interface{}) {
		fPln("Query Schema onEvicted:", k, v)
	})).(*lru.Cache)

	// LCJSON *** ID : JSON *** JSON query cache
	LCJSON = Must(lru.NewWithEvict(NLRU, func(k, v interface{}) {
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
