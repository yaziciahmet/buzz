package buzz

import (
	"errors"
	"testing"
)

func Test_SliceForEachSuccess(t *testing.T) {
	slice := []int{1, 2, 3}

	if err := Slice[int]().ForEach(func(v int) error {
		if v == 5 {
			return errors.New("you shall not pass")
		}
		return nil
	}).Validate(slice); err != nil {
		t.FailNow()
	}
}

func Test_SliceForEachFail(t *testing.T) {
	slice := []int{1, 2, 3}

	if err := Slice[int]().ForEach(func(v int) error {
		if v == 3 {
			return errors.New("you shall not pass")
		}
		return nil
	}).Validate(slice); err == nil {
		t.FailNow()
	}
}

func Test_SliceForEachWithFieldValidationSuccess(t *testing.T) {
	emails := []string{"ahmet@mail.com", "yazici@mail.com"}

	if err := Slice[string]().ForEach(func(v string) error {
		return String().Email().Validate(v)
	}).Validate(emails); err != nil {
		t.FailNow()
	}
}
