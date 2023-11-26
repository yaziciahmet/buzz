package buzz

import (
	"reflect"
	"testing"
)

func Test_MapType(t *testing.T) {
	m := map[string]int{}
	mType := reflect.TypeOf(m)

	if Map[string, int]().Type() != mType {
		t.FailNow()
	}
}

func Test_MapContainsKey(t *testing.T) {
	if err := Map[string, int]().ContainsKey("test").Validate(map[string]int{"test": 123}); err != nil {
		t.FailNow()
	}
}
