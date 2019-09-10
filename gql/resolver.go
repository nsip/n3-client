package gql

import (
	"encoding/json"
	"io/ioutil"
	"os"

	g "github.com/nsip/n3-client/global"
	pp "github.com/nsip/n3-client/preprocess"

	"github.com/playlyfe/go-graphql"
)

func qSchemaList(qSchemaDir string) (fnames []string) {
	files := must(ioutil.ReadDir(qSchemaDir)).([]os.FileInfo)
	for _, file := range files {
		fname := file.Name()
		if !S(fname).HP("_") {
			fnames = append(fnames, S(fname).RmTailFromLast(".").V())
		}
	}
	return
}

// GetInfoFromID : (FOR testing use)
func GetInfoFromID(ctx, infoType, objID string) string {
	defer clrQueryCache()
	if objID == "" {
		return ""
	}
	queryObject(ctx, objID) //     *** GET root, mStruct, mArray, mValue ***

	if len(mValue) == 0 {
		return ""
	}

	switch infoType {
	case "SCHEMA":
		{
			schema := SchemaBuild("", root)
			schema = sRepAll(schema, "\t-", "\t")
			schema = sRepAll(schema, "\t#", "\t")
			return schema
		}
	case "JSON":
		{
			JSONBuild(root)
			_, _, json := JSONWrapRoot(JSONMakeRep(mIPathObj, g.DELIPath), root)
			return json
		}
	default:
		return "ERROR: Need SCHEMA or JSON or QRYTXT for infoType"
	}
}

// GetDataFromID :
func GetDataFromID(ctx string, objIDs []string) (mSchema, mJSON map[string]string) {

	mSchema, mJSON = make(map[string]string), make(map[string]string)

	for _, objID := range objIDs {

		clrQueryCache()

		if objID == "" {
			mSchema[objID], mJSON[objID] = "", ""
			continue
		}

		queryObject(ctx, objID) //    *** GET root, mStruct, mArray, mValue ***

		if len(mStruct) == 0 || len(mValue) == 0 {
			mSchema[objID], mJSON[objID] = "", ""
			continue
		}

		schema := SchemaBuild("", root)
		schema = sRepAll(schema, "\t-", "\t")
		schema = sRepAll(schema, "\t#", "\t")
		mSchema[objID] = schema

		JSONBuild(root)
		_, _, json := JSONWrapRoot(JSONMakeRep(mIPathObj, g.DELIPath), root)
		mJSON[objID] = json
	}

	return
}

func rsvResource(objIDs []string, mJSON, mReplace map[string]string) []byte {
	jsonAll := ""
	for _, ID := range objIDs {
		json := mJSON[ID]
		json = pp.FmtJSONStr(json, "../preprocess/util/", "./preprocess/util/", "./")
		json = ASCIIToOri(json) //                                                      *** ascii back to original ***
		for k, v := range mReplace {
			json = sRepAll(json, k, v)
		}
		ioutil.WriteFile("./debug_qry/"+ID+".json", []byte(json), 0666) //              *** DEBUG ***
		jsonAll = JSONObjectMerge(jsonAll, json)
	}
	return []byte(jsonAll)
}

// Query : if id is known, use Query
func Query(ctx string, objIDs []string, qSchemaDir, qTxt string, variables map[string]interface{}, mReplace map[string]string) (rstJSON string) {

	mSchema, mJSON := GetDataFromID(ctx, objIDs)
	autoSchema := mSchema[objIDs[0]]
	if autoSchema == "" {
		return ""
	}

	// fPln(root)
	qSchema := string(must(ioutil.ReadFile(qSchemaDir + root + ".gql")).([]byte)) //  *** content must be related to resolver path ***

	schs := sSpl(qSchema, DELISchema)
	pc(len(schs) < 2, fEf("qSchema header error, missing array schema?\n"))
	qSchema = trueAssign(len(objIDs) == 1, len(objIDs) > 1, schs[0], schs[1], qSchema).(string)

	schema := qSchema + autoSchema
	for k, v := range mReplace {
		schema = sRepAll(schema, k, v)
	}

	// ioutil.WriteFile("./debug/"+objIDs[0]+".gql", []byte(schema), 0666) //    *** DEBUG ***

	resolvers := map[string]interface{}{}
	for _, fname := range qSchemaList(qSchemaDir) {
		resolvers["Query/"+fname] = func(params *graphql.ResolveParams) (interface{}, error) { // *** PATH : related to <querySchema> ***
			jsonBytes := rsvResource(objIDs, mJSON, mReplace) //                                      *** Get Reconstructed JSON ***
			jsonMap := make(map[string]interface{})
			pe(json.Unmarshal(jsonBytes, &jsonMap))
			return jsonMap[root], nil
		}
	}

	context := map[string]interface{}{}
	// variables := map[string]interface{}{}
	executor, _ := graphql.NewExecutor(schema, "Query", "", resolvers)
	// executor.ResolveType = func(value interface{}) string {
	// 	if object, ok := value.(map[string]interface{}); ok {
	// 		return object["__typename"].(string)
	// 	}
	// 	return ""
	// }

	result := must(executor.Execute(context, qTxt, variables, ""))
	rstJSON = string(must(json.Marshal(result)).([]byte))
	// ioutil.WriteFile("temp.json", []byte(rstJSON), 0666)
	return

	// ********* DEMO ************
	// resolvers["DemoQuery/GetStudentProgress"] = func(params *graphql.ResolveParams) (interface{}, error) {
	// 	tgID := objIDs[0]
	// 	jsonTeachingGroup := "test"
	// 	jsonGradingAssignment := "test"
	// 	fPln(tgID, jsonTeachingGroup, jsonGradingAssignment)
	// 	return nil, nil
	// }

	// resolvers["Query/staff"] = func(params *graphql.ResolveParams) (interface{}, error) {
	// 	if len(params.Field.Arguments) == 0 {
	// 		return nil, nil
	// 	}
	// 	args := params.Args
	// 	argName := params.Field.Arguments[0].Name.Value
	// 	argValue := args[argName]
	// 	fmt.Println(argName, argValue)
	// 	data := []map[string]interface{}{
	// 		{
	// 			"LocalId":          "abc",
	// 			"StateProvinceId":  "456",
	// 			"Title":            "Principal",
	// 			"EmploymentStatus": "A",
	// 		},
	// 		{
	// 			"LocalId":          "123 222",
	// 			"StateProvinceId":  "456 222",
	// 			"Title":            "Principal 222",
	// 			"EmploymentStatus": "A 222",
	// 		},
	// 	}
	// 	rt := []map[string]interface{}{}
	// 	for _, mp := range data {
	// 		if v, ok := mp[argName]; ok && v == argValue {
	// 			rt = append(rt, mp)
	// 		}
	// 	}
	// 	return rt, nil
	// }

	// context := map[string]interface{}{}
	// // variables := map[string]interface{}{}
	// executor, _ := graphql.NewExecutor(schema, "Query", "", resolvers)
	// // executor.ResolveType = func(value interface{}) string {
	// // 	if object, ok := value.(map[string]interface{}); ok {
	// // 		return object["__typename"].(string)
	// // 	}
	// // 	return ""
	// // }

	// result := Must(executor.Execute(context, queryStr, variables, ""))
	// rstJSON = string(Must(json.Marshal(result)).([]byte))
	// // ioutil.WriteFile("temp.json", []byte(rstJSON), 0666)
	// return
}
