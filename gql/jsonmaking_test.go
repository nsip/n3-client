package gql

// func jumpinto(src, mark, property, value string) (string, bool) {
// 	left, right, rp := "", "", 0
// 	property = u.Str(property).MkQuotes(u.QDouble) + ": "
// 	if value[0] != '{' && value[0] != '[' {
// 		value = u.Str(value).MkQuotes(u.QDouble)
// 	}

// 	if mark == "" { //							          *** empty json file ***
// 		if pe := sLI(src, "}"); pe > 0 {
// 			left, right = src[:pe], src[pe:]
// 		} else {
// 			return src, false
// 		}
// 	} else {
// 		mark = u.Str(mark).MkQuotes(u.QDouble) + ":" // ***
// 		if p := sI(src, mark); p > 0 {
// 			srcFromMark := src[p:]
// 			_, _, r := u.Str(srcFromMark).BracketsPos(u.BCurly, 1, 1)
// 			rp = r + p
// 			left, right = src[:rp], src[rp:]
// 		}
// 	}

// 	if !sHS(left, "{") {
// 		left += ","
// 	}

// 	return left + " " + property + value + " " + right, true
// }

// func TestJumpInto(t *testing.T) {
// 	r, b := jumpinto(`{
// 		"StaffPersonal": {
// 			"-RefId": {},
// 			"LocalId": "946379881",
// 			"StateProvinceId": "C2345681",
// 			"OtherIdList": {
// 				"OtherId": {
// 					"-Type": "0004",
// 					"#content": "333333333"
// 				}
// 			},
// 			"PersonInfo": {
// 				"Name": {
// 					"-Type": "LGL",
// 					"FamilyName": "Smith",
// 					"GivenName": "Fred",
// 					"FullName": "Fred Smith"
// 				},
// 				"OtherNames": {
// 					"Name": [
// 						{},
// 						{}
// 					]
// 				},
// 			}
// 		}
// 	}`,

// 		"Name",
// 		"FamilyName",
// 		"Anderson")
// 	fPln(r, b)
// }
