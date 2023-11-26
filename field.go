package buzz

import "reflect"

type BuzzField interface {
	Name() string
	SetName(n string)
	Type() reflect.Type
	Validate(v any) error
	Clone() BuzzField
}

func Field(name string, field BuzzField) BuzzField {
	clone := field.Clone()
	clone.SetName(name)
	return clone
}
