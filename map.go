package buzz

import (
	"fmt"
	"reflect"
)

type BuzzMapValidateFunc[K comparable, V any] func(map[K]V) error

type BuzzMap[K comparable, V any] struct {
	name          string
	validateFuncs []BuzzMapValidateFunc[K, V]
	refType       reflect.Type
	nullable      bool
}

func Map[K comparable, V any]() *BuzzMap[K, V] {
	return &BuzzMap[K, V]{
		refType:  reflect.TypeOf(*new(map[K]V)),
		nullable: true,
	}
}

func (m *BuzzMap[K, V]) Name() string {
	return m.name
}

func (m *BuzzMap[K, V]) Type() reflect.Type {
	return m.refType
}

func (m *BuzzMap[K, V]) Validate(v any) error {
	if v == nil {
		if m.nullable {
			return nil
		}

		return notNullableFieldErr(m.name)
	}

	vMap, ok := v.(map[K]V)
	if !ok {
		return fmt.Errorf(invalidTypeMsg, m.refType, v)
	}

	errAggr := NewFieldErrorAggregator()
	for _, valFn := range m.validateFuncs {
		if err := valFn(vMap); err != nil {
			if errAggr.Handle(err) != nil {
				return err
			}
		}
	}

	return errAggr.OrNil()
}

func (m *BuzzMap[K, V]) WithName(name string) BuzzField {
	m.name = name
	return m
}

func (m *BuzzMap[K, V]) Clone() BuzzField {
	return &BuzzMap[K, V]{
		name:          m.name,
		validateFuncs: m.validateFuncs,
		refType:       m.refType,
		nullable:      m.nullable,
	}
}

func (m *BuzzMap[K, V]) Nonnil() *BuzzMap[K, V] {
	m.nullable = false
	return m
}

func (m *BuzzMap[K, V]) Nonempty() *BuzzMap[K, V] {
	m.registerValidateFunc(func(v map[K]V) error {
		if len(v) == 0 {
			return nonEmptyFieldErr(m.name)
		}
		return nil
	})
	return m
}

func (m *BuzzMap[K, V]) ContainsKey(key K) *BuzzMap[K, V] {
	m.registerValidateFunc(func(v map[K]V) error {
		if _, ok := v[key]; !ok {
			return MakeFieldError(m.name, "ContainsKey", fmt.Sprintf("map %s does not contain key %v", m.name, key))
		}
		return nil
	})
	return m
}

func (m *BuzzMap[K, V]) Custom(fn BuzzMapValidateFunc[K, V]) *BuzzMap[K, V] {
	m.registerValidateFunc(fn)
	return m
}

func (m *BuzzMap[K, V]) registerValidateFunc(fn BuzzMapValidateFunc[K, V]) {
	m.validateFuncs = append(m.validateFuncs, fn)
}
