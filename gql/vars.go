package gql

import (
	"fmt"
	"log"
	"strings"

	c "../config"
	u "github.com/cdutwhu/go-util"
)

var (
	PE   = u.PanicOnError
	PE1  = u.PanicOnError1
	PC   = u.PanicOnCondition
	PH   = u.PanicHandle
	PHE  = u.PanicHandleEx
	LE   = u.LogOnError
	Must = u.Must

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf

	lPln = log.Println

	sSpl    = strings.Split
	sI      = strings.Index
	sLI     = strings.LastIndex
	sHP     = strings.HasPrefix
	sHS     = strings.HasSuffix
	sRepAll = strings.ReplaceAll

	e         error
	Cfg       *c.Config
	root      = ""
	mapStruct = map[string]string{}
	mapValue  = map[string][]*valver{}
	mapArray  = map[string]int{}
)

type valver struct {
	value string
	ver   int64
}
