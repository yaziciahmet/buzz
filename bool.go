package buzz

import (
	"fmt"
	"reflect"
)

var (
	boolReflectType = reflect.TypeOf(false)
)

type BuzzBoolValidateFunc func(bool) error

type BuzzBool struct {
	name          string
	validateFuncs []BuzzBoolValidateFunc
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
		return fmt.Errorf(invalidTypeMsg, boolReflectType, v)
	}

	for _, valFn := range b.validateFuncs {
		if err := valFn(vBool); err != nil {
			return err
		}
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
		validateFuncs: b.validateFuncs,
	}
}

func (b *BuzzBool) True() *BuzzBool {
	b.addValidateFunc(func(v bool) error {
		if v {
			return nil
		}
		return MakeFieldError("", "true", "expected true")
	})
	return b
}

func (b *BuzzBool) False() *BuzzBool {
	b.addValidateFunc(func(v bool) error {
		if !v {
			return nil
		}
		return MakeFieldError("", "false", "expected false")
	})
	return b
}

func (b *BuzzBool) Custom(fn BuzzBoolValidateFunc) *BuzzBool {
	b.addValidateFunc(fn)
	return b
}

func (b *BuzzBool) addValidateFunc(fn BuzzBoolValidateFunc) {
	b.validateFuncs = append(b.validateFuncs, fn)
}
