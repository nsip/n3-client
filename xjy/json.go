package xjy

import (
	u "github.com/cdutwhu/go-util"
)

const (
	lenGUID = 36 // GUID 36 chars length
)

// JSONScanObjects : (return the whole json string) (must have top-level "id" like `"id": "6690e6c9-3ef0-4ed3-8b37-7f3964730bee",` )
func JSONScanObjects(json, idmark string) (ids, objstrs []string, posarr []int) {
	idmark = u.Str(idmark).MkQuotes(u.QDouble)
	idmark = u.Str(idmark).MkSuffix(":")

	level, arrLevel, done1, done2 := 0, 0, false, false
	for i, c := range json {
		switch c {
		case '{':
			level++
		case '}':
			level--
		case '[':
			arrLevel++
		case ']':
			arrLevel--
		}

		/* single object */
		if level == 1 && arrLevel == 0 && done1 {
			continue
		}
		if level == 1 && arrLevel == 0 && !done1 {
			if p, pe, parr := sI(json[i:], idmark), sI(json[i:], "\","), sI(json[i:], "["); p > 0 && pe > 0 && parr < 0 {
				posarr = append(posarr, i)
				ids = append(ids, json[i+pe-lenGUID:i+pe])
				done1 = true
			}
		}
		if level == 0 && arrLevel == 0 && done1 {
			objstrs = append(objstrs, json[posarr[0]:i+1])
			break
		}

		/* object array */
		if level == 2 && arrLevel == 1 && done2 {
			continue
		}
		if level == 2 && arrLevel == 1 && !done2 {
			if p, pe := sI(json[i:], idmark), sI(json[i:], "\","); p > 0 && pe > 0 {
				posarr = append(posarr, i)
				ids = append(ids, json[i+pe-lenGUID:i+pe])
				done2 = true
			}
		}
		if level == 1 && arrLevel == 1 && done2 {
			objstrs = append(objstrs, json[posarr[len(posarr)-1]:i+1])
			done2 = false
		}
	}
	return
}

// JSONObjStrByID is
func JSONObjStrByID(json, idmark, ID string) string {
	ids, objstrs, _ := JSONScanObjects(json, idmark)
	for i, id := range ids {
		if id == ID {
			return objstrs[i]
		}
	}
	return ""
}

// JSONEleStrByTag is
func JSONEleStrByTag(json, tag string) string {
	l := len(json)

	if l == 0 || json[0] != '{' || json[l-1] != '}' {
		fPln(json)
		PE(fEf("Not a valid json section"))
		return ""
	}

	tag = u.Str(tag).MkQuotes(u.QDouble)
	tag = u.Str(tag).MkSuffix(":")

	level, arrLevel := 0, 0
	for _, c := range json {
		switch c {
		case '{':
			level++
		case '}':
			level--
		case '[':
			arrLevel++
		case ']':
			arrLevel--
		}

		if p := sI(json, tag); p >= 0 && level == 1 {
			peR := sI(json[p:], "\",")
			bFlat := !u.Str(json[p:p+peR+1]).HasAny('{', '}')

			if peR > 0 && bFlat { /* not last one, flat one */
				return u.Str(json[p : p+peR+1]).MkBrackets(u.BCurly)
			}
			if peR < 0 && bFlat {
				peR = sLI(json[p:], "\"")
				return u.Str(json[p : p+peR+1]).MkBrackets(u.BCurly)
			}
			if !bFlat { /* complex one */
				str, _, _ := u.Str(json[p:]).BracketsPos(u.BCurly, 1, 1)
				//return u.Str(json[p : p+rR+1]).MkBrackets(u.BCurly)
				return str
			}
		}
	}
	return u.Str("").MkBrackets(u.BCurly)
}

// JSONFindChildren :
func JSONFindChildren(jsonele string) (children []string, childList string) {
	l := len(jsonele)
	if l == 0 || jsonele[0] != '{' || jsonele[l-1] != '}' {
		fPln(jsonele)
		PE(fEf("Not a valid json section"))
		return nil, "nil"
	}

	level, childposl, childposr := 0, []int{}, []int{}
	for i, c := range jsonele[1:] { // skip the first '{'
		i++

		if c == '{' {
			level++
		}
		if c == '}' {
			level--
		}
		if level == 1 && c == '"' {
			if pR := sI(jsonele[i+1:], "\""); pR >= 0 && jsonele[i+1+pR+1] == ':' {
				// fPln(string(jsonele[i+1]))
				childposl = append(childposl, i)
				childposr = append(childposr, i+1+pR)
			}
		}
	}

	for i := range childposl {
		child := jsonele[childposl[i] : childposr[i]+1]

		/* deal with array element */
		childvalue := ""
		if i < len(childposl)-1 {
			childvalue = jsonele[childposr[i]+2 : childposl[i+1]]
		} else {
			childvalue = jsonele[childposr[i]+2:]
		}
		count := u.Str(childvalue).BracketPairCount(u.BCurly)

		if count == 0 { /* not an array element */
			children = append(children, u.Str(child).RmQuotes())
		} else {
			for j := 0; j < count; j++ {
				children = append(children, u.Str(child).RmQuotes())
			}
		}
	}

	// if len(children) > 1 && u.AllAreIdentical(children...) {
	// 	return children, spf("[%d]%s", len(children), children[0])
	// }

	return children, sJ(children, " + ")
}

// JSONYieldFamilyTree :
// func JSONYieldFamilyTree(jsonstr string, objs []string, skipNoChild bool, mapkeyprefix string, mapEleChildList *map[string]string) {
// 	if len(mapkeyprefix) > 0 {
// 		mapkeyprefix += "."
// 	}
// 	for _, obj := range objs {
// 		if _, ok := (*mapEleChildList)[mapkeyprefix+obj]; ok {
// 			continue
// 		}
// 		jsonele := JSONEleStrByTag(jsonstr, obj)
// 		children, childlist := JSONFindChildren(jsonele)

// 		if skipNoChild {
// 			if len(children) > 0 {

// 			}
// 		}
// 	}
// }
