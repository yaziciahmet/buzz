package buzz

import (
	"fmt"
	"reflect"
)

type BuzzSliceValidateFunc[T any] func(v []T) error
type BuzzSliceElementValidateFunc[T any] func(v T) error

type BuzzSlice[T any] struct {
	name          string
	validateFuncs []BuzzSliceValidateFunc[T]
	refType       reflect.Type
	nullable      bool
}

func Slice[T any]() *BuzzSlice[T] {
	return &BuzzSlice[T]{
		refType:  reflect.TypeOf(*new([]T)),
		nullable: true,
	}
}

func (s *BuzzSlice[T]) Name() string {
	return s.name
}

func (s *BuzzSlice[T]) Type() reflect.Type {
	return s.refType
}

func (s *BuzzSlice[T]) Validate(v any) error {
	if v == nil {
		if s.nullable {
			return nil
		}

		return MakeFieldError("", "nonnil", "slice not nullable")
	}

	vTSlice, ok := v.([]T)
	if !ok {
		return fmt.Errorf(invalidTypeMsg, s.refType, v)
	}

	for _, valFn := range s.validateFuncs {
		if err := valFn(vTSlice); err != nil {
			return err
		}
	}
	return nil
}

func (s *BuzzSlice[T]) WithName(name string) BuzzField {
	s.name = name
	return s
}

func (s *BuzzSlice[T]) Clone() BuzzField {
	return &BuzzSlice[T]{
		name:          s.name,
		validateFuncs: s.validateFuncs,
		refType:       s.refType,
		nullable:      s.nullable,
	}
}

func (s *BuzzSlice[T]) Min(min int) *BuzzSlice[T] {
	s.addValidateFunc(func(v []T) error {
		if min > len(v) {
			return MakeFieldError("", "min", "min failed")
		}
		return nil
	})
	return s
}

func (s *BuzzSlice[T]) Max(max int) *BuzzSlice[T] {
	s.addValidateFunc(func(v []T) error {
		if max < len(v) {
			return MakeFieldError("", "max", "max failed")
		}
		return nil
	})
	return s
}

func (s *BuzzSlice[T]) Len(l int) *BuzzSlice[T] {
	s.addValidateFunc(func(v []T) error {
		if l != len(v) {
			return MakeFieldError("", "len", "len failed")
		}
		return nil
	})
	return s
}

func (s *BuzzSlice[T]) Nonempty() *BuzzSlice[T] {
	s.addValidateFunc(func(v []T) error {
		if len(v) == 0 {
			return MakeFieldError("", "nonempty", "nonempty failed")
		}
		return nil
	})
	return s
}

func (s *BuzzSlice[T]) Nonnil() *BuzzSlice[T] {
	s.nullable = false
	return s
}

func (s *BuzzSlice[T]) ForEach(fn BuzzSliceElementValidateFunc[T]) *BuzzSlice[T] {
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

func (s *BuzzSlice[T]) Custom(fn BuzzSliceValidateFunc[T]) *BuzzSlice[T] {
	s.addValidateFunc(fn)
	return s
}

func (s *BuzzSlice[T]) addValidateFunc(fn BuzzSliceValidateFunc[T]) {
	s.validateFuncs = append(s.validateFuncs, fn)
}
