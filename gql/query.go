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

// GetSchemaFromID :
func GetSchemaFromID(objID string, rmStructs ...string) (schema string) {
	defer clrQueryBuf()
	if objID == "" {
		return ""
	}
	queryObject(objID) //                                 *** GET root, mapStruct, mapArray, mapValue ***
	modStructMap(rmStructs...)
	schema = SchemaBuild("", root)
	schema = sRepAll(schema, "\t-", "\t")
	schema = sRepAll(schema, "\t#", "\t")
	return
}

// GetJSONFromID :
func GetJSONFromID(objID string, rmStructs ...string) (json string) {
	defer clrQueryBuf()
	if objID == "" {
		return ""
	}
	queryObject(objID) //                                 *** GET root, mapStruct, mapArray, mapValue ***
	modStructMap(rmStructs...)
	JSONBuild(root)
	json = JSONMakeRep(mIPathObj, PATH_DEL)
	// ioutil.WriteFile("GetJSONFromID.json", []byte(json), 0666)
	return
}

// GetJSONGQLFromID :
// func GetJSONGQLFromID(objID string) (json, schema string) {
// 	defer clrQueryBuf()
// 	queryObject(objID) //                                 *** GET root, mapStruct, mapArray, mapValue ***
// 	modStructMap() //                               *** eliminate some unnecessary properties ***
// 	json = JSONMake("", root, pathDel, childDel)
// 	json = sRepAll(json, `"-`, `"`)
// 	json = sRepAll(json, `"#`, `"`)
// 	schema = SchemaMake("", root, pathDel, childDel)
// 	schema = sRepAll(schema, "\t-", "\t")
// 	schema = sRepAll(schema, "\t#", "\t")
// 	return
// }

func rsvResource(objIDs []string, rmStructs []string) []byte {	
	jsonAll := ""
	for _, objID := range objIDs {
		jsonstr := GetJSONFromID(objID, rmStructs...)
		_, _, jsonstr = JSONWrapRoot(jsonstr, "xapi")
		ioutil.WriteFile("./debug/"+objIDs[0]+".json", []byte(jsonstr), 0666)
		jsonAll = JSONObjectMerge(jsonAll, jsonstr)
	}
	return []byte(jsonAll)
}

// Query : if id is known, use Query
func Query(objIDs []string, querySchema, queryStr string, variables map[string]interface{}, rmStructs []string) (rstJSON string) {

	autoSchema := GetSchemaFromID(objIDs[0], rmStructs...)
	if autoSchema == "" {
		return ""
	}

	schema := querySchema + autoSchema //                                    *** querySchema has mannually coded ***
	// schema = sRepAll(schema, "en-US", "enUS")	
	ioutil.WriteFile("./debug/"+objIDs[0]+".gql", []byte(schema), 0666) //   *** DEBUG ***

	resolvers := map[string]interface{}{}
	// *** PATH : related to <querySchema> ***
	resolvers["QueryRoot/xapi"] = func(params *graphql.ResolveParams) (interface{}, error) {
		jsonBytes := rsvResource(objIDs, rmStructs) //                       *** Get Reconstructed JSON ***
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

	// resolvers := map[string]interface{}{}
	// resolvers["DemoQuery/TeachingGroupByName"] = func(params *graphql.ResolveParams) (interface{}, error) {
	// 	jsonBytes := rsvResource(objIDs, rmStructs)
	// 	jsonMap := make(map[string]interface{})
	// 	PE(json.Unmarshal(jsonBytes, &jsonMap))
	// 	return jsonMap[root], nil
	// }
	// resolvers["DemoQuery/TeachingGroupByStaffID"] = func(params *graphql.ResolveParams) (interface{}, error) {
	// 	jsonBytes := rsvResource(objIDs, rmStructs)
	// 	jsonMap := make(map[string]interface{})
	// 	PE(json.Unmarshal(jsonBytes, &jsonMap))
	// 	return jsonMap[root], nil
	// }
	// resolvers["DemoQuery/TeachingGroup"] = func(params *graphql.ResolveParams) (interface{}, error) {
	// 	jsonBytes := rsvResource(objIDs, rmStructs)
	// 	jsonMap := make(map[string]interface{})
	// 	PE(json.Unmarshal(jsonBytes, &jsonMap))
	// 	return jsonMap[root], nil
	// }
	// resolvers["DemoQuery/GradingAssignment"] = func(params *graphql.ResolveParams) (interface{}, error) {
	// 	jsonBytes := rsvResource(objIDs, rmStructs)
	// 	jsonMap := make(map[string]interface{})
	// 	PE(json.Unmarshal(jsonBytes, &jsonMap))
	// 	return jsonMap[root], nil
	// }
	// resolvers["DemoQuery/StudentAttendance"] = func(params *graphql.ResolveParams) (interface{}, error) {
	// 	jsonBytes := rsvResource(objIDs, rmStructs)
	// 	jsonMap := make(map[string]interface{})
	// 	PE(json.Unmarshal(jsonBytes, &jsonMap))
	// 	return jsonMap[root], nil
	// }
	// resolvers["DemoQuery/QueryXAPI"] = func(params *graphql.ResolveParams) (interface{}, error) {
	// 	jsonBytes := rsvResource(objIDs, rmStructs)
	// 	jsonMap := make(map[string]interface{})
	// 	PE(json.Unmarshal(jsonBytes, &jsonMap))
	// 	return jsonMap[root], nil
	// }

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
