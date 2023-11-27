package buzz

import (
	"fmt"
	"net/mail"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

var (
	stringReflectType = reflect.TypeOf("")
)

type BuzzStringValidateFunc func(string) error

type BuzzString struct {
	name          string
	validateFuncs []BuzzStringValidateFunc
}

func String() *BuzzString {
	return &BuzzString{}
}

func (s *BuzzString) Name() string {
	return s.name
}

func (s *BuzzString) Type() reflect.Type {
	return stringReflectType
}

func (s *BuzzString) Validate(v any) error {
	vstr, ok := v.(string)
	if !ok {
		return fmt.Errorf(invalidTypeMsg, stringReflectType, v)
	}

	errAggr := NewFieldErrorAggregator()
	for _, valFn := range s.validateFuncs {
		if err := valFn(vstr); err != nil {
			if errAggr.Handle(err) != nil {
				return err
			}
		}
	}

	return errAggr.OrNil()
}

func (s *BuzzString) WithName(name string) BuzzField {
	s.name = name
	return s
}

func (s *BuzzString) Clone() BuzzField {
	return &BuzzString{
		name:          s.name,
		validateFuncs: s.validateFuncs,
	}
}

func (s *BuzzString) Min(min int) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if min > len(v) {
			return MakeFieldError(s.name, "Min", fmt.Sprintf("%s must have minimum %d characters", s.name, min))
		}
		return nil
	})
	return s
}

func (s *BuzzString) Max(max int) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if max < len(v) {
			return MakeFieldError(s.name, "Max", fmt.Sprintf("%s must have maximum %d characters", s.name, max))
		}
		return nil
	})
	return s
}

func (s *BuzzString) Len(l int) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if l != len(v) {
			return MakeFieldError(s.name, "Len", fmt.Sprintf("%s must have %d characters", s.name, l))
		}
		return nil
	})
	return s
}

func (s *BuzzString) Email() *BuzzString {
	s.addValidateFunc(func(v string) error {
		if _, err := mail.ParseAddress(v); err != nil {
			return MakeFieldError(s.name, "Email", "invalid email")
		}
		return nil
	})
	return s
}

func (s *BuzzString) URL() *BuzzString {
	s.addValidateFunc(func(v string) error {
		if _, err := url.ParseRequestURI(v); err != nil {
			return MakeFieldError(s.name, "URL", "invalid url")
		}
		return nil
	})
	return s
}

func (s *BuzzString) Regex(regex string) *BuzzString {
	r, err := regexp.Compile(regex)
	s.addValidateFunc(func(v string) error {
		if err != nil {
			return err
		}

		if !r.MatchString(v) {
			return MakeFieldError(s.name, "Regex", fmt.Sprintf("regex does not satisfy %s", s.name))
		}

		return nil
	})
	return s
}

func (s *BuzzString) Contains(str string) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if !strings.Contains(v, str) {
			return MakeFieldError(s.name, "Contains", fmt.Sprintf("%s must contain %s", s.name, str))
		}
		return nil
	})
	return s
}

func (s *BuzzString) StartsWith(str string) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if !strings.HasPrefix(v, str) {
			return MakeFieldError(s.name, "StartsWith", fmt.Sprintf("%s must start with %s", s.name, str))
		}
		return nil
	})
	return s
}

func (s *BuzzString) EndsWith(str string) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if !strings.HasSuffix(v, str) {
			return MakeFieldError(s.name, "EndsWith", fmt.Sprintf("%s must end with %s", s.name, str))
		}
		return nil
	})
	return s
}

func (s *BuzzString) UUID() *BuzzString {
	s.addValidateFunc(func(v string) error {
		if len(v) != 36 {
			return MakeFieldError(s.name, "UUID", fmt.Sprintf("%s has invalid length", s.name))
		}

		if v[8] != '-' || v[13] != '-' || v[18] != '-' || v[23] != '-' {
			return MakeFieldError(s.name, "UUID", fmt.Sprintf("%s has invalid UUID format", s.name))
		}

		for _, r := range v {
			if r == '-' {
				continue
			}

			if !s.isHexB(byte(r)) {
				return MakeFieldError(s.name, "UUID", fmt.Sprintf("%s has invalid UUID format", s.name))
			}
		}

		return nil
	})
	return s
}

func (s *BuzzString) Custom(fn BuzzStringValidateFunc) *BuzzString {
	s.addValidateFunc(fn)
	return s
}

func (s *BuzzString) addValidateFunc(fn BuzzStringValidateFunc) {
	s.validateFuncs = append(s.validateFuncs, fn)
}

func (s *BuzzString) isHexB(b byte) bool {
	return (b >= 48 && b <= 57) || (b >= 65 && b <= 70) || (b >= 97 && b <= 102)
}
