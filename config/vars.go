package config

import (
	"fmt"

	u "github.com/cdutwhu/go-util"
)

var (
	pe   = u.PanicOnError
	ph   = u.PanicHandle
	must = u.Must
	fPln = fmt.Println
)
