package xjy

import (
	"io/ioutil"
	"testing"

	c "../config"
	u "github.com/cdutwhu/go-util"
)

func TestXMLScanObjects(t *testing.T) {
	cfg := c.FromFile("./config.toml", "../config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog) }()

	sifbytes := Must(ioutil.ReadFile("./files/sif.xml")).([]byte)
	sif := u.Str(sifbytes)
	sif.SetEnC()

	mapStructRecord := map[string]string{}
	n := sif.XMLSegsCount()
	prevEnd := 0
	for iObj := 1; iObj <= n; iObj++ {
		nextStart := u.TerOp(iObj == 1, 0, prevEnd+1).(int)
		_, xml, _, end := sif.S(nextStart, u.ALL).XMLSegPos(1, 1)
		prevEnd = end + nextStart

		// fPf("%d SIF *****************************************\n", iObj)

		procIdx := 1
		XMLModelInfo(xml, "RefId", pathDel, childDel,
			func(p, v string) {
				if prevV, ok := mapStructRecord[p]; !ok || (ok && v != prevV && u.Str(v).FieldsSeqContain(prevV, childDel)) {
					mapStructRecord[p] = v
					fPf("S%4d ---> %-70s:: %s\n", procIdx, p, v)
					procIdx++
				}
			},
			func(p, v string, n int) {
				fPf("A%4d ---> %-70s[] %s  -- [%d]\n", procIdx, p, v, n)
				procIdx++
			},
		)
	}

	// idx := 1
	// XMLModelInfo(string(sifbytes), "RefId", pathDel, childDel,
	// 	func(p, v string) {
	// 		fPf("S ---> %-5d: %-90s:: %s\n", idx, p, v)
	// 		idx++
	// 	},
	// 	func(p, v string, n int) {
	// 		fPf("A ---> %-5d: %-90s[] %s  -- [%d]\n", idx, p, v, n)
	// 		idx++
	// 	},
	// )
	// fPf("finish:\n")

	// ids, objtags, psarr := XMLScanObjects(string(sifbytes), "RefId")
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
