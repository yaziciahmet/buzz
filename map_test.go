package buzz

import (
	"reflect"
	"testing"
)

func Test_MapType(t *testing.T) {
	mType := reflect.TypeOf(map[string]int{})
	if Map[string, int]().Type() != mType {
		t.FailNow()
	}
}

func Test_MapContainsKey(t *testing.T) {
	if err := Map[string, int]().ContainsKey("test").Validate(map[string]int{"test": 123}); err != nil {
		t.FailNow()
	}
}

func Test_MapNil(t *testing.T) {
	if err := Map[any, any]().Validate(nil); err != nil {
		t.FailNow()
	}
}

func Test_MapNonnil(t *testing.T) {
	if err := Map[any, any]().Nonnil().Validate(nil); err == nil {
		t.FailNow()
	}
}
