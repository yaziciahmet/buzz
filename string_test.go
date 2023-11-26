package buzz

import (
	"testing"
)

func Test_StringEmailSuccess(t *testing.T) {
	if err := String().Email().Validate("ahmet@mail.com"); err != nil {
		t.FailNow()
	}
}

func Test_StringEmailFail(t *testing.T) {
	if err := String().Email().Validate("ahmetmail.com"); err == nil {
		t.FailNow()
	}
}

func Test_StringMinSuccess(t *testing.T) {
	if err := String().Min(2).Validate("ahmet"); err != nil {
		t.FailNow()
	}
}

func Test_StringMinFail(t *testing.T) {
	if err := String().Min(2).Validate("a"); err == nil {
		t.FailNow()
	}
}

func Test_StringMaxSuccess(t *testing.T) {
	if err := String().Max(10).Validate("ahmet"); err != nil {
		t.FailNow()
	}
}

func Test_StringMaxFail(t *testing.T) {
	if err := String().Max(10).Validate("ahmetyazici"); err == nil {
		t.FailNow()
	}
}

func Test_StringLenSuccess(t *testing.T) {
	if err := String().Len(5).Validate("ahmet"); err != nil {
		t.FailNow()
	}
}

func Test_StringLenFail(t *testing.T) {
	if err := String().Len(5).Validate("ahmetyazici"); err == nil {
		t.FailNow()
	}
}
