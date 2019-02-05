package xjy

// Xstr2Y is
func Xstr2Y(xmlstr string) string {
	return Jstr2Y(Xstr2J(xmlstr))
}

// Xfile2Y is
func Xfile2Y(xmlfile string) string {
	//return string(Jb2Yb([]byte(Xfile2J(xmlfile))))
	return Jstr2Y(Xfile2J(xmlfile))
}
