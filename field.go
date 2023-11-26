package buzz

import "reflect"

type BuzzField interface {
	Name() string
	Type() reflect.Type
	Validate(v any) error
	WithName(n string) BuzzField
	Clone() BuzzField
}

func Field(name string, field BuzzField) BuzzField {
	clone := field.Clone()
	return clone.WithName(name)
}
