package buzz

import (
	"reflect"
)

type BuzzPtr struct {
	name     string
	field    BuzzField
	refType  reflect.Type
	nullable bool
}

func Ptr(field BuzzField) *BuzzPtr {
	return &BuzzPtr{
		field:    field,
		refType:  reflect.New(field.Type()).Type(),
		nullable: true,
	}
}

func (p *BuzzPtr) Name() string {
	return p.name
}

func (p *BuzzPtr) Type() reflect.Type {
	return p.refType
}

func (p *BuzzPtr) Validate(v any) error {
	refValue := reflect.ValueOf(v)
	refKind := refValue.Kind()

	if refKind == reflect.Invalid {
		if p.nullable {
			return nil
		}

		return MakeFieldError("", "nonnil", "pointer is not nullable")
	}

	if refKind != reflect.Pointer {
		return MakeFieldError("", "type", "type must be pointer")
	}

	if refValue.IsNil() {
		if p.nullable {
			return nil
		}

		return MakeFieldError("", "nonnil", "pointer is not nullable")
	}

	return p.field.Validate(refValue.Elem().Interface())
}

func (p *BuzzPtr) WithName(name string) BuzzField {
	p.name = name
	return p
}

func (p *BuzzPtr) Clone() BuzzField {
	return &BuzzPtr{
		name:     p.name,
		field:    p.field,
		refType:  p.refType,
		nullable: p.nullable,
	}
}

func (p *BuzzPtr) Nonnil() *BuzzPtr {
	p.nullable = false
	return p
}
