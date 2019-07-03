package gql

import (
	"fmt"
	"log"
	"strings"

	gjxy "github.com/cdutwhu/go-gjxy"
	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	Str  = w.Str
	Strs = w.Strs
	I32  = w.I32
	I32s = w.I32s
)

var (
	pe   = u.PanicOnError
	pe1  = u.PanicOnError1
	pc   = u.PanicOnCondition
	ph   = u.PanicHandle
	phe  = u.PanicHandleEx
	le   = u.LogOnError
	must = u.Must
	IF   = u.IF

	Max           = w.Max
	Min           = w.Min
	GetMapKeys    = w.GetMapKeys
	IArrSeqCtns   = w.IArrSeqCtns
	IArrStrJoinEx = w.IArrStrJoinEx
	SortIntArr2D  = w.SortIntArr2D
	IArrMake      = w.IArrMake
	ASCIIToOri    = w.ASCIIToOri

	JSONWrapRoot    = gjxy.JSONWrapRoot
	JSONMake        = gjxy.JSONMakeIPath
	JSONMakeRep     = gjxy.JSONMakeIPathRep
	JSONObjectMerge = gjxy.JSONObjectMerge
	SchemaMake      = gjxy.SchemaMake

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf

	lPln = log.Println

	sSpl    = strings.Split
	sRepAll = strings.ReplaceAll
	sJ      = strings.Join
	sCnt    = strings.Count
	
	root            = ""
	mStruct         = map[string]string{}
	mValue          = map[string][]*valver{}
	mArray          = map[string]int{}
	mIndicesList    = map[string][][]int{}  // *** FOR JSONBuilld ***
	mIPathObj       = map[string]string{}   // *** FOR JSONBuilld ***
	mIPathSubIPaths = map[string][]string{} // *** FOR JSONBuilld ***
)

type valver struct {
	value string
	ver   int64
}

const (
	PATH_DEL  = " ~ "
	CHILD_DEL = " + "
	BLANK     = w.BLANK
)
