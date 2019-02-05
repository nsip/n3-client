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

	sI  = strings.Index
	sLI = strings.LastIndex
	sT  = strings.Trim
	sTL = strings.TrimLeft
	sTR = strings.TrimRight
	sHP = strings.HasPrefix
	sHS = strings.HasSuffix
	sFF = strings.FieldsFunc
	sC  = strings.Contains
	sJ  = strings.Join

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf
)
