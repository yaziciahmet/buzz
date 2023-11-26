package buzz

import (
	"fmt"
	"reflect"
	"unicode"
)

type BuzzSchemaValidateFunc[T any] func(T) error

type BuzzSchema[T any] struct {
	name          string
	fields        []BuzzField
	validateFuncs []BuzzSchemaValidateFunc[T]
	refType       reflect.Type
}

func Schema[T any](fields ...BuzzField) *BuzzSchema[T] {
	refObj := *new(T)

	refType := reflect.TypeOf(refObj)
	if refType.Kind() != reflect.Struct {
		panic("buzz: reference object is not struct")
	}

	refFields := reflect.VisibleFields(refType)

	var exportedFields []BuzzField
	for _, field := range fields {
		fieldName := field.Name()
		fieldType := field.Type()

		if unicode.IsLower(rune(fieldName[0])) {
			continue
		}

		found := false
		for _, refField := range refFields {
			if refField.Name == fieldName {
				if refField.Type != fieldType {
					panic(fmt.Sprintf("buzz: field '%s' has mismatching types", fieldName))
				}

				exportedFields = append(exportedFields, field)
				found = true
				break
			}
		}

		if !found {
			panic(fmt.Sprintf("buzz: field '%s' not found in the reference object", fieldName))
		}
	}

	return &BuzzSchema[T]{
		fields:  exportedFields,
		refType: refType,
	}
}

func (s *BuzzSchema[T]) Validate(obj any) error {
	objT, ok := obj.(T)
	if !ok {
		return fmt.Errorf(invalidTypeMsg, s.refType, obj)
	}

	valueObj := reflect.ValueOf(obj)
	for _, f := range s.fields {
		valueField := valueObj.FieldByName(f.Name())
		if err := f.Validate(valueField.Interface()); err != nil {
			return err
		}
	}

	for _, valFn := range s.validateFuncs {
		if err := valFn(objT); err != nil {
			return err
		}
	}

	return nil
}

func (s *BuzzSchema[T]) Name() string {
	return s.name
}

func (s *BuzzSchema[T]) Type() reflect.Type {
	return s.refType
}

func (s *BuzzSchema[T]) WithName(name string) BuzzField {
	s.name = name
	return s
}

func (s *BuzzSchema[T]) Clone() BuzzField {
	return &BuzzSchema[T]{
		name:          s.name,
		fields:        s.fields,
		validateFuncs: s.validateFuncs,
		refType:       s.refType,
	}
}

func (s *BuzzSchema[T]) Fields() []BuzzField {
	return s.fields
}

func (s *BuzzSchema[T]) Custom(fn BuzzSchemaValidateFunc[T]) *BuzzSchema[T] {
	s.addValidateFunc(fn)
	return s
}

func (s *BuzzSchema[T]) addValidateFunc(fn BuzzSchemaValidateFunc[T]) {
	s.validateFuncs = append(s.validateFuncs, fn)
}

func Extend[T, K any](schema *BuzzSchema[K], fields ...BuzzField) *BuzzSchema[T] {
	newFields := append(fields, schema.fields...)
	return Schema[T](newFields...)
}

func Pick[T, K any](schema *BuzzSchema[K], fieldNames ...string) *BuzzSchema[T] {
	var newFields []BuzzField
	for _, name := range fieldNames {
		for _, field := range schema.fields {
			if field.Name() == name {
				newFields = append(newFields, field)
				break
			}
		}
	}

	return Schema[T](newFields...)
}

func Omit[T, K any](schema *BuzzSchema[K], fieldNames ...string) *BuzzSchema[T] {
	var newFields []BuzzField
	for _, field := range schema.fields {
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

	return Schema[T](newFields...)
}
