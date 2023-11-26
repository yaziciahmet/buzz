package buzz

import (
	"reflect"
	"time"
)

var (
	timeReflectType = reflect.TypeOf(time.Time{})
)

type BuzzTimeValidateFunc func(t time.Time) error

type BuzzTime struct {
	name          string
	validateFuncs []BuzzTimeValidateFunc
}

func Time() *BuzzTime {
	return &BuzzTime{}
}

func (t *BuzzTime) Name() string {
	return t.name
}

func (t *BuzzTime) Type() reflect.Type {
	return timeReflectType
}

func (t *BuzzTime) WithName(name string) BuzzField {
	t.name = name
	return t
}

func (t *BuzzTime) Validate(v any) error {
	vtime, ok := v.(time.Time)
	if !ok {
		return makeValidationError("", "type", "type not string")
	}

	for _, valFn := range t.validateFuncs {
		if err := valFn(vtime); err != nil {
			return err
		}
	}
	return nil
}

func (t *BuzzTime) Clone() BuzzField {
	return &BuzzTime{
		name:          t.name,
		validateFuncs: t.validateFuncs,
	}
}

func (t *BuzzTime) After(timestamp time.Time) *BuzzTime {
	t.addValidateFunc(func(v time.Time) error {
		if v.After(timestamp) {
			return nil
		}
		return makeValidationError("", "after", "after failed")
	})
	return t
}

func (t *BuzzTime) Before(timestamp time.Time) *BuzzTime {
	t.addValidateFunc(func(v time.Time) error {
		if v.Before(timestamp) {
			return nil
		}
		return makeValidationError("", "before", "before failed")
	})
	return t
}

func (t *BuzzTime) NotAfter(timestamp time.Time) *BuzzTime {
	t.addValidateFunc(func(v time.Time) error {
		if v.After(timestamp) {
			return makeValidationError("", "notAfter", "notAfter failed")
		}
		return nil
	})
	return t
}

func (t *BuzzTime) NotBefore(timestamp time.Time) *BuzzTime {
	t.addValidateFunc(func(v time.Time) error {
		if v.Before(timestamp) {
			return makeValidationError("", "notBefore", "notBefore failed")
		}
		return nil
	})
	return t
}

func (t *BuzzTime) Custom(fn BuzzTimeValidateFunc) *BuzzTime {
	t.addValidateFunc(fn)
	return t
}

func (t *BuzzTime) addValidateFunc(fn BuzzTimeValidateFunc) {
	t.validateFuncs = append(t.validateFuncs, fn)
}
