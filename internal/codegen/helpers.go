package codegen

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

var (
	ErrNotAStruct           = errors.New("given objects underlying type is not a struct")
	ErrUnsupportedFieldType = errors.New("given field type is not supported")

	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
	mapRegex      = regexp.MustCompile(`map\[(.*)\](.*)`)
)

func Deduplicate(s []string) []string {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}
	var res []string
	for k := range m {
		res = append(res, k)
	}
	return res
}

func makeDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func stringInList(s string, list []string) bool {
	for _, s2 := range list {
		if s == s2 {
			return true
		}
	}
	return false
}

func isTimeField(s string) bool {
	return s == "time.Time" || s == "*time.Time"
}

func isTagField(s string) bool {
	return s == "Tags" || s == "TagSet" || s == "TagList"
}

func isUnmodelableField(s string) bool {
	return s == "github.com/Azure/go-autorest/autorest/date.Time" || s == "*github.com/Azure/go-autorest/autorest/date.Time"
}

func typeToParquetType(s string) string {
	switch s {
	case "string":
		return "BYTE_ARRAY"
	case "int32":
		return "INT32"
	case "int":
		return "INT32"
	case "bool":
		return "BOOLEAN"
	case "int64":
		return "INT64"
	case "float32":
		return "FLOAT"
	case "float64":
		return "DOUBLE"
	default:
		return ""
	}
}

func typeToParquetConvertedType(s string) string {
	switch s {
	case "string":
		return "UTF8"
	default:
		return ""
	}
}
