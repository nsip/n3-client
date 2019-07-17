package global

// // RmIDsInLRU : Remove ObjectIDs from LRU Cache
// func RmIDsInLRU(IDs ...string) {
// 	for _, ID := range IDs {
// 		LCSchema.Remove(ID)
// 		LCJSON.Remove(ID)
// 		LCRoot.Remove(ID)
// 	}
// }

// // RmQryIDsCache : set IDs to nil to "remove" this
// func RmQryIDsCache(IDs ...string) {
// 	for _, ID := range IDs {
// 	OUT:
// 		for i, qry := range CacheQryIDs {
// 			if qry.IDs == nil {
// 				continue
// 			}
// 			for _, id := range qry.IDs {
// 				if id == ID {
// 					CacheQryIDs[i].IDs = nil
// 					continue OUT
// 				}
// 			}
// 		}
// 	}
// }

// // ClrAllIDsInLRU :
// func ClrAllIDsInLRU() {
// 	IDs := []string{}
// 	for _, id := range LCRoot.Keys() {
// 		IDs = append(IDs, id.(string))
// 	}
// 	if len(IDs) > 0 {
// 		IDs = IArrRmRep(Ss(IDs)).([]string)
// 		RmIDsInLRU(IDs...)
// 		RmQryIDsCache(IDs...)
// 	}
// }
