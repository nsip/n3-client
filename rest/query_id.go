package rest

import (
	"../query"
	u "github.com/cdutwhu/go-util"
)

func idListByPO(queryType string, mapParamPath map[string]string, mapParamValue map[string]interface{}) (IDs []string) {
	Fn := query.Sif
	switch queryType {
	case "sif", "Sif", "SIF":
		Fn = query.Sif
	case "xapi", "Xapi", "XAPI":
		Fn = query.Xapi
	}

	n := len(mapParamValue)
	idsList := make([][]string, n)
	for i := 0; i < n; i++ {
		idsList[i] = []string{}
	}

	idx := 0
	for param, value := range mapParamValue {
		s, _, _, _ := Fn("", mapParamPath[param], value.(string))
		for _, eachID := range s {
			idsList[idx] = append(idsList[idx], eachID)
		}
		idx++
	}

	if idx > 0 {
		IDs = idsList[0]
		for i := 1; i < idx; i++ {
			ga := u.Strs(IDs).ToG()
			gaNext := u.Strs(idsList[i]).ToG()
			IDs = u.ToGA(ga.InterSec(gaNext...)...).ToStrs()
		}
	}

	return
}

// IDsByPOFromSIF :
func IDsByPOFromSIF(mapParamPath map[string]string, mapParamValue map[string]interface{}) (IDs []string) {
	return idListByPO("sif", mapParamPath, mapParamValue)
}

// IDsByPOFromXAPI :
func IDsByPOFromXAPI(mapParamPath map[string]string, mapParamValue map[string]interface{}) (IDs []string) {
	return idListByPO("xapi", mapParamPath, mapParamValue)
}
