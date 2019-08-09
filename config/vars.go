package config

import (
	"fmt"

	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	S = w.Str
)

var (
	pe   = u.PanicOnError
	ph   = u.PanicHandle
	must = u.Must
	fPln = fmt.Println
	fSf  = fmt.Sprintf
)
