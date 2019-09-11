package rest

import (
	"fmt"
	"log"
	"strings"
	"sync"

	gjxy "github.com/cdutwhu/go-gjxy"
	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	// S is
	S = w.Str
	// Ss is
	Ss = w.Strs
)

var (
	// IF is
	IF   = u.IF
	pe   = u.PanicOnError
	pc   = u.PanicOnCondition
	phe  = u.PanicHandleEx
	must = u.Must
	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf
	lPln = log.Println
	sSpl = strings.Split
	// IArrIntersect is
	IArrIntersect = w.IArrIntersect
	// GetMapKVs is
	GetMapKVs = w.GetMapKVs
	// SlcD2ToD1 is
	SlcD2ToD1 = w.SlcD2ToD1
	// IArrRmRep is
	IArrRmRep = w.IArrRmRep
	// IsJSON is
	IsJSON = gjxy.IsJSON
	// Get1stObjInQry is
	Get1stObjInQry = gjxy.Get1stObjInQry

	mtxQry  = &sync.Mutex{}
	mtxQry2 = &sync.Mutex{}
	mtxPub  = &sync.Mutex{}
	mtxDel  = &sync.Mutex{}
	mtxID   = &sync.Mutex{}
	mtxObj  = &sync.Mutex{}
	mtxScm  = &sync.Mutex{}
)

const (
	// BLANK is
	BLANK = w.BLANK
)
