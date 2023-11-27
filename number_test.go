package buzz

import (
	"fmt"
	"testing"
)

func Test_NumberGteSuccess(t *testing.T) {
	if err := Number[int]().Gte(10).Validate(15); err != nil {
		t.FailNow()
	}
}

func Test_NumberGtFail(t *testing.T) {
	if err := Number[int]().Gt(10).Validate(5); err == nil {
		t.FailNow()
	}
}

func Test_NumberLteSuccess(t *testing.T) {
	if err := Number[int]().Lte(10).Validate(5); err != nil {
		t.FailNow()
	}
}

func Test_NumberLtFail(t *testing.T) {
	if err := Number[int]().Lt(10).Validate(15); err == nil {
		t.FailNow()
	}
}

func Test_NumberPositiveSuccess(t *testing.T) {
	if err := Number[int64]().Positive().Validate(15); err == nil {
		t.FailNow()
	}
}

func Test_NumberInvalidTypeError(t *testing.T) {
	err := Number[int]().Validate("hello")
	if err.Error() != fmt.Sprintf(invalidTypeMsg, "int", "") {
		t.FailNow()
	}
}
