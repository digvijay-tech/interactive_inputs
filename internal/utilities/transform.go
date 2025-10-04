package utilities

import (
	"reflect"
	"strings"
)

/*
Returns the tyoe of x, if x is an array/slice it will return type prefixed with '[]', You can set second argument removeBrackets to true if you prefer to
remove '[]' from the return value for array/slice types.
If x is nil interface value this function will return an empty string ""
*/
func FindType[T any](x T, removeBrackets bool) string {
	result := reflect.TypeOf(x)

	if result == nil {
		return ""
	}

	if removeBrackets {
		return strings.TrimPrefix(result.String(), "[]")
	}

	return result.String()
}

/*
Transforms given text by applying transformations such as uppercase, lowercase or capitalise.
If transformation string matches none it will return the original string.
*/
func TextTransform(transformation, text string) string {
	if transformation == "uppercase" {
		return strings.ToUpper(text)
	}

	if transformation == "lowercase" {
		return strings.ToLower(text)
	}

	if transformation == "capitalise" {
		return toCapitalise(text)
	}

	return text
}

func toCapitalise(s string) string {
	clean := strings.TrimSpace(s)

	if clean == "" {
		return clean
	}

	words := strings.Split(clean, " ")
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	}

	return strings.Join(words, " ")
}
