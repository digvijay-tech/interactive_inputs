package utilities

import (
	"fmt"
	"reflect"
	"strings"
)

func FindArrayType[T any](list []T, removeBrackets bool) (string, error) {
	result := reflect.TypeOf(list)

	if result == nil {
		return "", fmt.Errorf("provided list type %v is invalid", result)
	}

	if removeBrackets {
		return strings.TrimPrefix(result.String(), "[]"), nil
	}

	return result.String(), nil
}

func ToCapitalise(s string) string {
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
