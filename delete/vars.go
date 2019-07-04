package delete

import (
	"fmt"
	"log"

	u "github.com/cdutwhu/go-util"
)

var (
	pe   = u.PanicOnError
	pc   = u.PanicOnCondition
	ph   = u.PanicHandle
	IF   = u.IF
	fEf  = fmt.Errorf
	lPln = log.Println
)
