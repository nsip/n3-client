package xjy

import (
	"io/ioutil"
	"testing"

	c "github.com/nsip/n3-client/config"
	g "github.com/nsip/n3-client/global"
)

func TestYAMLScan(t *testing.T) {
	cfg := c.FromFile("../build/config.toml")
	defer func() { ph(recover(), cfg.ErrLog) }()

	databytes := must(ioutil.ReadFile("./files/xapi.json")).([]byte)                                  //   *** change file name ***
	YAMLScan(string(databytes), "ROOT", g.DELIPath, nil, g.JSON, func(path, value, id string) error { //   *** change idmark & DataType ***
		fPf("%s : %-70s : %s\n", id, path, value)
		return nil
	})
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
