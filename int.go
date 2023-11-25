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

func (f *BuzzInt) Name() string {
	return f.name
}

func (f *BuzzInt) Type() reflect.Type {
	return intReflectType
}

func (f *BuzzInt) Validate(v any) error {
	for _, valFn := range f.validateFuncs {
		if err := valFn(v.(int)); err != nil {
			return err
		}
	}
	return nil
}

func (f *BuzzInt) Min(min int) *BuzzInt {
	f.validateFuncs = append(f.validateFuncs, func(v int) error {
		if min > v {
			return makeValidationError(f.name, "min", "min failed")
		}
		return nil
	})
	return f
}

func (f *BuzzInt) Max(max int) *BuzzInt {
	f.validateFuncs = append(f.validateFuncs, func(v int) error {
		if max < v {
			return makeValidationError(f.name, "max", "max failed")
		}
		return nil
	})
	return f
}
