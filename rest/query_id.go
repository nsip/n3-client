package rest

import (
	"reflect"

	g "../global"
	q "../query"
)

// GetIDs :
func GetIDs(ctx string, mParamPath map[string]string, mParamValue map[string]interface{}, all bool) []string {
	for _, qry := range g.CacheQryIDs {
		if reflect.DeepEqual(qry.Qry, mParamValue) && qry.IDs != nil {
			return qry.IDs
		}
	}
	IDs := IDsByPO(ctx, mParamPath, mParamValue, all)
	if len(IDs) > 0 && IDs[0] != "" {
		g.CacheQryIDsPtr++
		g.CacheQryIDsPtr %= g.NQryIDsCache
		g.CacheQryIDs[g.CacheQryIDsPtr] = g.QryIDs{Ctx: g.CurCtx, Qry: mParamValue, IDs: IDs}
	}
	return IDs
}

// IDsByPO :
func IDsByPO(ctx string, mParamPath map[string]string, mParamValue map[string]interface{}, all bool) (IDs []string) {

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
		s, _, _, _ := q.Data(ctx, IF(all, "*", "").(string), mParamPath[param], value.(string))
		for _, eachID := range s {
			idsList[idx] = append(idsList[idx], eachID)
		}
		idx++
	}

	if idx > 0 {
		IDs = idsList[0]
		for i := 1; i < idx; i++ {
			if len(IDs) > 0 {
				if rst := IArrIntersect(Ss(IDs), Ss(idsList[i])); rst != nil {
					IDs = rst.([]string)
				} else {
					IDs = []string{}
				}
			}
		}
	}

	if len(IDs) > 1 {
		IDs = IArrRmRep(Ss(IDs)).([]string)
	}
	return
}
