package buzz

import (
	"reflect"
	"time"
)

var (
	timeReflectType = reflect.TypeOf(time.Time{})
)

type BuzzTimestampValidateFunc func(t time.Time) error

type BuzzTimestamp struct {
	name          string
	validateFuncs []BuzzTimestampValidateFunc
}

func Timestamp(name string) *BuzzTimestamp {
	return &BuzzTimestamp{name: name}
}

func (t *BuzzTimestamp) Name() string {
	return t.name
}

func (t *BuzzTimestamp) Type() reflect.Type {
	return timeReflectType
}

func (t *BuzzTimestamp) Validate(v any) error {
	for _, valFn := range t.validateFuncs {
		if err := valFn(v.(time.Time)); err != nil {
			return err
		}
	}
	return nil
}

func (t *BuzzTimestamp) After(timestamp time.Time) *BuzzTimestamp {
	t.addValidateFunc(func(v time.Time) error {
		if v.After(timestamp) {
			return nil
		}
		return makeValidationError(t.name, "after", "after failed")
	})
	return t
}

func (t *BuzzTimestamp) Before(timestamp time.Time) *BuzzTimestamp {
	t.addValidateFunc(func(v time.Time) error {
		if v.Before(timestamp) {
			return nil
		}
		return makeValidationError(t.name, "before", "before failed")
	})
	return t
}

func (t *BuzzTimestamp) NotAfter(timestamp time.Time) *BuzzTimestamp {
	t.addValidateFunc(func(v time.Time) error {
		if v.After(timestamp) {
			return makeValidationError(t.name, "notAfter", "notAfter failed")
		}
		return nil
	})
	return t
}

func (t *BuzzTimestamp) NotBefore(timestamp time.Time) *BuzzTimestamp {
	t.addValidateFunc(func(v time.Time) error {
		if v.Before(timestamp) {
			return makeValidationError(t.name, "notBefore", "notBefore failed")
		}
		return nil
	})
	return t
}

func (t *BuzzTimestamp) Custom(fn BuzzTimestampValidateFunc) *BuzzTimestamp {
	t.addValidateFunc(fn)
	return t
}

func (t *BuzzTimestamp) addValidateFunc(fn BuzzTimestampValidateFunc) {
	t.validateFuncs = append(t.validateFuncs, fn)
}
