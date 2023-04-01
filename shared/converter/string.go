package converter

import (
	"strings"
	"unicode"
)

func CamelCaseToSnakeCase(str string) string {
	var builder strings.Builder
	for i, c := range str {
		if unicode.IsUpper(c) {
			if i > 0 {
				builder.WriteByte('_')
			}
			builder.WriteRune(unicode.ToLower(c))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
