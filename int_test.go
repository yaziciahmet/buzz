package buzz

import "testing"

func Test_IntMinSuccess(t *testing.T) {
	if err := Int("Id").Min(10).Validate(15); err != nil {
		t.FailNow()
	}
}

func Test_IntMinFail(t *testing.T) {
	if err := Int("Id").Min(10).Validate(5); err == nil {
		t.FailNow()
	}
}

func Test_IntMaxSuccess(t *testing.T) {
	if err := Int("Id").Max(10).Validate(5); err != nil {
		t.FailNow()
	}
}

func Test_IntMaxFail(t *testing.T) {
	if err := Int("Id").Max(10).Validate(15); err == nil {
		t.FailNow()
	}
}
