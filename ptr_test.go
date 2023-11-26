package buzz

import (
	"reflect"
	"testing"
)

func Test_PtrType(t *testing.T) {
	str := ""
	if Ptr(String()).Type() != reflect.TypeOf(&str) {
		t.FailNow()
	}
}

func Test_PtrValidate(t *testing.T) {
	str := "ahmet@mail.com"
	if err := Ptr(String().Email()).Validate(&str); err != nil {
		t.FailNow()
	}
}

func Test_PtrValidateNull(t *testing.T) {
	if err := Ptr(String().Email()).Nullable().Validate(nil); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func Test_PtrValidateNullWithType(t *testing.T) {
	var str *string
	if err := Ptr(String().Email()).Nullable().Validate(str); err != nil {
		t.FailNow()
	}
}
