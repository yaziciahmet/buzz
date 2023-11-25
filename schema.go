package buzz

import (
	"fmt"
	"reflect"
)

type BuzzField interface {
	Name() string
	Validate(v any) error
	Type() reflect.Type
}

type BuzzSchema struct {
	name    string
	fields  []BuzzField
	refType reflect.Type
}

func Schema(name string, refObj any, fields ...BuzzField) *BuzzSchema {
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

		found := false
		for _, reflectField := range refFields {
			if reflectField.Name == fieldName {
				if reflectField.Type != field.Type() {
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

	return &BuzzSchema{
		name:    name,
		fields:  fields,
		refType: refType,
	}
}

func (s *BuzzSchema) Validate(obj any) error {
	valueObj := reflect.ValueOf(obj)
	for _, f := range s.fields {
		valueField := valueObj.FieldByName(f.Name())
		if err := f.Validate(valueField.Interface()); err != nil {
			return err
		}
	}
	return nil
}

func (s *BuzzSchema) Type() reflect.Type {
	return s.refType
}

func (s *BuzzSchema) Name() string {
	return s.name
}
