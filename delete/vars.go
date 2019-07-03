package delete

import (
	"fmt"
	"log"

	u "github.com/cdutwhu/go-util"
)

var (
	pe = u.PanicOnError
	pc = u.PanicOnCondition
	ph = u.PanicHandle
	IF = u.IF

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf

	lPln = log.Println
)

const (
	DEADMARK = "TOMBSTONE"
)
