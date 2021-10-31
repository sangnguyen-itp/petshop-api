package tools

import "fmt"

func buildFunction(fnName string, content string, ptr string, objectName string, params string, returns string)  string {
	fn := fmt.Sprintf("func (%s *%s) %s(%s) %s { \n", ptr, objectName, fnName, params, returns)
	fn = fmt.Sprintf("%s %s \n", fn, content)
	fn = fmt.Sprintf("%s }", fn)
	return fn
}

