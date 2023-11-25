package buzz

import (
	"fmt"
	"reflect"
)

type BuzzSchemaValidateFunc[T any] func(T) error

type BuzzField interface {
	Name() string
	Validate(v any) error
	Type() reflect.Type
}

type BuzzSchema[T any] struct {
	name          string
	fields        []BuzzField
	validateFuncs []BuzzSchemaValidateFunc[T]
	refType       reflect.Type
}

func Schema[T any](refObj T, fields ...BuzzField) *BuzzSchema[T] {
	refType := reflect.TypeOf(refObj)
	if refType.Kind() != reflect.Struct {
		panic("buzz: reference object is not struct")
	}

	refFields := reflect.VisibleFields(refType)

	if len(refFields) != len(fields) {
		panic("buzz: reference object's field count does not match to the number of fields")
	}

	for _, field := range fields {
		fieldName := field.Name()
		fieldType := field.Type()

		found := false
		for _, reflectField := range refFields {
			if reflectField.Name == fieldName {
				if reflectField.Type != fieldType {
					panic(fmt.Sprintf("buzz: field '%s' has mismatching types", fieldName))
				}

				found = true
				break
			}
		}

		if !found {
			panic(fmt.Sprintf("buzz: field '%s' not found in the reference object", fieldName))
		}
	}

	return &BuzzSchema[T]{
		fields:  fields,
		refType: refType,
	}
}

func (s *BuzzSchema[T]) Validate(obj any) error {
	valueObj := reflect.ValueOf(obj)
	for _, f := range s.fields {
		valueField := valueObj.FieldByName(f.Name())
		if err := f.Validate(valueField.Interface()); err != nil {
			return err
		}
	}

	for _, valFn := range s.validateFuncs {
		if err := valFn(obj.(T)); err != nil {
			return err
		}
	}

	return nil
}

func (s *BuzzSchema[T]) Type() reflect.Type {
	return s.refType
}

func (s *BuzzSchema[T]) Name() string {
	return s.name
}

func (s *BuzzSchema[T]) Extend(refObj any, fields ...BuzzField) *BuzzSchema[any] {
	newFields := append(fields, s.fields...)
	return Schema(refObj, newFields...)
}

func (s *BuzzSchema[T]) Pick(refObj any, fieldNames ...string) *BuzzSchema[any] {
	var newFields []BuzzField
	for _, name := range fieldNames {
		for _, field := range s.fields {
			if field.Name() == name {
				newFields = append(newFields, field)
				break
			}
		}
	}

	return Schema(refObj, newFields...)
}

func (s *BuzzSchema[T]) Omit(refObj any, fieldNames ...string) *BuzzSchema[any] {
	var newFields []BuzzField
	for _, field := range s.fields {
		fieldName := field.Name()

		found := false
		for _, name := range fieldNames {
			if fieldName == name {
				found = true
				break
			}
		}

		if !found {
			newFields = append(newFields, field)
		}
	}

	return Schema(refObj, newFields...)
}

func (s *BuzzSchema[T]) Fields() []BuzzField {
	return s.fields
}

func (s *BuzzSchema[T]) WithName(name string) *BuzzSchema[T] {
	return &BuzzSchema[T]{
		name:    name,
		fields:  s.fields,
		refType: s.refType,
	}
}

func (s *BuzzSchema[T]) Custom(fn func(T) error) *BuzzSchema[T] {
	s.addValidateFunc(fn)
	return s
}

func (s *BuzzSchema[T]) addValidateFunc(fn BuzzSchemaValidateFunc[T]) {
	s.validateFuncs = append(s.validateFuncs, fn)
}
