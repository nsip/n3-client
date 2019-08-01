package rest

import "testing"
import g "github.com/nsip/n3-client/global"

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

	ids := IDsByPO(g.CurCtx, mapPP, mapPV, false)
	for _, id := range ids {
		fPln(id)
	}
}
