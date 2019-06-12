package send

import (
	"fmt"
	"log"
	"strings"

	c "../config"
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
	PE   = u.PanicOnError
	PE1  = u.PanicOnError1
	PC   = u.PanicOnCondition
	PH   = u.PanicHandle
	PHE  = u.PanicHandleEx
	LE   = u.LogOnError
	Must = u.Must
	IF   = u.IF

	XMLAttributes = gjxy.XMLAttributes
	XMLSegPos     = gjxy.XMLSegPos
	XMLSegsCount  = gjxy.XMLSegsCount
	XMLFamilyTree = gjxy.XMLFamilyTree
	XMLCntInfo    = gjxy.XMLCntInfo
	IsJSON        = gjxy.IsJSON

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf

	lPln = log.Println
	LPln = lPln

	sCtns = strings.Contains
	sJ    = strings.Join
	sCnt  = strings.Count

	e   error
	CFG *c.Config
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
