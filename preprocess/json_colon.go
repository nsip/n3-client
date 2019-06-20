package preprocess

// HasColonInValue :
func HasColonInValue(str string) bool {
	return sCnt(str, "\":") != sCnt(str, ":")
}

// RplcColonInValueTo :
func RplcColonInValueTo(str, colonTo string) string {
	s := Str(str).Replace("\":", "#TAGEND#")
	s = s.Replace(":", colonTo)
	s = s.Replace("#TAGEND#", "\":")
	return s.V()
}
