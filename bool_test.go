package buzz

import "testing"

func Test_Bool(t *testing.T) {
	if err := Bool().Validate(true); err != nil {
		t.FailNow()
	}

	if err := Bool().True().Validate(true); err != nil {
		t.FailNow()
	}

	if err := Bool().False().Validate(false); err != nil {
		t.FailNow()
	}
}
