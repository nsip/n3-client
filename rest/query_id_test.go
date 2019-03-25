package rest

import "testing"

func TestQueryIDs(t *testing.T) {

	mapPP := map[string]string{
		"fname": "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ Name ~ FamilyName",
		"gname": "TeachingGroup ~ TeacherList ~ TeachingGroupTeacher ~ Name ~ GivenName",
		"rmNo":  "TeachingGroup ~ TeachingGroupPeriodList ~ TeachingGroupPeriod ~ RoomNumber",
	}
	mapPV := map[string]string{
		"fname": "Knoll",
		"gname": "Ina",
		"rmNo":  "141",
	}

	ids := IDsByPOFromSIF(mapPP, mapPV)
	for _, id := range ids {
		fPln(id)
	}
}
