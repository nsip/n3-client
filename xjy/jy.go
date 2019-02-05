package xjy

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Jstr2Y is
func Jstr2Y(jsonstr string) string {
	yamlbytes, err := yaml.JSONToYAML([]byte(jsonstr))
	PE1(err, "error on yaml.JSONToYAML")
	//return string(yamlbytes)
	return YAMLRmHangStr(string(yamlbytes)) /* avoid hanging string value */
}

// Jb2Yb is
func Jb2Yb(jsonbytes []byte) []byte {
	yamlbytes, err := yaml.JSONToYAML(jsonbytes)
	PE1(err, "error on yaml.JSONToYAML")
	return yamlbytes
}

// Jfile2Y is
func Jfile2Y(jsonfile string) string {
	jsonbytes, err := ioutil.ReadFile(jsonfile)
	PE1(err, "error on ioutil.ReadFile")
	return Jstr2Y(string(jsonbytes))
}

// Jfile2Yb is
func Jfile2Yb(jsonfile string) []byte {
	jsonbytes, err := ioutil.ReadFile(jsonfile)
	PE1(err, "error on ioutil.ReadFile")
	return Jb2Yb(jsonbytes)
}
