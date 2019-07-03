package xjy

import (
	"fmt"
	"strings"

	gjxy "github.com/cdutwhu/go-gjxy"
	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	Str      = w.Str
	I32s     = w.I32s
	DataType int
)

const (
	DT_XML DataType = iota
	DT_JSON
)

var (
	pe   = u.PanicOnError
	pe1  = u.PanicOnError1
	ph   = u.PanicHandle
	pc   = u.PanicOnCondition
	must = u.Must
	IF   = u.IF

	min = w.Min

	JSONXPathValue        = gjxy.JSONXPathValue
	JSONWrapRoot          = gjxy.JSONWrapRoot
	JSONArrInfo           = gjxy.JSONArrInfo
	IsJSONArray           = gjxy.IsJSONArray
	IsJSONSingle          = gjxy.IsJSONSingle
	XMLAttributes         = gjxy.XMLAttributes
	XMLSegPos             = gjxy.XMLSegPos
	XMLSegsCount          = gjxy.XMLSegsCount
	XMLFamilyTree         = gjxy.XMLFamilyTree
	XMLCntInfo            = gjxy.XMLCntInfo
	YAMLTag               = gjxy.YAMLTag
	YAMLValue             = gjxy.YAMLValue
	YAMLInfo              = gjxy.YAMLInfo
	YAMLJoinSplittedLines = gjxy.YAMLJoinSplittedLines
	Xstr2Y                = gjxy.Xstr2Y
	Jstr2Y                = gjxy.Jstr2Y

	sJ = strings.Join

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf
)

const (
	PATH_DEL  = " ~ "
	CHILD_DEL = " + "

	JT_NULL = gjxy.JT_NULL
	JT_OBJ  = gjxy.JT_OBJ
	JT_ARR  = gjxy.JT_ARR
	JT_STR  = gjxy.JT_STR
	JT_NUM  = gjxy.JT_NUM
	JT_BOOL = gjxy.JT_BOOL
	JT_MIX  = gjxy.JT_MIX
	JT_UNK  = gjxy.JT_UNK

	QDouble = w.QDouble
	ALL     = w.ALL
	LAST    = w.LAST
	BCurly  = w.BCurly
	BLANK   = w.BLANK
)
