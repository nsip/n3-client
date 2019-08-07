package gql

import (
	"io/ioutil"
	"testing"

	g "github.com/nsip/n3-client/global"
	q "github.com/nsip/n3-client/query"
)

func TestGQL(t *testing.T) {
	objID := "ca669951-9511-4e53-ae92-50845d3bdcd6"
	ctx := g.CurCtx
	_, _, o, _ := q.Data(ctx, objID, "") //           *** get root ***
	if len(o) > 0 {
		root := o[0]
		qSchemaDir, qTxtDir := "./qSchema/", "./qTxt/"
		qTxt := string(must(ioutil.ReadFile(qTxtDir + root + ".txt")).([]byte)) // *** change ***
		result := Query(
			ctx,
			[]string{objID},
			qSchemaDir,
			qTxt,
			map[string]interface{}{},
			g.MpQryRstRplc,
		)

		ioutil.WriteFile(fSf("./yield/%s.json", objID), []byte(result), 0666)
		return
	}
	fPln("wrong objID")
}

func TestQSchemaList(t *testing.T) {
	fnames := qSchemaList("./qSchema")
	for _, f := range fnames {
		fPln(f)
	}
}
