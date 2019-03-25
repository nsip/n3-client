package gql

import (
	"encoding/json"

	u "github.com/cdutwhu/go-util"
	"github.com/playlyfe/go-graphql"
)

func modStructMap(rmStructs ...string) {
	if v, ok := mapStruct[root]; ok {
		// v = sRepAll(v, "+ StudentList +", "+") //                *** remove some children which are not needed ***
		for _, child := range rmStructs {
			v = sRepAll(v, child, "")
			v = u.Str(v).T(u.BLANK + "+").V()
			v = sRepAll(v, "+  +", "+")
		}
		mapStruct[root] = v
	}
}

// GetSchemaFromID :
func GetSchemaFromID(objID, fType string, rmStructs ...string) (schema string) {
	defer clrQueryBuf()
	queryObject(objID, fType) //                                 *** GET root, mapStruct, mapArray, mapValue ***
	modStructMap(rmStructs...)
	schema = SchemaMake("", root, pathDel, childDel)
	schema = sRepAll(schema, "\t-", "\t")
	schema = sRepAll(schema, "\t#", "\t")
	return
}

// GetJSONFromID :
func GetJSONFromID(objID, fType string, rmStructs ...string) (json string) {
	defer clrQueryBuf()
	queryObject(objID, fType) //                                 *** GET root, mapStruct, mapArray, mapValue ***
	modStructMap(rmStructs...)
	json = JSONMake("", root, pathDel, childDel)
	json = sRepAll(json, `"-`, `"`)
	json = sRepAll(json, `"#`, `"`)
	return
}

// GetJSONGQLFromID :
// func GetJSONGQLFromID(objID, fType string) (json, schema string) {
// 	defer clrQueryBuf()
// 	queryObject(objID, fType) //                                 *** GET root, mapStruct, mapArray, mapValue ***
// 	modStructMap() //                               *** eliminate some unnecessary properties ***
// 	json = JSONMake("", root, pathDel, childDel)
// 	json = sRepAll(json, `"-`, `"`)
// 	json = sRepAll(json, `"#`, `"`)
// 	schema = SchemaMake("", root, pathDel, childDel)
// 	schema = sRepAll(schema, "\t-", "\t")
// 	schema = sRepAll(schema, "\t#", "\t")
// 	return
// }

func rsvResource(objIDs []string, fType string, rmStructs []string) []byte {
	jsonAll := ""
	for _, objID := range objIDs {
		jsonstr := GetJSONFromID(objID, fType, rmStructs...)
		jsonAll = u.Str(jsonAll).JSONObjectMerge(jsonstr)
	}
	return []byte(jsonAll)
}

// GQuery : if we know id, use GQuery
func GQuery(objIDs []string, fType, schemaQuery, queryStr string, variables map[string]interface{}, rmStructs []string) (rstJSON string) {

	schemaRuntime := GetSchemaFromID(objIDs[0], fType, rmStructs...)
	// ioutil.WriteFile("debug_schema.gql", []byte(schema), 0666)
	schema := schemaQuery + schemaRuntime //                            *** schemaQuery has mannually coded schemas ***
	schema = sRepAll(schema, "en-US", "enUS")
	//ioutil.WriteFile("./yield/"+objID+".json", []byte(jsonstr), 0666) //  *** DEBUG ***
	//ioutil.WriteFile("./yield/"+objID+".gql", []byte(schema), 0666)   //  *** DEBUG ***
	//return

	resolvers := map[string]interface{}{}
	resolvers["DemoQuery/TeachingGroupByName"] = func(params *graphql.ResolveParams) (interface{}, error) {
		jsonBytes := rsvResource(objIDs, fType, rmStructs)
		jsonMap := make(map[string]interface{})
		PE(json.Unmarshal(jsonBytes, &jsonMap))
		return jsonMap[root], nil
	}

	resolvers["DemoQuery/TeachingGroupByStaffID"] = func(params *graphql.ResolveParams) (interface{}, error) {
		jsonBytes := rsvResource(objIDs, fType, rmStructs)
		jsonMap := make(map[string]interface{})
		PE(json.Unmarshal(jsonBytes, &jsonMap))
		return jsonMap[root], nil
	}

	resolvers["DemoQuery/TeachingGroup"] = func(params *graphql.ResolveParams) (interface{}, error) {
		jsonBytes := rsvResource(objIDs, fType, rmStructs)
		jsonMap := make(map[string]interface{})
		PE(json.Unmarshal(jsonBytes, &jsonMap))
		return jsonMap[root], nil
	}

	resolvers["DemoQuery/GradingAssignment"] = func(params *graphql.ResolveParams) (interface{}, error) {
		jsonBytes := rsvResource(objIDs, fType, rmStructs)
		jsonMap := make(map[string]interface{})
		PE(json.Unmarshal(jsonBytes, &jsonMap))
		return jsonMap[root], nil
	}

	resolvers["DemoQuery/StudentAttendance"] = func(params *graphql.ResolveParams) (interface{}, error) {
		jsonBytes := rsvResource(objIDs, fType, rmStructs)
		jsonMap := make(map[string]interface{})
		PE(json.Unmarshal(jsonBytes, &jsonMap))
		return jsonMap[root], nil
	}

	resolvers["DemoQuery/QueryXAPI"] = func(params *graphql.ResolveParams) (interface{}, error) {
		jsonBytes := rsvResource(objIDs, fType, rmStructs)
		jsonMap := make(map[string]interface{})
		PE(json.Unmarshal(jsonBytes, &jsonMap))
		return jsonMap[root], nil
	}

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

	context := map[string]interface{}{}
	// variables := map[string]interface{}{}
	executor, _ := graphql.NewExecutor(schema, "DemoQuery", "", resolvers)
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
}
