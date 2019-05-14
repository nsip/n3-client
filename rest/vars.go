package rest

import (
	"fmt"
	"log"
	"strings"

	c "../config"
	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
	gjxy "github.com/cdutwhu/go-gjxy"
)

type (
	Str  = w.Str
	Strs = w.Strs
)

var (
	PE   = u.PanicOnError
	PE1  = u.PanicOnError1
	PC   = u.PanicOnCondition
	PH   = u.PanicHandle
	PHE  = u.PanicHandleEx
	LE   = u.LogOnError
	Must = u.Must
	IF = u.IF

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf

	lPln = log.Println

	sCtn    = strings.Contains
	sRepAll = strings.ReplaceAll
	sJ      = strings.Join

	IArrIntersect = w.IArrIntersect
	IsJSON = gjxy.IsJSON

	e   error
	CFG *c.Config
	ver int64 = 1
)

const (
	TERMMARK  = "ENDENDEND"
	DELAY     = 2000
	PATH_DEL  = " ~ "
	CHILD_DEL = " + "
)
