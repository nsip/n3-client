package rest

import (
	"fmt"
	"log"
	"strings"

	c "../config"
	gjxy "github.com/cdutwhu/go-gjxy"
	u "github.com/cdutwhu/go-util"
	w "github.com/cdutwhu/go-wrappers"
)

type (
	Str  = w.Str
	Strs = w.Strs
)

var (
	PE   = u.PanicOnError
	PE1  = u.PanicOnError1
	PC   = u.PanicOnCondition
	PH   = u.PanicHandle
	PHE  = u.PanicHandleEx
	LE   = u.LogOnError
	Must = u.Must
	IF   = u.IF

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf

	lPln = log.Println

	sCtn    = strings.Contains
	sRepAll = strings.ReplaceAll
	sJ      = strings.Join
	sSpl    = strings.Split

	IArrIntersect = w.IArrIntersect
	IsJSON        = gjxy.IsJSON

	e   error
	CFG *c.Config
	ver int64 = 1
)

const (
	TERMMARK  = "--------------------------------------"
	DELAY     = 2000
	PATH_DEL  = " ~ "
	CHILD_DEL = " + "
)

// var (
// 	mPP = map[string]string{
// 		"fname":           "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ Name ~ FamilyName",
// 		"gname":           "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ Name ~ GivenName",
// 		"staffid":         "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ StaffPersonalRefId",
// 		"teachinggroupid": "TeachingGroup ~ -RefId",
// 		"tgid":            "GradingAssignment ~ TeachingGroupRefId",
// 		"studentid":       "StudentAttendanceTimeList ~ StudentPersonalRefId",
// 		"objid":           "xapi ~ object ~ id",
// 	}
// )
