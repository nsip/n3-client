package xjy

import (
	"fmt"
	"strings"

	gjxy "github.com/cdutwhu/go-gjxy"
	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	S  = w.Str
	Ss = w.Strs
)

var (
	sJ = strings.Join

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf

	ph   = u.PanicHandle
	pc   = u.PanicOnCondition
	must = u.Must
	IF   = u.IF

	getMapKeys  = w.GetMapKeys
	IArrEleIn   = w.IArrEleIn
	UTF8ToASCII = w.UTF8ToASCII

	JSONChildValue        = gjxy.JSONChildValue
	JSONXPathValue        = gjxy.JSONXPathValue
	JSONWrapRoot          = gjxy.JSONWrapRoot
	JSONArrInfo           = gjxy.JSONArrInfo
	JSONObjChildren       = gjxy.JSONObjChildren
	IsJSONArr             = gjxy.IsJSONArr
	IsJSONArrOnFmtL0      = gjxy.IsJSONArrOnFmtL0
	IsJSONSingle          = gjxy.IsJSONSingle
	IsJSON                = gjxy.IsJSON
	XMLAttributes         = gjxy.XMLAttributes
	XMLSegPos             = gjxy.XMLSegPos
	XMLSegsCount          = gjxy.XMLSegsCount
	XMLCntInfo            = gjxy.XMLCntInfo
	YAMLTag               = gjxy.YAMLTag
	YAMLValue             = gjxy.YAMLValue
	YAMLInfo              = gjxy.YAMLInfo
	YAMLJoinSplittedLines = gjxy.YAMLJoinSplittedLines
	Xstr2Y                = gjxy.Xstr2Y
	Jstr2Y                = gjxy.Jstr2Y
)

const (
	J_NULL = gjxy.J_NULL
	J_OBJ  = gjxy.J_OBJ
	J_ARR  = gjxy.J_ARR
	J_STR  = gjxy.J_STR
	J_NUM  = gjxy.J_NUM
	J_BOOL = gjxy.J_BOOL
	J_MIX  = gjxy.J_MIX
	J_UNK  = gjxy.J_UNK

	QDouble = w.QDouble
	ALL     = w.ALL
	BLANK   = w.BLANK
)
