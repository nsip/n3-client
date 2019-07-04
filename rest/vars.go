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
	Str  = w.Str
	Strs = w.Strs
)

var (
	pe            = u.PanicOnError
	pc            = u.PanicOnCondition
	phe           = u.PanicHandleEx
	must          = u.Must
	IF            = u.IF
	fPln          = fmt.Println
	fEf           = fmt.Errorf
	fSf           = fmt.Sprintf
	lPln          = log.Println
	sSpl          = strings.Split
	IArrIntersect = w.IArrIntersect
	IArrRmRep     = w.IArrRmRep
	IsJSON        = gjxy.IsJSON

	mtxQry = &sync.Mutex{}
	mtxPub = &sync.Mutex{}
	mtxDel = &sync.Mutex{}
	mtxID  = &sync.Mutex{}
	mtxObj = &sync.Mutex{}
	mtxScm = &sync.Mutex{}
)

const (
	BLANK = w.BLANK
)
