package send

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
	FPln = fPln
	fPf  = fmt.Printf
	FPf  = fPf
	fEf  = fmt.Errorf
	FEf  = fEf
	fSf  = fmt.Sprintf
	FSf  = fSf

	lPln = log.Println
	LPln = lPln

	sC = strings.Contains
	SC = sC

	e   error
	Cfg *c.Config
	ver int64 = 1
)

const (
	TERMMARK = "ENDENDEND"
	HEADTRIM = "sif."
	DELAY    = 300
)
