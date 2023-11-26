package buzz

import (
	"reflect"
)

type BuzzInterfaceValidateFunc[T any] func(T) error

type BuzzInterface[T any] struct {
	name          string
	validateFuncs []BuzzInterfaceValidateFunc[T]
	refType       reflect.Type
	nullable      bool
}

func Interface[T any]() *BuzzInterface[T] {
	refType := reflect.TypeOf(new(T)).Elem()
	if refType.Kind() != reflect.Interface {
		panic("noninterface is passed as generic parameter")
	}

	return &BuzzInterface[T]{
		refType:  refType,
		nullable: true,
	}
}

func (i *BuzzInterface[T]) Name() string {
	return i.name
}

func (i *BuzzInterface[T]) Type() reflect.Type {
	return i.refType
}

func (i *BuzzInterface[T]) Validate(v any) error {
	if v == nil {
		if i.nullable {
			return nil
		}

		return makeValidationError("", "nonnil", "interface not nullable")
	}

	viface, ok := v.(T)
	if !ok {
		return makeValidationError("", "type", "interface type not T")
	}

	for _, valFn := range i.validateFuncs {
		if err := valFn(viface); err != nil {
			return err
		}
	}
	return nil
}

func (i *BuzzInterface[T]) WithName(name string) BuzzField {
	i.name = name
	return i
}

func (i *BuzzInterface[T]) Clone() BuzzField {
	return &BuzzInterface[T]{
		name:          i.name,
		validateFuncs: i.validateFuncs,
		refType:       i.refType,
		nullable:      i.nullable,
	}
}

func (i *BuzzInterface[T]) Nonnil() *BuzzInterface[T] {
	i.nullable = false
	return i
}

func (i *BuzzInterface[T]) MustBeType(typ T) *BuzzInterface[T] {
	expectedType := reflect.TypeOf(typ)
	i.addValidateFunc(func(v T) error {
		actualType := reflect.TypeOf(v)
		if expectedType != actualType {
			return makeValidationError("", "mustbetype", "mustbetype failed")
		}

		return nil
	})

	return i
}

func (i *BuzzInterface[T]) Custom(fn BuzzInterfaceValidateFunc[T]) *BuzzInterface[T] {
	i.addValidateFunc(fn)
	return i
}

func (i *BuzzInterface[T]) addValidateFunc(fn BuzzInterfaceValidateFunc[T]) {
	i.validateFuncs = append(i.validateFuncs, fn)
}
