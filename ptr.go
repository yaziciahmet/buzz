package buzz

import "reflect"

type BuzzPtr struct {
	name     string
	nullable bool
	field    BuzzField
	refType  reflect.Type
}

func Ptr(field BuzzField) *BuzzPtr {
	return &BuzzPtr{
		nullable: false,
		field:    field,
		refType:  reflect.New(field.Type()).Type(),
	}
}

func (p *BuzzPtr) Name() string {
	return p.name
}

func (p *BuzzPtr) SetName(name string) {
	p.name = name
}

func (p *BuzzPtr) Type() reflect.Type {
	return p.refType
}

func (p *BuzzPtr) Validate(v any) error {
	refValue := reflect.ValueOf(v)
	if refValue.Kind() != reflect.Pointer {
		return makeValidationError("", "type", "type must be pointer")
	}

	if refValue.IsNil() {
		if p.nullable {
			return nil
		}

		return makeValidationError("", "nullable", "pointer is not nullable")
	}

	return p.field.Validate(refValue.Elem().Interface())
}

func (p *BuzzPtr) Nullable() *BuzzPtr {
	p.nullable = true
	return p
}
