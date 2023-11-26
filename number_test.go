package buzz

import "testing"

func Test_NumberMinSuccess(t *testing.T) {
	if err := Number[int]().Min(10).Validate(15); err != nil {
		t.FailNow()
	}
}

func Test_NumberMinFail(t *testing.T) {
	if err := Number[int]().Min(10).Validate(5); err == nil {
		t.FailNow()
	}
}

func Test_NumberMaxSuccess(t *testing.T) {
	if err := Number[int]().Max(10).Validate(5); err != nil {
		t.FailNow()
	}
}

func Test_NumberMaxFail(t *testing.T) {
	if err := Number[int]().Max(10).Validate(15); err == nil {
		t.FailNow()
	}
}

func Test_NumberPositiveSuccess(t *testing.T) {
	if err := Number[int64]().Max(10).Validate(15); err == nil {
		t.FailNow()
	}
}
