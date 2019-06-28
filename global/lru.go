package global

// RmIDsInLRU : Remove ObjectIDs from LRU Cache
func RmIDsInLRU(IDs ...string) {
	for _, ID := range IDs {
		LCSchema.Remove(ID)
		LCJSON.Remove(ID)
		LCRoot.Remove(ID)
	}
}
