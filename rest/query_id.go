package rest

import (
	q "../query"
)

// IDsByPO :
func IDsByPO(mapParamPath map[string]string, mapParamValue map[string]interface{}) (IDs []string) {

	n := len(mapParamValue)
	idsList := make([][]string, n)
	for i := 0; i < n; i++ {
		idsList[i] = []string{}
	}

	idx := 0
	for param, value := range mapParamValue {
		s, _, _, _ := q.Data("", mapParamPath[param], value.(string))
		for _, eachID := range s {
			idsList[idx] = append(idsList[idx], eachID)
		}
		idx++
	}

	if idx > 0 {
		IDs = idsList[0]
		for i := 1; i < idx; i++ {
			if len(IDs) > 0 {
				if rst := IArrIntersect(Strs(IDs), Strs(idsList[i])); rst != nil {
					IDs = rst.([]string)
				} else {
					IDs = []string{}
				}
			}
		}
	}

	return
}
