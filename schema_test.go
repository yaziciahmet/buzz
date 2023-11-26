package buzz

import (
	"errors"
	"testing"
)

func Test_SchemaBasicStructSuccess(t *testing.T) {
	if err := Schema(
		User{},
		Field("Id", Int().Min(0).Max(1000)),
		Field("Name", String().Min(2).Max(20)),
		Field("Email", String().Email()),
	).Validate(User{
		Id:    100,
		Name:  "ah",
		Email: "ahmet@mail.com",
	}); err != nil {
		t.FailNow()
	}
}

func Test_SchemaNestedStructSuccess(t *testing.T) {
	if err := Schema(
		UserWithAddress{},
		Field("Id", Int().Min(0).Max(1000)),
		Field("Name", String().Min(2).Max(20)),
		Field("Address", Schema[any](
			Address{},
			Field("ZipCode", Int().Min(0).Max(100000)),
			Field("Text", String().Min(5).Max(100)),
		)),
	).Validate(UserWithAddress{
		Id:   1000,
		Name: "ah",
		Address: Address{
			ZipCode: 33333,
			Text:    "Istanbul",
		},
	}); err != nil {
		t.FailNow()
	}
}

func Test_SchemaStructWithUnexportedFields(t *testing.T) {
	if err := Schema(
		UserWithUnexportedFields{},
		Field("Id", Int().Min(0).Max(1000)),
	).Validate(UserWithUnexportedFields{
		Id:   100,
		name: "gonna be ignored brother",
	}); err != nil {
		t.FailNow()
	}
}

func Test_SchemaStructWithSlice(t *testing.T) {
	if err := Schema(
		StructWithSlice{},
		Field("Id", Int()),
		Field("List", Slice[string]().Nonempty()),
		Field("List2", Slice[int]().ForEach(func(v int) error {
			if v > 5 {
				return errors.New("you shall not pass")
			}
			return nil
		})),
	).Validate(StructWithSlice{
		Id:    100,
		List:  []string{"brotherhood"},
		List2: []int{1, 2, 3},
	}); err != nil {
		t.FailNow()
	}
}

func Test_SchemaExtendSuccess(t *testing.T) {
	if err := Extend(
		Schema(
			User{},
			Field("Id", Int().Min(0).Max(1000)),
			Field("Name", String().Min(2).Max(20)),
			Field("Email", String().Email()),
		),
		UserExtended{},
		Field("String", String().Min(5)),
	).Validate(UserExtended{
		Id:     1000,
		Name:   "ragnarok",
		Email:  "ahmet@gmail.com",
		String: "brother",
	}); err != nil {
		t.FailNow()
	}
}

func Test_SchemaExtendFail(t *testing.T) {
	if err := Extend(
		Schema(
			User{},
			Field("Id", Int().Min(0).Max(1000)),
			Field("Name", String().Min(2).Max(20)),
			Field("Email", String().Email()),
		),
		UserExtended{},
		Field("String", String().Min(5)),
	).Validate(UserExtended{
		Id:     10000,
		Name:   "ragnarok",
		Email:  "ahmet@gmail.com",
		String: "bro",
	}); err == nil {
		t.FailNow()
	}
}

func Test_SchemaPickSuccess(t *testing.T) {
	if err := Pick(
		Schema(
			User{},
			Field("Id", Int().Min(0).Max(1000)),
			Field("Name", String().Min(2).Max(20)),
			Field("Email", String().Email()),
		),
		Id{},
		"Id",
	).Validate(Id{
		Id: 100,
	}); err != nil {
		t.FailNow()
	}
}

func Test_SchemaPickFail(t *testing.T) {
	if err := Pick(
		Schema(
			User{},
			Field("Id", Int().Min(0).Max(1000)),
			Field("Name", String().Min(2).Max(20)),
			Field("Email", String().Email()),
		),
		Id{},
		"Id",
	).Validate(Id{
		Id: 10000,
	}); err == nil {
		t.FailNow()
	}
}

func Test_SchemaOmitSuccess(t *testing.T) {
	if err := Omit(
		Schema(
			User{},
			Field("Id", Int().Min(0).Max(1000)),
			Field("Name", String().Min(2).Max(20)),
			Field("Email", String().Email()),
		),
		Id{},
		"Name",
		"Email",
	).Validate(Id{
		Id: 100,
	}); err != nil {
		t.FailNow()
	}
}

func Test_SchemaOmitFail(t *testing.T) {
	if err := Omit(
		Schema(
			User{},
			Field("Id", Int().Min(0).Max(1000)),
			Field("Name", String().Min(2).Max(20)),
			Field("Email", String().Email()),
		),
		Id{},
		"Name",
		"Email",
	).Validate(Id{
		Id: 10000,
	}); err == nil {
		t.FailNow()
	}
}

func Test_SchemaCustomSuccess(t *testing.T) {
	if err := Schema(
		User{},
		Field("Id", Int().Min(0).Max(1000)),
		Field("Name", String().Min(2).Max(20)),
		Field("Email", String().Email()),
	).Custom(func(u User) error {
		if u.Email != "ahmet@mail.com" {
			return errors.New("you shall not pass")
		}
		return nil
	}).Validate(User{
		Id:    100,
		Name:  "ahmet",
		Email: "ahmet@mail.com",
	}); err != nil {
		t.FailNow()
	}
}

func Test_SchemaCustomOnPickFail(t *testing.T) {
	if err := Pick(
		Schema(
			User{},
			Field("Id", Int().Min(0).Max(1000)),
			Field("Name", String().Min(2).Max(20)),
			Field("Email", String().Email()),
		),
		Id{},
		"Id",
	).Custom(func(id Id) error {
		if id.Id != 1000 {
			return errors.New("you shall not pass")
		}
		return nil
	}).Validate(Id{
		Id: 100,
	}); err == nil {
		t.FailNow()
	}
}
