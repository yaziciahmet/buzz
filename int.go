package main

import "reflect"

var (
	intReflectType = reflect.TypeOf(0)
)

type BuzzIntValidateFunc func(v int) error

type BuzzInt struct {
	name          string
	validateFuncs []BuzzIntValidateFunc
}

func Int(name string) *BuzzInt {
	return &BuzzInt{name: name}
}

func (i *BuzzInt) Name() string {
	return i.name
}

func (i *BuzzInt) Type() reflect.Type {
	return intReflectType
}

func (i *BuzzInt) Validate(v any) error {
	for _, valFn := range i.validateFuncs {
		if err := valFn(v.(int)); err != nil {
			return err
		}
	}
	return nil
}

func (i *BuzzInt) Min(min int) *BuzzInt {
	i.addValidateFunc(func(v int) error {
		if min > v {
			return makeValidationError(i.name, "min", "min failed")
		}
		return nil
	})
	return i
}

func (i *BuzzInt) Max(max int) *BuzzInt {
	i.addValidateFunc(func(v int) error {
		if max < v {
			return makeValidationError(i.name, "max", "max failed")
		}
		return nil
	})
	return i
}

func (i *BuzzInt) addValidateFunc(fn BuzzIntValidateFunc) {
	i.validateFuncs = append(i.validateFuncs, fn)
}
