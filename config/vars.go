package config

import (
	"fmt"
	"log"

	u "github.com/cdutwhu/go-util"
)

var (
	pe   = u.PanicOnError	
	ph   = u.PanicHandle	
	must = u.Must

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf
	lPln = log.Println
)
