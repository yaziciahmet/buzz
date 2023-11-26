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

func Schema[T any](refObj T, fields ...BuzzField) *BuzzSchema[T] {
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
					fmt.Println(refField.Type, fieldType)
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
	valueObj := reflect.ValueOf(obj)
	for _, f := range s.fields {
		valueField := valueObj.FieldByName(f.Name())
		if err := f.Validate(valueField.Interface()); err != nil {
			return err
		}
	}

	for _, valFn := range s.validateFuncs {
		objT, ok := obj.(T)
		if !ok {
			return makeValidationError("", "type", "type not T")
		}

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

func Extend[T, K any](schema *BuzzSchema[K], refObj T, fields ...BuzzField) *BuzzSchema[T] {
	newFields := append(fields, schema.fields...)
	return Schema(refObj, newFields...)
}

func Pick[T, K any](schema *BuzzSchema[K], refObj T, fieldNames ...string) *BuzzSchema[T] {
	var newFields []BuzzField
	for _, name := range fieldNames {
		for _, field := range schema.fields {
			if field.Name() == name {
				newFields = append(newFields, field)
				break
			}
		}
	}

	return Schema(refObj, newFields...)
}

func Omit[T, K any](schema *BuzzSchema[K], refObj T, fieldNames ...string) *BuzzSchema[T] {
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

	return Schema(refObj, newFields...)
}
