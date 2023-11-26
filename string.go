package buzz

import (
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
		return MakeFieldError("", "type", "type not string")
	}

	for _, valFn := range s.validateFuncs {
		if err := valFn(vstr); err != nil {
			return err
		}
	}
	return nil
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
			return MakeFieldError("", "min", "min failed")
		}
		return nil
	})
	return s
}

func (s *BuzzString) Max(max int) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if max < len(v) {
			return MakeFieldError("", "max", "max failed")
		}
		return nil
	})
	return s
}

func (s *BuzzString) Len(l int) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if l != len(v) {
			return MakeFieldError("", "len", "len failed")
		}
		return nil
	})
	return s
}

func (s *BuzzString) Email() *BuzzString {
	s.addValidateFunc(func(v string) error {
		if _, err := mail.ParseAddress(v); err != nil {
			return MakeFieldError("", "email", "email failed")
		}
		return nil
	})
	return s
}

func (s *BuzzString) URL() *BuzzString {
	s.addValidateFunc(func(v string) error {
		if _, err := url.ParseRequestURI(v); err != nil {
			return MakeFieldError("", "url", "url failed")
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
			return MakeFieldError("", "regex", "regex failed")
		}

		return nil
	})
	return s
}

func (s *BuzzString) Contains(str string) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if !strings.Contains(v, str) {
			return MakeFieldError("", "contains", "contains failed")
		}
		return nil
	})
	return s
}

func (s *BuzzString) StartsWith(str string) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if !strings.HasPrefix(v, str) {
			return MakeFieldError("", "startsWith", "startsWith failed")
		}
		return nil
	})
	return s
}

func (s *BuzzString) EndsWith(str string) *BuzzString {
	s.addValidateFunc(func(v string) error {
		if !strings.HasSuffix(v, str) {
			return MakeFieldError("", "endsWith", "endsWith failed")
		}
		return nil
	})
	return s
}

func (s *BuzzString) UUID() *BuzzString {
	s.addValidateFunc(func(v string) error {
		if len(v) != 36 {
			return MakeFieldError("", "uuid", "invalid uuid length")
		}

		if v[8] != '-' || v[13] != '-' || v[18] != '-' || v[23] != '-' {
			return MakeFieldError("", "uuid", "invalid uuid format")
		}

		for _, r := range v {
			if r == '-' {
				continue
			}

			if !s.isHexB(byte(r)) {
				return MakeFieldError("", "uuid", "invalid uuid format")
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
