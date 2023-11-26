package buzz

import (
	"reflect"
	"time"
)

type Condition interface {
	*BuzzInt | *BuzzString | *BuzzSlice[any] | *BuzzTimestamp | *BuzzSchema[any]
}

type BuzzField struct {
	buzzInt       *BuzzInt
	buzzString    *BuzzString
	buzzSlice     *BuzzSlice[any]
	buzzTimestamp *BuzzTimestamp
	buzzSchema    *BuzzSchema[any]

	name    string
	refType reflect.Type
}

func Field[T Condition](name string, condition T) *BuzzField {
	field := &BuzzField{
		name: name,
	}

	switch f := any(condition).(type) {
	case *BuzzInt:
		field.buzzInt = f
		field.refType = f.Type()
		break
	case *BuzzString:
		field.buzzString = f
		field.refType = f.Type()
		break
	case *BuzzSlice[any]:
		field.buzzSlice = f
		field.refType = f.Type()
		break
	case *BuzzTimestamp:
		field.buzzTimestamp = f
		field.refType = f.Type()
		break
	case *BuzzSchema[any]:
		field.buzzSchema = f
		field.refType = f.Type()
	}

	return field
}

func (f *BuzzField) Name() string {
	return f.name
}

func (f *BuzzField) Type() reflect.Type {
	return f.refType
}

func (f *BuzzField) Validate(v any) error {
	var err error
	switch {
	case f.buzzInt != nil:
		err = f.buzzInt.Validate(v.(int))
		break
	case f.buzzString != nil:
		err = f.buzzString.Validate(v.(string))
		break
	case f.buzzSlice != nil:
		err = f.buzzSlice.Validate(v.([]any))
		break
	case f.buzzTimestamp != nil:
		err = f.buzzTimestamp.Validate(v.(time.Time))
		break
	case f.buzzSchema != nil:
		err = f.buzzSchema.Validate(v)
	}

	return err
}
