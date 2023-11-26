package buzz

import (
	"testing"
	"time"
)

func Test_TimeAfterSuccess(t *testing.T) {
	now := time.Now()
	if err := Time().After(now).Validate(now.Add(1)); err != nil {
		t.FailNow()
	}
}

func Test_TimeBeforeSuccess(t *testing.T) {
	now := time.Now()
	if err := Time().Before(now).Validate(now.Add(-1)); err != nil {
		t.FailNow()
	}
}
