package main

import "reflect"

var (
	stringReflectType = reflect.TypeOf("")
)

type BuzzStringValidateFunc func(string) error

type BuzzString struct {
	name          string
	validateFuncs []BuzzStringValidateFunc
}

func String(name string) *BuzzString {
	return &BuzzString{name: name}
}

func (s *BuzzString) Name() string {
	return s.name
}

func (s *BuzzString) Type() reflect.Type {
	return stringReflectType
}

func (s *BuzzString) Validate(v any) error {
	for _, valFn := range s.validateFuncs {
		if err := valFn(v.(string)); err != nil {
			return err
		}
	}
	return nil
}

func (s *BuzzString) Min(min int) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if min > len(v) {
			return makeValidationError(s.name, "min", "min failed")
		}
		return nil
	})
	return s
}

func (s *BuzzString) Max(max int) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if max < len(v) {
			return makeValidationError(s.name, "max", "max failed")
		}
		return nil
	})
	return s
}

func (s *BuzzString) addValidateFunc(fn BuzzStringValidateFunc) {
	s.validateFuncs = append(s.validateFuncs, fn)
}
