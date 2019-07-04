package preprocess

import "strconv"

// HasColonInValue :
func HasColonInValue(str string) bool {
	return sCnt(str, "\":") != sCnt(str, ":")
}

// RplcValueColons :
func RplcValueColons(str string) string {
	colonTo := "^" + strconv.FormatInt((int64)(':'), 36) + "$"
	s := S(str).Replace("\":", "#TAGEND#")
	s = s.Replace(":", colonTo)
	s = s.Replace("#TAGEND#", "\":")
	return s.V()
}
