package xjy

import (
	"io/ioutil"
	"strings"

	xj "github.com/basgys/goxml2json"
)

// Xstr2J is
func Xstr2J(xmlstr string) string {
	xml := strings.NewReader(xmlstr)
	json, err := xj.Convert(xml)
	PE1(err, "error on xj.Convert")
	return json.String()
}

// Xfile2J is
func Xfile2J(xmlfile string) string {
	xmlbytes, err := ioutil.ReadFile(xmlfile)
	PE1(err, "error on ioutil.ReadFile")
	return Xstr2J(string(xmlbytes))
}
