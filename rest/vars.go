package rest

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

	sCtn    = strings.Contains
	sRepAll = strings.ReplaceAll
	sJ      = strings.Join

	e   error
	CFG *c.Config
	ver int64 = 1
)

const (
	TERMMARK = "ENDENDEND"
	HEADTRIM = "sif."
	DELAY    = 2000
	pathDel  = " ~ "
	childDel = " + "
)
