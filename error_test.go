package buzz

import "testing"

func Test_ErrorSimple(t *testing.T) {
	if err := Error().MustBeType(&MyError{}).Validate(&MyError{}); err != nil {
		t.FailNow()
	}
}

func Test_ErrorNil(t *testing.T) {
	if err := Error().Validate(nil); err != nil {
		t.FailNow()
	}
}
