package buzz

import "reflect"

var (
	boolReflectType = reflect.TypeOf(false)
)

type BuzzBool struct {
	name          string
	expectedValue *bool
}

func Bool() *BuzzBool {
	return &BuzzBool{}
}

func (b *BuzzBool) Name() string {
	return b.name
}

func (b *BuzzBool) Type() reflect.Type {
	return boolReflectType
}

func (b *BuzzBool) Validate(v any) error {
	vBool, ok := v.(bool)
	if !ok {
		return makeValidationError("", "type", "expected bool type")
	}

	if b.expectedValue == nil {
		return nil
	}

	if vBool != *b.expectedValue {
		return makeValidationError("", "value", "expected different bool value")
	}

	return nil
}

func (b *BuzzBool) WithName(name string) BuzzField {
	b.name = name
	return b
}

func (b *BuzzBool) Clone() BuzzField {
	return &BuzzBool{
		name:          b.name,
		expectedValue: b.expectedValue,
	}
}

func (b *BuzzBool) True() *BuzzBool {
	b.expectedValue = new(bool)
	*b.expectedValue = true
	return b
}

func (b *BuzzBool) False() *BuzzBool {
	b.expectedValue = new(bool)
	*b.expectedValue = false
	return b
}
