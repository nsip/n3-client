package query

import (
	"fmt"
	"log"

	c "../config"
	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	Str  = w.Str
	Strs = w.Strs
)

var (
	PE         = u.PanicOnError
	PE1        = u.PanicOnError1
	PC         = u.PanicOnCondition
	PH         = u.PanicHandle
	PHE        = u.PanicHandleEx
	LE         = u.LogOnError
	Must       = u.Must
	IF         = u.IF	

	IArrIntersect = w.IArrIntersect

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf

	lPln = log.Println

	e   error
	CFG *c.Config
)
