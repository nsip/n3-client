package gql

import (
	"encoding/json"
	"io/ioutil"
	"os"

	g "../global"

	"github.com/playlyfe/go-graphql"
)

func qSchemaList(qSchemaDir string) (fnames []string) {
	files := Must(ioutil.ReadDir(qSchemaDir)).([]os.FileInfo)
	for _, file := range files {
		fname := file.Name()
		if !Str(fname).HP("_") {
			fnames = append(fnames, Str(fname).RmTailFromLast(".").V())
		}
	}
	return
}

func modStructMap(rmStructs ...string) {
	if v, ok := mStruct[root]; ok {
		// v = sRepAll(v, "+ StudentList +", "+") //      *** remove some children which are not needed ***
		for _, child := range rmStructs {
			v = sRepAll(v, child, "")
			v = Str(v).T(BLANK + "+").V()
			v = sRepAll(v, "+  +", "+")
		}
		mStruct[root] = v
	}
}

// GetInfoFromID : (FOR testing use)
func GetInfoFromID(infoType, objID string, rmStructs ...string) string {
	defer clrQueryCache()
	if objID == "" {
		return ""
	}
	queryObject(objID) //         *** GET root, mStruct, mArray, mValue ***
	modStructMap(rmStructs...)
	if len(mValue) == 0 {
		return ""
	}

	switch infoType {
	case "SCHEMA":
		{
			schema := SchemaBuild("", root)
			schema = sRepAll(schema, "\t-", "\t")
			schema = sRepAll(schema, "\t#", "\t")
			g.LCSchema.Add(objID, schema) //           *** LRU ***
			return schema
		}
	case "JSON":
		{
			JSONBuild(root)
			_, _, json := JSONWrapRoot(JSONMakeRep(mIPathObj, PATH_DEL), root)
			g.LCJSON.Add(objID, json) //               *** LRU ***
			return json
		}
	default:
		return "ERROR: Need SCHEMA or JSON for infoType"
	}
}

// GetResourceFromID :
func GetResourceFromID(objIDs []string, rmStructs ...string) (mSchema, mJSON map[string]string) {

	mSchema, mJSON = make(map[string]string), make(map[string]string)

	for _, objID := range objIDs {

		ok1, ok2, ok3 := false, false, false
		if schema, ok := g.LCSchema.Get(objID); ok {
			mSchema[objID], ok1 = schema.(string), ok
		}
		if json, ok := g.LCJSON.Get(objID); ok {
			mJSON[objID], ok2 = json.(string), ok
		}
		if rt, ok := g.LCRoot.Get(objID); ok {
			root, ok3 = rt.(string), ok
		}
		if ok1 && ok2 && ok3 {
			continue
		}

		// ********************************************************************* //

		clrQueryCache()
		g.RmIDsInLRU(objID)

		if objID == "" {
			mSchema[objID], mJSON[objID] = "", ""
			continue
		}

		queryObject(objID)         //    *** GET root, mStruct, mArray, mValue ***
		modStructMap(rmStructs...) //    *** eliminate some unnecessary properties ***

		if len(mStruct) == 0 || len(mValue) == 0 {
			mSchema[objID], mJSON[objID] = "", ""
			continue
		}

		schema := SchemaBuild("", root)
		schema = sRepAll(schema, "\t-", "\t")
		schema = sRepAll(schema, "\t#", "\t")
		mSchema[objID] = schema
		g.LCSchema.Add(objID, schema) //      *** LRU ***

		JSONBuild(root)
		_, _, json := JSONWrapRoot(JSONMakeRep(mIPathObj, PATH_DEL), root)
		mJSON[objID] = json
		g.LCJSON.Add(objID, json) //          *** LRU ***

		g.LCRoot.Add(objID, root) //          *** LRU ***
	}

	return
}

func rsvResource(objIDs []string, mJSON, mReplace map[string]string) []byte {
	jsonAll := ""
	for _, objID := range objIDs {
		jsonstr := mJSON[objID]
		ioutil.WriteFile("./debug/"+objIDs[0]+".json", []byte(jsonstr), 0666) //   *** DEBUG ***
		jsonAll = JSONObjectMerge(jsonAll, jsonstr)
	}
	for k, v := range mReplace {
		jsonAll = sRepAll(jsonAll, k, v)
	}
	return []byte(jsonAll)
}

// Query : if id is known, use Query
func Query(objIDs []string, qSchema, qSchemaDir, qTxt string, variables map[string]interface{}, rmStructs []string, mReplace map[string]string) (rstJSON string) {

	mSchema, mJSON := GetResourceFromID(objIDs, rmStructs...)

	autoSchema := mSchema[objIDs[0]]
	if autoSchema == "" {
		return ""
	}

	schema := qSchema + autoSchema //                                         *** qSchema is mannually coded ***
	for k, v := range mReplace {
		schema = sRepAll(schema, k, v)
	}

	ioutil.WriteFile("./debug/"+objIDs[0]+".gql", []byte(schema), 0666) //    *** DEBUG ***

	fResolver := func(params *graphql.ResolveParams) (interface{}, error) {
		jsonBytes := rsvResource(objIDs, mJSON, mReplace) //                  *** Get Reconstructed JSON ***
		jsonMap := make(map[string]interface{})
		PE(json.Unmarshal(jsonBytes, &jsonMap))
		// fPln(root)
		return jsonMap[root], nil
	}

	resolvers := map[string]interface{}{}
	for _, fname := range qSchemaList(qSchemaDir) {
		resolvers["QueryRoot/"+fname] = fResolver //             *** PATH : related to <querySchema> ***
	}

	context := map[string]interface{}{}
	// variables := map[string]interface{}{}
	executor, _ := graphql.NewExecutor(schema, "QueryRoot", "", resolvers)
	// executor.ResolveType = func(value interface{}) string {
	// 	if object, ok := value.(map[string]interface{}); ok {
	// 		return object["__typename"].(string)
	// 	}
	// 	return ""
	// }

	result := Must(executor.Execute(context, qTxt, variables, ""))
	rstJSON = string(Must(json.Marshal(result)).([]byte))
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

	// resolvers["QueryRoot/staff"] = func(params *graphql.ResolveParams) (interface{}, error) {
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
	// executor, _ := graphql.NewExecutor(schema, "QueryRoot", "", resolvers)
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
