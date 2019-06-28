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
)

const (
	XML SQDType = iota
	JSON
	NLRU = 8192
)

var (
	N3clt *n3grpc.Client

	fPln = fmt.Println
	fSf  = fmt.Sprintf
	Must = u.Must

	OriExePath, _ = os.Getwd()

	LCRoot = Must(lru.NewWithEvict(NLRU, func(k, v interface{}) {
		fPln("Query Root onEvicted:", k, v)
	})).(*lru.Cache)

	LCSchema = Must(lru.NewWithEvict(NLRU, func(k, v interface{}) {
		fPln("Query Schema onEvicted:", k, v)
	})).(*lru.Cache)

	LCJSON = Must(lru.NewWithEvict(NLRU, func(k, v interface{}) {
		fPln("Query JSON onEvicted:", k, v)
	})).(*lru.Cache)

	MpQryRstRplc = map[string]string{
		`en-US`: `en_US`,
		// `"'`:    `"`,
		// `'"`:    `"`,
	}
)
