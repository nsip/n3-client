package rest

import "testing"

func TestQueryIDs(t *testing.T) {

	mapPP := map[string]string{
		"learning_area": "Content ~ learning_area",
		"subject":       "Content ~ subject",
		"stage":         "Content ~ stage",
	}
	mapPV := map[string]interface{}{
		"learning_area": "HSIE",
		"subject":       "Geography",
		"stage":         "1",
	}

	ids := IDsByPO(mapPP, mapPV, false)
	for _, id := range ids {
		fPln(id)
	}
}
