package rest

import "testing"

func TestQueryIDs(t *testing.T) {

	mapPP := map[string]string{
		"ccode":    "SchoolCourseInfo ~ CourseCode",
		"ccredits": "SchoolCourseInfo ~ CourseCredits",
		"ctitle":  "SchoolCourseInfo ~ CourseTitle",
	}
	mapPV := map[string]interface{}{
		"ccode":    "Mathematics 701",
		"ccredits": "2",
		"ctitle":  "Mathematics 7",
	}

	ids := IDsByPO(mapPP, mapPV)
	for _, id := range ids {
		fPln(id)
	}
}
