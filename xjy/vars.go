package xjy

import (
	"fmt"
	"strings"

	gjxy "github.com/cdutwhu/go-gjxy"
	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	S        = w.Str
	DataType int
)

const (
	DT_XML DataType = iota
	DT_JSON
)

var (
	ph   = u.PanicHandle
	must = u.Must
	IF   = u.IF
	sJ   = strings.Join
	fPln = fmt.Println
	fPf  = fmt.Printf

	JSONXPathValue        = gjxy.JSONXPathValue
	JSONWrapRoot          = gjxy.JSONWrapRoot
	JSONArrInfo           = gjxy.JSONArrInfo
	IsJSONArray           = gjxy.IsJSONArray
	IsJSONSingle          = gjxy.IsJSONSingle
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
	BLANK   = w.BLANK
)
