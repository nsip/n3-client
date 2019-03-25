package main

import (
	"fmt"
	"strings"

	u "github.com/cdutwhu/go-util"
)

var (
	PE  = u.PanicOnError
	PE1 = u.PanicOnError1
	PH  = u.PanicHandle
	PC  = u.PanicOnCondition

	sFF = strings.FieldsFunc
	sC  = strings.Contains
	sJ  = strings.Join

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf
)
