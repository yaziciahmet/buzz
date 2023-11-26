package buzz

import (
	"reflect"
)

type AnyNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type BuzzNumberValidateFunc[T AnyNumber] func(v T) error

type BuzzNumber[T AnyNumber] struct {
	name          string
	validateFuncs []BuzzNumberValidateFunc[T]
	refType       reflect.Type
}

func Number[T AnyNumber]() *BuzzNumber[T] {
	return &BuzzNumber[T]{
		refType: reflect.TypeOf(*new(T)),
	}
}

func (n *BuzzNumber[T]) Name() string {
	return n.name
}

func (n *BuzzNumber[T]) Type() reflect.Type {
	return n.refType
}

func (n *BuzzNumber[T]) Validate(v any) error {
	for _, valFn := range n.validateFuncs {
		vint, ok := v.(T)
		if !ok {
			return makeValidationError("", "type", "type not T")
		}

		if err := valFn(vint); err != nil {
			return err
		}
	}
	return nil
}

func (n *BuzzNumber[T]) WithName(name string) BuzzField {
	n.name = name
	return n
}

func (n *BuzzNumber[T]) Clone() BuzzField {
	return &BuzzNumber[T]{
		name:          n.name,
		validateFuncs: n.validateFuncs,
		refType:       n.refType,
	}
}

func (n *BuzzNumber[T]) Min(min T) *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if min > v {
			return makeValidationError("", "min", "min failed")
		}
		return nil
	})
	return n
}

func (n *BuzzNumber[T]) Max(max T) *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if max < v {
			return makeValidationError("", "max", "max failed")
		}
		return nil
	})
	return n
}

func (n *BuzzNumber[T]) Positive() *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if v <= 0 {
			return makeValidationError("", "positive", "positive failed")
		}
		return nil
	})
	return n
}

func (n *BuzzNumber[T]) Nonnegative() *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if v < 0 {
			return makeValidationError("", "nonnegative", "nonnegative failed")
		}
		return nil
	})
	return n
}

func (n *BuzzNumber[T]) Negative() *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if v >= 0 {
			return makeValidationError("", "negative", "negative failed")
		}
		return nil
	})
	return n
}

func (n *BuzzNumber[T]) NonPositive() *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if v > 0 {
			return makeValidationError("", "nonpositive", "nonpositive failed")
		}
		return nil
	})
	return n
}

func (n *BuzzNumber[T]) Custom(fn BuzzNumberValidateFunc[T]) *BuzzNumber[T] {
	n.addValidateFunc(fn)
	return n
}

func (n *BuzzNumber[T]) addValidateFunc(fn BuzzNumberValidateFunc[T]) {
	n.validateFuncs = append(n.validateFuncs, fn)
}
