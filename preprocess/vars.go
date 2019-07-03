package preprocess

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

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf

	lPln = log.Println
	LPln = lPln

	sJ   = strings.Join
	sCnt = strings.Count	
)

const (
	ALL     = w.ALL
	LAST    = w.LAST
	QSingle = w.QSingle
)
