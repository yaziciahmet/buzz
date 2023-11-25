package buzz

import (
	"reflect"
)

var (
	sliceReflectType = reflect.TypeOf([]int{})
)

type BuzzSliceValidateFunc[T any] func(v []T) error

type BuzzSlice[T any] struct {
	name          string
	validateFuncs []BuzzSliceValidateFunc[T]
}

func Slice[T any](name string) *BuzzSlice[T] {
	return &BuzzSlice[T]{name: name}
}

func (s *BuzzSlice[T]) Name() string {
	return s.name
}

func (s *BuzzSlice[T]) Type() reflect.Type {
	return stringReflectType
}

func (s *BuzzSlice[T]) Validate(v any) error {
	for _, valFn := range s.validateFuncs {
		if err := valFn(v.([]T)); err != nil {
			return err
		}
	}
	return nil
}

func (s *BuzzSlice[T]) Min(min int) *BuzzSlice[T] {
	s.addValidateFunc(func(v []T) error {
		if min > len(v) {
			return makeValidationError(s.name, "min", "min failed")
		}
		return nil
	})
	return s
}

func (s *BuzzSlice[T]) Max(max int) *BuzzSlice[T] {
	s.addValidateFunc(func(v []T) error {
		if max < len(v) {
			return makeValidationError(s.name, "max", "max failed")
		}
		return nil
	})
	return s
}

func (s *BuzzSlice[T]) Len(l int) *BuzzSlice[T] {
	s.addValidateFunc(func(v []T) error {
		if l != len(v) {
			return makeValidationError(s.name, "len", "len failed")
		}
		return nil
	})
	return s
}

func (s *BuzzSlice[T]) ForEach(fn func(T) error) *BuzzSlice[T] {
	s.addValidateFunc(func(v []T) error {
		for _, el := range v {
			if err := fn(el); err != nil {
				return err
			}
		}
		return nil
	})
	return s
}

func (s *BuzzSlice[T]) addValidateFunc(fn BuzzSliceValidateFunc[T]) {
	s.validateFuncs = append(s.validateFuncs, fn)
}
