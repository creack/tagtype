package tagtype

import (
	"fmt"
	"reflect"
	"strings"
)

// typeMap maps the common types for printf.
// map[go type]printf type
// TODO: use reflect to lookup type.
var typeMap = map[string]string{
	"[]byte": "[]uint8",
	"byte":   "uint8",
}

// inArray returns true if `search` is found within `array`.
// It also maps fields based on the `typeMap`.
func inArray(search string, array []string) bool {
	for _, elem := range array {
		if elem == search || typeMap[elem] == search {
			return true
		}
	}
	return false
}

// Validates iterates the given struct and make sure the value
// of each field is consistent with it's `tt` struct tag.
func Validate(in interface{}) bool {
	// Fetch type and value data from reflect
	st := reflect.TypeOf(in)
	s := reflect.ValueOf(in)
	if reflect.TypeOf(in).Kind() == reflect.Ptr {
		st = st.Elem()
		s = s.Elem()
	}

	// Iterate through the struct fields
	for i := 0; i < s.NumField(); i++ {
		// Fetch type and value data for the current field.
		fieldType, fieldValue := st.Field(i), s.Field(i)
		tagStr := fieldType.Tag.Get("tt")
		if tagStr == "" {
			// If no `tt` tag present, do not validate
			continue
		}
		actualType := fmt.Sprintf("%T", fieldValue.Interface())
		if actualType == "<nil>" {
			// If the value is <nil>, do not validate
			continue
		}
		tags := strings.Split(tagStr, ",")
		if !inArray(actualType, tags) {
			return false
		}
	}
	return true
}
