package xjy

import (
	"io/ioutil"
	"testing"

	"github.com/cdutwhu/go-gjxy"
	c "github.com/nsip/n3-client/config"
	g "github.com/nsip/n3-client/global"
)

func TestXMLScanObjects(t *testing.T) {
	sifbytes := must(ioutil.ReadFile("./files/sif.xml")).([]byte)
	objtags, ids, idtags, starts, ends := XMLScanObjects(string(sifbytes))
	fPln(len(objtags))
	for i := range ids {
		fPf("%25s -- %s -- %s -- %6d -- %6d\n", objtags[i], ids[i], idtags[i], starts[i], ends[i])
	}
}

func TestXMLObjStrByID(t *testing.T) {
	sifbytes := must(ioutil.ReadFile("./files/sif.xml")).([]byte)
	xmlobj := XMLObjStrByID(string(sifbytes), "1822AF7A-F9CB-4F0D-96EA-9280DD0B6AB2")
	fPln(xmlobj)
	fPln()
	fPln(XMLAttributes(xmlobj))
}

func TestXMLInfoScan(t *testing.T) {
	cfg := c.FromFile("../build/config.toml")
	defer func() { ph(recover(), cfg.ErrLog) }()

	sifbytes := must(ioutil.ReadFile("./files/sif.xml")).([]byte)
	XMLInfoScan(string(sifbytes), g.DELIPath,
		func(p, id string, v []string, lastOne bool) error {
			fPln("S --->>> ", p, " : ", v)
			return nil
		},
		func(p, id string, n int, lastOne bool) error {
			if n > 1 {
				fPf("A --->>> %-100s : %s : %d\n", p, id, n)
			}
			return nil
		},
	)
}

func TestXMLEleStrByTag(t *testing.T) {
	fPln(gjxy.XMLTagEleEx(`
	<OtherNames>
		<Name Type="AKA">
			<FamilyName>Anderson</FamilyName>
			<GivenName>Samuel</GivenName>
			<FullName>Samuel Anderson</FullName>
		</Name>
		<Name Type="PRF">
			<FamilyName>Rowinski</FamilyName>
			<GivenName>Sam</GivenName>
			<FullName>Sam Rowinski </FullName>
		</Name>
	</OtherNames>`, "Name", 3))
}
