package mapy

import (
	"fmt"
	"reflect"
)

type ErrInvalidKeyType struct {
	expected string
	got      string
}

func (e ErrInvalidKeyType) Error() string {
	return fmt.Sprintf("invalid key type; expected '%s', got '%s'", e.expected, e.got)
}

// StringKeys returns the string keys of m. If the keys are not of type string,
// ErrInvalidKeyType is returned.
func StringKeys(m interface{}) ([]string, error) {
	const expectedKeyType = "string"
	t := reflect.TypeOf(m)
	keyType := t.Key().String()
	if keyType != expectedKeyType {
		return nil, ErrInvalidKeyType{
			expected: expectedKeyType,
			got:      keyType,
		}
	}

	keys := reflect.ValueOf(m).MapKeys()
	strKeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strKeys[i] = keys[i].String()
	}

	return strKeys, nil
}
