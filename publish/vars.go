package publish

import (
	"fmt"
	"strings"

	gjxy "github.com/cdutwhu/go-gjxy"
	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	S = w.Str
)

var (
	pe           = u.PanicOnError
	pc           = u.PanicOnCondition
	ph           = u.PanicHandle
	must         = u.Must
	IF           = u.IF
	UTF8ToASCII  = w.UTF8ToASCII
	IsJSON       = gjxy.IsJSON
	JSONWrapRoot = gjxy.JSONWrapRoot
	fPln         = fmt.Println
	fEf          = fmt.Errorf
	fSf          = fmt.Sprintf
	sJ           = strings.Join
)
