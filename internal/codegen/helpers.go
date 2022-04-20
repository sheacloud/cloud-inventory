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

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToSnakeHyphen(str string) string {
	return strings.Replace(ToSnakeCase(str), "_", "-", -1)
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
