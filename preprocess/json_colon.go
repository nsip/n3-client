package preprocess

// HasColonInValue :
func HasColonInValue(str string) bool {
	return sCnt(str, "\":") != sCnt(str, ":")
}

// RplcColonInValue :
func RplcColonInValue(str, colonTo string) string {
	s := Str(str).Replace("\":", "#TAGEND#")
	s = s.Replace(":", colonTo)
	s = s.Replace("#TAGEND#", "\":")
	return s.V()
}
