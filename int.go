package buzz

import "reflect"

var (
	intReflectType = reflect.TypeOf(0)
)

type BuzzIntValidateFunc func(v int) error

type BuzzInt struct {
	name          string
	validateFuncs []BuzzIntValidateFunc
}

func Int() *BuzzInt {
	return &BuzzInt{}
}

func (i *BuzzInt) Name() string {
	return i.name
}

func (i *BuzzInt) SetName(name string) {
	i.name = name
}

func (i *BuzzInt) Type() reflect.Type {
	return intReflectType
}

func (i *BuzzInt) Validate(v any) error {
	for _, valFn := range i.validateFuncs {
		vint, ok := v.(int)
		if !ok {
			return makeValidationError("", "type", "type not int")
		}

		if err := valFn(vint); err != nil {
			return err
		}
	}
	return nil
}

func (i *BuzzInt) Min(min int) *BuzzInt {
	i.addValidateFunc(func(v int) error {
		if min > v {
			return makeValidationError("", "min", "min failed")
		}
		return nil
	})
	return i
}

func (i *BuzzInt) Max(max int) *BuzzInt {
	i.addValidateFunc(func(v int) error {
		if max < v {
			return makeValidationError("", "max", "max failed")
		}
		return nil
	})
	return i
}

func (i *BuzzInt) Positive() *BuzzInt {
	i.addValidateFunc(func(v int) error {
		if v <= 0 {
			return makeValidationError("", "positive", "positive failed")
		}
		return nil
	})
	return i
}

func (i *BuzzInt) Nonnegative() *BuzzInt {
	i.addValidateFunc(func(v int) error {
		if v < 0 {
			return makeValidationError("", "nonnegative", "nonnegative failed")
		}
		return nil
	})
	return i
}

func (i *BuzzInt) Negative() *BuzzInt {
	i.addValidateFunc(func(v int) error {
		if v >= 0 {
			return makeValidationError("", "negative", "negative failed")
		}
		return nil
	})
	return i
}

func (i *BuzzInt) NonPositive() *BuzzInt {
	i.addValidateFunc(func(v int) error {
		if v > 0 {
			return makeValidationError("", "nonpositive", "nonpositive failed")
		}
		return nil
	})
	return i
}

func (i *BuzzInt) Custom(fn BuzzIntValidateFunc) *BuzzInt {
	i.addValidateFunc(fn)
	return i
}

func (i *BuzzInt) addValidateFunc(fn BuzzIntValidateFunc) {
	i.validateFuncs = append(i.validateFuncs, fn)
}
