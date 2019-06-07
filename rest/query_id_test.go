package rest

import "testing"

func TestQueryIDs(t *testing.T) {

	mapPP := map[string]string{
		"userId":    "lesson ~ userId",		
	}
	mapPV := map[string]interface{}{
		"userId":    "Angie",		
	}

	ids := IDsByPO(mapPP, mapPV)
	for _, id := range ids {
		fPln(id)
	}
}
