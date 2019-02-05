package xjy

import (
	"io/ioutil"
	"testing"

	c "../config"
)

func TestXMLScanObjects(t *testing.T) {
	cfg := c.GetConfig("./config.toml", "../config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog, true) }()

	//xmlbytes, err := ioutil.ReadFile("./files/staffpersonal.xml")
	xmlbytes := Must(ioutil.ReadFile("./files/nswdig.xml")).([]byte)

	XMLModelInfo(string(xmlbytes), "RefId", true,
		func(p, v string) {
			fPf("%-90s:: %s\n", p, v)
		},
		func(p, v string, n int) {
			fPf("%-90s:: %s  -- [%d]\n", p, v, n)
		},
	)
	fPf("finish:\n")

	// ids, objtags, psarr := XMLScanObjects(string(xmlbytes), "RefId")
	// fPln(len(objtags))
	// for _, objtag := range objtags {
	// 	fPln(objtag)
	// }
	// for i := range ids {
	// 	fPf("%s -- %s -- %d\n", objtags[i], ids[i], psarr[i])
	// }

	//fmt.Print(string(xmlbytes[psarr[1]:psarr[2]]))

	// xmlobj := XMLObjStrByID(string(xmlbytes), "RefId", "D3E34F41-9D75-101A-8C3D-00AA001A1652")
	// fPln(xmlobj)
	// fPln()
	// fPln(XMLFindAttributes(xmlobj))
}

func TestXMLEleStrByTag(t *testing.T) {
	fPln(XMLEleStrByTag(`		<OtherNames>
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
</OtherNames>`, "Name", 1))
}
