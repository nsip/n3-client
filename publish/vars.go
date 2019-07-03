package publish

import (
	"fmt"
	"strings"

	xjy "../xjy"
	gjxy "github.com/cdutwhu/go-gjxy"
	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	Str  = w.Str
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

	UTF8ToASCII = w.UTF8ToASCII
	ASCIIToOri  = w.ASCIIToOri

	XMLAttributes = gjxy.XMLAttributes
	XMLSegPos     = gjxy.XMLSegPos
	XMLSegsCount  = gjxy.XMLSegsCount
	XMLFamilyTree = gjxy.XMLFamilyTree
	XMLCntInfo    = gjxy.XMLCntInfo
	IsJSON        = gjxy.IsJSON
	JSONWrapRoot  = gjxy.JSONWrapRoot

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf

	sCtns = strings.Contains
	sJ    = strings.Join
	sCnt  = strings.Count

	e error
	// ver int64 = 1
)

const (
	TERMMARK  = "--------------------------------------"
	DELAY     = 2000
	PATH_DEL  = " ~ "
	CHILD_DEL = " + "
	ALL       = w.ALL
	LAST      = w.LAST
	QSingle   = w.QSingle
	DT_XML    = xjy.DT_XML
	DT_JSON   = xjy.DT_JSON
)
