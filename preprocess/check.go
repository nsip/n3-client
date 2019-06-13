package preprocess

func hasColonInValue(str string) bool {
	return sCnt(str, "\":") != sCnt(str, ":")
}

func hasSQuoteInValue(str string) bool {
	return Str(str).Contains("'")	
}

// func hasHyphenInTag(str string) bool {
// 	return sCtns(str, "-")
// }
