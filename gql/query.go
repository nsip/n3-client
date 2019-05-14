package gql

import (
	"encoding/json"
	"io/ioutil"

	"github.com/playlyfe/go-graphql"
)

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

// GetSchemaFromID : (FOR testing use)
func GetSchemaFromID(objID string, rmStructs ...string) (schema string) {
	defer clrQueryBuf()
	if objID == "" {
		return ""
	}
	queryObject(objID) //                                 *** GET root, mapStruct, mapArray, mapValue ***
	modStructMap(rmStructs...)
	if len(mStruct) == 0 || len(mValue) == 0 {
		return ""
	}

	schema = SchemaBuild("", root)
	schema = sRepAll(schema, "\t-", "\t")
	schema = sRepAll(schema, "\t#", "\t")
	return
}

// GetJSONFromID : (FOR testing use)
func GetJSONFromID(objID string, rmStructs ...string) (json string) {
	defer clrQueryBuf()
	if objID == "" {
		return ""
	}
	queryObject(objID) //                                 *** GET root, mapStruct, mapArray, mapValue ***
	modStructMap(rmStructs...)
	if len(mValue) == 0 {
		return ""
	}

	JSONBuild(root)
	json = JSONMakeRep(mIPathObj, PATH_DEL)
	return
}

// GetResourceFromID :
func GetResourceFromID(objIDs []string, rmStructs ...string) (mSchema, mJSON map[string]string) {
	mSchema, mJSON = make(map[string]string), make(map[string]string)
	for _, objID := range objIDs {
		clrQueryBuf()
		
		if objID == "" {
			mSchema[objID], mJSON[objID] = "", ""
			continue
		}

		queryObject(objID)         //                     *** GET root, mapStruct, mapArray, mapValue ***
		modStructMap(rmStructs...) //                     *** eliminate some unnecessary properties ***
		
		if len(mStruct) == 0 || len(mValue) == 0 {
			mSchema[objID], mJSON[objID] = "", ""
			continue
		}

		schema := SchemaBuild("", root)
		schema = sRepAll(schema, "\t-", "\t")
		schema = sRepAll(schema, "\t#", "\t")
		mSchema[objID] = schema

		JSONBuild(root)
		json := JSONMakeRep(mIPathObj, PATH_DEL)
		mJSON[objID] = json
	}
	return
}

func rsvResource(objIDs []string, mJSON map[string]string) []byte {
	jsonAll := ""
	for _, objID := range objIDs {
		jsonstr := mJSON[objID]
		_, _, jsonstr = JSONWrapRoot(jsonstr, "xapi")
		ioutil.WriteFile("./debug/"+objIDs[0]+".json", []byte(jsonstr), 0666)
		jsonAll = JSONObjectMerge(jsonAll, jsonstr)
	}
	return []byte(jsonAll)
}

// Query : if id is known, use Query
func Query(objIDs []string, querySchema, queryStr string, variables map[string]interface{}, rmStructs []string) (rstJSON string) {

	mSchema, mJSON := GetResourceFromID(objIDs, rmStructs...)

	autoSchema := mSchema[objIDs[0]] // GetSchemaFromID(objIDs[0], rmStructs...)
	if autoSchema == "" {
		return ""
	}

	schema := querySchema + autoSchema //                                    *** querySchema is mannually coded ***
	// schema = sRepAll(schema, "en-US", "enUS")
	ioutil.WriteFile("./debug/"+objIDs[0]+".gql", []byte(schema), 0666) //   *** DEBUG ***

	resolvers := map[string]interface{}{}
	// *** PATH : related to <querySchema> ***
	resolvers["QueryRoot/xapi"] = func(params *graphql.ResolveParams) (interface{}, error) {
		jsonBytes := rsvResource(objIDs, mJSON) //                           *** Get Reconstructed JSON ***
		jsonMap := make(map[string]interface{})
		PE(json.Unmarshal(jsonBytes, &jsonMap))
		return jsonMap[root], nil
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

	result := Must(executor.Execute(context, queryStr, variables, ""))
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
