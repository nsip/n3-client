package global

import (
	"fmt"

	u "github.com/cdutwhu/go-util"
	lru "github.com/hashicorp/golang-lru"
	"github.com/nsip/n3-messages/n3grpc"
)

var (
	N3clt *n3grpc.Client

	fPln = fmt.Println
	Must = u.Must

	LCRoot = Must(lru.NewWithEvict(LRUCOUNT, func(k, v interface{}) {
		fPln("Query Root onEvicted:", k, v)
	})).(*lru.Cache)

	LCSchema = Must(lru.NewWithEvict(LRUCOUNT, func(k, v interface{}) {
		fPln("Query Schema onEvicted:", k, v)
	})).(*lru.Cache)

	LCJSON = Must(lru.NewWithEvict(LRUCOUNT, func(k, v interface{}) {
		fPln("Query JSON onEvicted:", k, v)
	})).(*lru.Cache)

	LCQryID = Must(lru.NewWithEvict(LRUCOUNT, func(k, v interface{}) {
		fPln("Query-IDs onEvicted:", k, v)
	})).(*lru.Cache)
)

type (
	SQDType int
)

const (
	XML SQDType = iota
	JSON
	LRUCOUNT = 4096
)
