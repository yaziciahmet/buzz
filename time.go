package buzz

import (
	"fmt"
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
		return fmt.Errorf(invalidTypeMsg, timeReflectType, v)
	}

	errAggr := NewFieldErrorAggregator()
	for _, valFn := range t.validateFuncs {
		if err := valFn(vtime); err != nil {
			if errAggr.Handle(err) != nil {
				return err
			}
		}
	}

	return errAggr.OrNil()
}

func (t *BuzzTime) Clone() BuzzField {
	return &BuzzTime{
		name:          t.name,
		validateFuncs: t.validateFuncs,
	}
}

func (t *BuzzTime) After(timestamp time.Time) *BuzzTime {
	t.registerValidateFunc(func(v time.Time) error {
		if v.After(timestamp) {
			return nil
		}
		return MakeFieldError(t.name, "After", fmt.Sprintf("%s must be after %s", t.name, timestamp))
	})
	return t
}

func (t *BuzzTime) Before(timestamp time.Time) *BuzzTime {
	t.registerValidateFunc(func(v time.Time) error {
		if v.Before(timestamp) {
			return nil
		}
		return MakeFieldError(t.name, "Before", fmt.Sprintf("%s must be before %s", t.name, timestamp))
	})
	return t
}

func (t *BuzzTime) NotAfter(timestamp time.Time) *BuzzTime {
	t.registerValidateFunc(func(v time.Time) error {
		if v.After(timestamp) {
			return MakeFieldError(t.name, "NotAfter", fmt.Sprintf("%s must not be after %s", t.name, timestamp))
		}
		return nil
	})
	return t
}

func (t *BuzzTime) NotBefore(timestamp time.Time) *BuzzTime {
	t.registerValidateFunc(func(v time.Time) error {
		if v.Before(timestamp) {
			return MakeFieldError(t.name, "NotBefore", fmt.Sprintf("%s must not be before %s", t.name, timestamp))
		}
		return nil
	})
	return t
}

func (t *BuzzTime) Custom(fn BuzzTimeValidateFunc) *BuzzTime {
	t.registerValidateFunc(fn)
	return t
}

func (t *BuzzTime) registerValidateFunc(fn BuzzTimeValidateFunc) {
	t.validateFuncs = append(t.validateFuncs, fn)
}
