package xjy

import (
	"io/ioutil"
	"testing"

	c "../config"
	u "github.com/cdutwhu/go-util"
)

func TestYAMLScanAsync(t *testing.T) {
	cfg := c.FromFile("./config.toml", "../config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog) }()

	sifbytes := Must(ioutil.ReadFile("./files/sif.xml")).([]byte)
	sif := u.Str(sifbytes)
	sif.SetEnC()
	n := sif.XMLSegsCount()
	prevEnd := 0
	for i := 1; i <= n; i++ {
		nextStart := u.TerOp(i == 1, 0, prevEnd+1).(int)
		_, xml, _, end := sif.S(nextStart, u.ALL).XMLSegPos(1, 1)
		prevEnd = end + nextStart
		// if i == n {
		// 	fPln(tag)
		// 	fPln(xml)
		// }

		fPf("%d SIF*************************************************************\n", i)
		yamlstr, done := Xstr2Y(xml), make(chan int)
		// if i == 3 {
		// 	ioutil.WriteFile("sif3.yaml", []byte(yamlstr), 0666)
		// }

		idx := 0
		go YAMLScanAsync(yamlstr, "RefId", pathDel, XML, true, func(path, value, id string) {
			idx++
			// fPf("%06d : %s\n", idx, path)
			// fPf("%s\n", value)
			// fPf("%s\n", id)
			// fPln("----------------------------------------")
		}, done)

		fPf("finish: %d\n\n", <-done)
	}

	xapibytes := Must(ioutil.ReadFile("./files/xapi.json")).([]byte)
	xapi := u.Str(xapibytes)
	xapi.SetEnC()
	if ok, jsonType, n := xapi.IsJSONRootArray(); ok {
		if jsonType == "Object" {
			prevEnd := 0
			for i := 1; i <= n; i++ {
				nextStart := u.TerOp(i == 1, 0, prevEnd+1).(int)
				json, _, end := xapi.S(nextStart, u.ALL).BracketsPos(u.BCurly, 1, 1)
				prevEnd = end + nextStart
				// if i == n - 3 {
				// 	fPln(json)
				// }

				fPf("%d XAPI*************************************************************\n", i)
				yamlstr, done := Jstr2Y(json.V()), make(chan int)

				// if i == 1 {
				// 	ioutil.WriteFile("yaml1.yaml", []byte(yamlstr), 0666)
				// }

				idx := 0
				go YAMLScanAsync(yamlstr, "id", pathDel, JSON, true, func(path, value, id string) { // *** set "ID" & XML / JSON ***
					idx++
					fPf("%06d : %s\n", idx, path)
					fPf("%s\n", value)
					fPf("%s\n", id)
					fPln("-----------------------------------------")
				}, done)
				fPf("finish: %d\n\n", <-done)
			}
		}
	}

	//yamlstr, done := Xfile2Y("./files/sif.xml"), make(chan int)
	//ioutil.WriteFile(`./files/sif.yaml`, []byte(yamlstr), 0666)
	//yamlstr, done := Jfile2Y("./files/xapi.json"), make(chan int)
	//ioutil.WriteFile(`./files/xapi.yaml`, []byte(yamlstr), 0666)
	//yamlstr, done := Jfile2Y(`./files/content.json`), make(chan int)
	//ioutil.WriteFile(`./files/content.yaml`, []byte(yamlstr), 0666)

	// done := make(chan int)
	// yamlbytes, _ := ioutil.ReadFile("../temp.yaml")
	// yamlstr := string(yamlbytes)

	// idx := 0
	// go YAMLScanAsync(yamlstr, "RefId", pathDel, XML, true, func(path, value, id string) { // *** set "ID" & XML / JSON ***
	// 	idx++
	// 	fPf("%06d : %s\n", idx, path)
	// 	fPf("%s\n", value)
	// 	fPf("%s\n", id)
	// 	fPln("-----------------------------------------")
	// }, done)
	// fPf("finish: %d\n", <-done)

	//fbytes, err := ioutil.ReadFile("./files/nswdig.yaml")
	//PE(err)
}

func TestYAMLTag(t *testing.T) {
	fPln(YAMLTag(`- name: Andrew Downes`))
	fPln(YAMLTag(`actor:`))
	fPln(YAMLTag(`  mbox: mailto:teampb@example.com`))
	fPln(YAMLTag(`      homePage: http://www.example.com`))
	fPln(YAMLTag(`  - mbox_sha1sum: ebd31e95054c018b10727ccffd2ef2ec3a016ee9`))
	fPln(YAMLTag(`version: 1.0.0`))
	fPln(YAMLTag(`      - "9"`))
	fPln(YAMLTag(`- a`))
	fPln(YAMLTag(`-RefId: D3E34F41-9D75-101A-8C3D-00AA001A1652`))
}

func TestYAMLValue(t *testing.T) {
	fPln(YAMLValue(`- name: Andrew Downes`))
	fPln(YAMLValue(`actor:`))
	fPln(YAMLValue(`  mbox: mailto:teampb@example.com`))
	fPln(YAMLValue(`      homePage: http://www.example.com`))
	fPln(YAMLValue(`  - mbox_sha1sum: ebd31e95054c018b10727ccffd2ef2ec3a016ee9`))
	fPln(YAMLValue(`version: 1.0.0`))
	fPln(YAMLValue(`      - "9"`))
	fPln(YAMLValue(`- a`))
	fPln(YAMLValue(`-RefId: D3E34F41-9D75-101A-8C3D-00AA001A1652`))
}
