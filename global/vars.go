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
	// DataType : XML(0) or JSON(1)
	DataType int

	// QryIDs : ID-Query Cache Element
	QryIDs struct {
		Ctx string
		Qry map[string]interface{}
		IDs []string
	}
)

const (
	// XML : xml data type
	XML DataType = iota
	// JSON : json data type
	JSON
)

const (
	// DELIPath :
	DELIPath = " ~ "
	// DELIChild :
	DELIChild = " + "
	// MARKTerm :
	MARKTerm = "--------------------------------------"
	// MARKDead :
	MARKDead = "TOMBSTONE"
)

const (
	// NLRU : LRU Cache Capacity
	NLRU = 8192
	// NQryIDsCache : ID-Query Cache Capacity
	NQryIDsCache = 524288
)

var (
	fPln = fmt.Println
	fSf  = fmt.Sprintf
	must = u.Must

	// Cfg : Config File Struct
	Cfg *c.Config

	// N3clt : GRPC Client
	N3clt *n3grpc.Client

	// CurCtx : Current Context
	CurCtx = ""

	// OriExePath : Original Running Path
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

	// CacheQryIDs : Cache for ID-Query
	CacheQryIDs = make([]QryIDs, NQryIDsCache)

	// CacheQryIDsPtr : ID-Query Cache Current Pointer
	CacheQryIDsPtr = -1

	// MpQryRstRplc : Replacement Map for Query OutCome JSON
	MpQryRstRplc = map[string]string{
		`en-US`:      `en_US`,
		`": "true"`:  `": true`,
		`": "false"`: `": false`,
		`": "null"`:  `": null`,
	}
)
