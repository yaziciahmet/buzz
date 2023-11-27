package buzz

import (
	"fmt"
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
	vint, ok := v.(T)
	if !ok {
		return fmt.Errorf(invalidTypeMsg, n.refType, v)
	}

	errAggr := NewFieldErrorAggregator()
	for _, valFn := range n.validateFuncs {
		if err := valFn(vint); err != nil {
			if errAggr.Handle(err) != nil {
				return err
			}
		}
	}

	return errAggr.OrNil()
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

func (n *BuzzNumber[T]) Gt(num T) *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if v > num {
			return nil
		}
		return MakeFieldError(n.name, "Gt", fmt.Sprintf("%s must be greater than %v", n.name, num))
	})
	return n
}

func (n *BuzzNumber[T]) Gte(num T) *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if v >= num {
			return nil
		}
		return MakeFieldError(n.name, "Gte", fmt.Sprintf("%s must be greate than or equal to %v", n.name, num))
	})
	return n
}

func (n *BuzzNumber[T]) Lt(num T) *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if v < num {
			return nil
		}
		return MakeFieldError(n.name, "Lt", fmt.Sprintf("%s must be less than %v", n.name, num))
	})
	return n
}

func (n *BuzzNumber[T]) Lte(num T) *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if v <= num {
			return nil
		}
		return MakeFieldError(n.name, "Lte", fmt.Sprintf("%s must be less than or equal to %v", n.name, num))
	})
	return n
}

func (n *BuzzNumber[T]) Positive() *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if v <= 0 {
			return MakeFieldError(n.name, "Positive", fmt.Sprintf("%s must be positive", n.name))
		}
		return nil
	})
	return n
}

func (n *BuzzNumber[T]) Nonnegative() *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if v < 0 {
			return MakeFieldError(n.name, "Nonnegative", fmt.Sprintf("%s must be nonnegative", n.name))
		}
		return nil
	})
	return n
}

func (n *BuzzNumber[T]) Negative() *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if v >= 0 {
			return MakeFieldError(n.name, "Negative", fmt.Sprintf("%s must be negative", n.name))
		}
		return nil
	})
	return n
}

func (n *BuzzNumber[T]) Nonpositive() *BuzzNumber[T] {
	n.addValidateFunc(func(v T) error {
		if v > 0 {
			return MakeFieldError(n.name, "Nonpositive", fmt.Sprintf("%s must be nonpositive", n.name))
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
