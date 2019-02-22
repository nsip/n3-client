package xjy

import (
	"io/ioutil"
	"testing"

	c "../config"
)

func TestYAMLScanAsync(t *testing.T) {
	cfg := c.GetConfig("./config.toml", "../config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog) }()

	yamlstr, done := Xfile2Y("./files/nswdig.xml"), make(chan int)
	ioutil.WriteFile(`./files/nswdig.yaml`, []byte(yamlstr), 0666)
	//yamlstr, done := Xfile2Y("./files/staffpersonal.xml"), make(chan int)
	//ioutil.WriteFile(`./files/staffpersonal.yaml`, []byte(yamlstr), 0666)
	//yamlstr, done := Jfile2Y(`./files/xapifile.json`), make(chan int)
	//ioutil.WriteFile(`./files/xapifile.yaml`, []byte(yamlstr), 0666)

	// done := make(chan int)
	// yamlbytes, _ := ioutil.ReadFile("../temp.yaml")
	// yamlstr := string(yamlbytes)

	idx := 0
	go YAMLScanAsync(yamlstr, "RefId", XML, true, func(path, value, id string) {
		idx++
		fPf("%06d : %s\n", idx, path)
		fPf("%s\n", value)
		fPf("%s\n", id)
		fPln("-----------------------------------------")
	}, done)
	fPf("finish: %d\n", <-done)

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
