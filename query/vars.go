package query

import (
	"fmt"

	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	Ss = w.Strs
)

var (
	pc            = u.PanicOnCondition
	ph            = u.PanicHandle
	IF            = u.IF
	IArrIntersect = w.IArrIntersect
	fPln          = fmt.Println
	fPf           = fmt.Printf
	fEf           = fmt.Errorf
)
