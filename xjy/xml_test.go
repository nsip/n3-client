package xjy

import (
	"io/ioutil"
	"testing"

	c "../config"
	"github.com/cdutwhu/go-gjxy"
)

func TestXMLScanObjects(t *testing.T) {
	sifbytes := Must(ioutil.ReadFile("./files/sif.xml")).([]byte)
	ids, objtags, starts, ends := XMLScanObjects(string(sifbytes), "RefId")
	fPln(len(objtags))
	for i := range ids {
		fPf("%25s -- %s -- %6d -- %6d\n", objtags[i], ids[i], starts[i], ends[i])
	}
}

func TestXMLObjStrByID(t *testing.T) {
	sifbytes := Must(ioutil.ReadFile("./files/sif.xml")).([]byte)
	xmlobj := XMLObjStrByID(string(sifbytes), "RefId", "1822AF7A-F9CB-4F0D-96EA-9280DD0B6AB2")
	fPln(xmlobj)
	fPln()

	fPln(XMLAttributes(xmlobj, "-"))
}

func TestXMLInfoScan(t *testing.T) {
	cfg := c.FromFile("./config.toml", "../config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog) }()

	sifbytes := Must(ioutil.ReadFile("./files/sif.xml")).([]byte)
	XMLInfoScan(string(sifbytes), "RefId", PATH_DEL,
		func(p, id string, v []string, lastOne bool) {
			fPln("S --->>> ", p, " : ", v)
		},
		func(p, id string, n int, lastOne bool) {
			if n > 1 {
				fPf("A --->>> %-100s : %s : %d\n", p, id, n)
			}
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
