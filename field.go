package buzz

import "reflect"

type BuzzField interface {
	Name() string
	SetName(n string)
	Type() reflect.Type
	Validate(v any) error
}

func Field(name string, field BuzzField) BuzzField {
	field.SetName(name)
	return field
}
