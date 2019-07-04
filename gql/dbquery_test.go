package gql

import (
	"io/ioutil"
	"testing"
	"time"
	g "../global"
)

// func TestPermu(t *testing.T) {
// 	rst := PermuIndices([]int{2, 3, 4, 2})
// 	for _, item := range rst {
// 		fPln(item)
// 	}
// }

// func TestAugIntsArr(t *testing.T) {
// 	rst := AugIntsArr([][]int{[]int{11, 22, 3, 8}, []int{4, 5, 6, 9}})
// 	fPln(rst)
// }

func TestStrsJoinEx(t *testing.T) {
	fPln(IArrStrJoinEx(Ss([]string{"a", "b", "c", "d", "e"}), Ss([]string{"1", "2"}), "#", " ~ "))
	fPln(S("a#1 ~ b#2 ~ c#4 ~ d ~ e").SplitEx(" ~ ", "#", "string", "int"))
}

func TestAMapIndicesList(t *testing.T) {
	objID := "d6971037-e2ff-4db6-a669-ba5fbab0d318"
	queryObject(g.CurCtx, objID) //                                 *** get root, mapStruct, mapValue ***

	fPln()

	// rst := mArrayIndicesList()
	// for k, v := range rst {
	// 	fPln(k, v)
	// 	AugIntsArr(v)
	// }

	rst := IPathListBymArr("xapi ~ courses ~ content_areas ~ investigations")
	for _, r := range rst {
		fPln(r)
	}
}

func TestQueryObject(t *testing.T) {

	objID := "c58f2f19-32ed-4258-a71b-6506b2a2f33b"
	queryObject(g.CurCtx, objID) //                                 *** get root, mapStruct, mapValue ***

	fPln(root)
	fPln("<-------------------------------------------------------------------------------------------------------------->")

	// mIPathObj := map[string]string{} //                                                *** in vars.go ***

	// JSONMake(mIPathObj, "xapi", "learning_area", "HSIE")
	// JSONMake(mIPathObj, "xapi", "subject", "Geography")
	// JSONMake(mIPathObj, "xapi", "stage", "1")
	// JSONMake(mIPathObj, "xapi", "yrLvls", "1")
	// JSONMake(mIPathObj, "xapi", "yrLvls", "2")
	// JSONMake(mIPathObj, "xapi", "yrLvls", "3")
	// JSONMake(mIPathObj, "xapi", "courses", "xapi ~ courses#1")
	// JSONMake(mIPathObj, "xapi", "courses", "xapi ~ courses#2")
	// JSONMake(mIPathObj, "xapi ~ courses#1", "name", "Features Of Places")
	// JSONMake(mIPathObj, "xapi ~ courses#1", "outcomes", "xapi ~ courses#1 ~ outcomes#1")
	// JSONMake(mIPathObj, "xapi ~ courses#1 ~ outcomes#1", "description", "describes features of places and the connections people have with places")
	// JSONMake(mIPathObj, "xapi ~ courses#1", "outcomes", "xapi ~ courses#1 ~ outcomes#2")
	// JSONMake(mIPathObj, "xapi ~ courses#1 ~ outcomes#2", "description", "identifies ways in which people interact with and care for places")
	// JSONMake(mIPathObj, "xapi ~ courses#2", "name", "People and Places")
	// JSONMake(mIPathObj, "xapi ~ courses#2", "outcomes", "xapi ~ courses#2 ~ outcomes#1")
	// JSONMake(mIPathObj, "xapi ~ courses#2 ~ outcomes#1", "description", "describes features of places and the connections people have with places")

	JSONBuild(root)
	json := JSONMakeRep(mIPathObj, PATH_DEL)
	ioutil.WriteFile("temp.json", []byte(json), 0666)

	// schema := SchemaMake("", root, PATH_DEL, CHILD_DEL)
	// schema = sRepAll(schema, "\t-", "\t")
	// schema = sRepAll(schema, "\t#", "\t")
	// ioutil.WriteFile(fSf("./yield/%s.gql", objID), []byte(schema), 0666)

	// for k, v := range mapStruct {
	// 	fPf("%-100s%s\n", k, v)
	// }

	time.Sleep(time.Second * 1)
}
