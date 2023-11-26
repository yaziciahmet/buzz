package buzz

import (
	"errors"
	"testing"
)

func Test_InterfaceError(t *testing.T) {
	iface := Interface[error]().MustBeType(&MyError{})

	if err := iface.Validate(&MyError{}); err != nil {
		t.FailNow()
	}

	if err := iface.Validate(errors.New("dracarys")); err == nil {
		t.FailNow()
	}
}

func Test_InterfaceCustom(t *testing.T) {
	if err := Interface[MyInterface]().Custom(func(mi MyInterface) error {
		if mi.Method() == 100 {
			return errors.New("hell nah")
		}
		return nil
	}).Validate(&MyInterfaceStruct{}); err != nil {
		t.FailNow()
	}
}

func Test_InterfaceNil(t *testing.T) {
	if err := Interface[MyInterface]().Validate(nil); err != nil {
		t.FailNow()
	}
}
