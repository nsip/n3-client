package rest

import (
	q "../query"
)

// IDsByPO :
func IDsByPO(mParamPath map[string]string, mParamValue map[string]interface{}) (IDs []string) {

	// *** remove "" empty string value items from <mParamValue>
	for k, v := range mParamValue {
		if sv, ok := v.(string); ok && sv == "" {
			delete(mParamValue, k)
		}
	}
	// ***

	n := len(mParamValue)
	idsList := make([][]string, n)
	for i := 0; i < n; i++ {
		idsList[i] = []string{}
	}

	idx := 0
	for param, value := range mParamValue {
		s, _, _, _ := q.Data("", mParamPath[param], value.(string))
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

	if len(IDs) > 1 {
		IDs = IArrRmRep(Strs(IDs)).([]string)
	}
	return
}
