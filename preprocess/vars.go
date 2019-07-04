package preprocess

import (
	"fmt"
	"log"
	"strings"

	gjxy "github.com/cdutwhu/go-gjxy"
	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	S = w.Str
)

var (
	pe          = u.PanicOnError
	pc          = u.PanicOnCondition
	must        = u.Must
	UTF8ToASCII = w.UTF8ToASCII
	ASCIIToOri  = w.ASCIIToOri
	IsJSON      = gjxy.IsJSON
	fPln        = fmt.Println
	fEf         = fmt.Errorf
	lPln        = log.Println
	sCnt        = strings.Count
)

const (
	QSingle = w.QSingle
)
