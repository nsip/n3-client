package rest

import (
	"fmt"
	"log"
	"strings"
	"sync"

	c "../config"
	gjxy "github.com/cdutwhu/go-gjxy"
	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	Str  = w.Str
	Strs = w.Strs
)

var (
	PE   = u.PanicOnError
	PE1  = u.PanicOnError1
	PC   = u.PanicOnCondition
	PH   = u.PanicHandle
	PHE  = u.PanicHandleEx
	LE   = u.LogOnError
	Must = u.Must
	IF   = u.IF

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf

	lPln = log.Println

	sCtns   = strings.Contains
	sRepAll = strings.ReplaceAll
	sJ      = strings.Join
	sSpl    = strings.Split

	IArrIntersect = w.IArrIntersect
	IArrRmRep     = w.IArrRmRep
	ASCIIToOri    = w.ASCIIToOri
	UTF8ToASCII   = w.UTF8ToASCII

	IsJSON = gjxy.IsJSON

	e   error
	CFG *c.Config
	ver int64 = 1

	mtxQry    = &sync.Mutex{}
	mtxPub    = &sync.Mutex{}
	mtxDel    = &sync.Mutex{}
	mtxID     = &sync.Mutex{}
	mtxObj    = &sync.Mutex{}
	mtxScm    = &sync.Mutex{}
	mtxQryTxt = &sync.Mutex{}
)

const (
	TERMMARK  = "--------------------------------------"
	DELAY     = 2000
	PATH_DEL  = " ~ "
	CHILD_DEL = " + "
	BLANK     = w.BLANK
)
