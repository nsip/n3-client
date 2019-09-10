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
	S  = w.Str
	Ss = w.Strs
)

var (
	pe              = u.PanicOnError
	pc              = u.PanicOnCondition
	ph              = u.PanicHandle
	must            = u.Must
	IF              = u.IF
	trueAssign      = u.TrueAssign
	GetMapKeys      = w.GetMapKeys
	IArrStrJoinEx   = w.IArrStrJoinEx
	SortIntArr2D    = w.SortIntArr2D
	IArrMake        = w.IArrMake
	ASCIIToOri      = w.ASCIIToOri
	JSONWrapRoot    = gjxy.JSONWrapRoot
	JSONMake        = gjxy.JSONMakeIPath
	JSONMakeRep     = gjxy.JSONMakeIPathRep
	JSONObjectMerge = gjxy.JSONObjectMerge
	SchemaMake      = gjxy.SchemaMake
	fPln            = fmt.Println
	fPf             = fmt.Printf
	fEf             = fmt.Errorf
	fSf             = fmt.Sprintf
	lPln            = log.Println
	sSpl            = strings.Split
	sRepAll         = strings.ReplaceAll
	sJ              = strings.Join

	root            = ""
	mStruct         = map[string]string{}
	mValue          = map[string][]*valver{}
	mArray          = map[string]int{}
	mIndicesList    = map[string][][]int{}  // *** FOR JSONBuilld ***
	mIPathObj       = map[string]string{}   // *** FOR JSONBuilld ***
	mIPathSubIPaths = map[string][]string{} // *** FOR JSONBuilld ***
)

const (
	DELIPath   = " ~ "
	DELISchema = "###"
)

type valver struct {
	value string
	ver   int64
}
