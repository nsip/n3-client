package publish

import (
	"io/ioutil"

	pp "github.com/nsip/n3-client/preprocess"
)

func prepJSON(json string) string {

	// *** format json ***

	// json = pp.FmtJSONStr(json, "../preprocess/util/", "./")
	ioutil.WriteFile("../build/debug_pub/in.json", []byte(json), 0666)
	json = pp.FmtJSONFile("../../build/debug_pub/in.json", "../preprocess/util/", "./")
	ioutil.WriteFile("../build/debug_pub/infmt.json", []byte(json), 0666)

	// *** ': null' => ': "null"' ***
	json = S(json).Replace(`": null`, `": "null"`).V()

	// *** dealing with colon ***
	if pp.HasColonInValue(json) {
		json = pp.RplcValueColons(json)
	}

	// *** convert to ASCII ***
	if ascii, ajson := UTF8ToASCII(json); !ascii {
		fPln("is utf8")
		return ajson
	}
	return json
}

func prepXML(xml string) string {

	// *** convert to ASCII ***
	if ascii, axml := UTF8ToASCII(xml); !ascii {
		fPln("is utf8")
		return axml
	}
	return xml
}
