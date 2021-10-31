package tools

import (
	"fmt"
	"strings"
)

func parseList(list string) []string {
	arr := strings.Split(list, ",")
	for idx, a := range arr {
		arr[idx] = strings.TrimSpace(a)
	}
	return arr
}

func GetSetUpdateQuery(input string) string {
	list := parseList(input)
	var result []string
	for _, row := range list {
		if row == "" {
			continue
		}
		customRow := appendComment(row)
		result = append(result, customRow)
	}

	return join(result, "\n")
}

func GetTransformFunc(ptr, obj string, input string) string {
	list := parseList(input)
	var result []string
	for _, row := range list {
		if row == "" {
			continue
		}
		elements := strings.Split(row, " ")
		field := strings.Trim(elements[0], `"`)
		data := buildTransformIfCond(ptr, field)
		result = append(result, data)
	}
	content :=  join(result, "\n\n")
	transformReturnContent := fmt.Sprintf("%s \n return query", content)
	return buildFunction("transform", transformReturnContent, ptr, obj, "query string", "string")
}

func buildTransformIfCond(ob string, field string) string {
	r := fmt.Sprintf("if %s.Is%s { \n", ob, field)
	r = fmt.Sprintf(`%s query = strings.ReplaceAll(query, "--!Is%s", "")`, r, field)
	r = fmt.Sprintf("%s \n }", r)
	return r
}

func appendComment(row string) string {
	elements := strings.Split(row, " ")
	field := strings.Trim(elements[0], `"`)
	return fmt.Sprintf(`--!Is%s "%s" = :%s,`, field, field, field)
}

func join(list []string, sep string) string {
	result := strings.Join(list, sep)
	return result
}