package filewatcher

import (
	"fmt"
	"log"

	u "github.com/cdutwhu/go-util"
)

var (
	fPln = fmt.Println
	lPln = log.Println
	pe   = u.PanicOnError
	ph   = u.PanicHandle
	must = u.Must
)
