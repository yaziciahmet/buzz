package buzz

import (
	"errors"
	"testing"
)

func Test_SchemaBasicStructSuccess(t *testing.T) {
	if err := Schema(
		User{},
		Int("Id").Min(0).Max(1000),
		String("Name").Min(2).Max(20),
		String("Email").Email(),
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
		Int("Id").Min(0).Max(1000),
		String("Name").Min(2).Max(20),
		Schema(
			Address{},
			Int("ZipCode").Min(0).Max(100000),
			String("Text").Min(5).Max(100),
		).WithName("Address"),
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

func Test_SchemaExtendSuccess(t *testing.T) {
	if err := Extend(
		Schema(
			User{},
			Int("Id").Min(0).Max(1000),
			String("Name").Min(2).Max(20),
			String("Email").Email(),
		),
		UserExtended{},
		String("String").Min(5),
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
			Int("Id").Min(0).Max(1000),
			String("Name").Min(2).Max(20),
			String("Email").Email(),
		),
		UserExtended{},
		String("String").Min(5),
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
			Int("Id").Min(0).Max(1000),
			String("Name").Min(2).Max(20),
			String("Email").Email(),
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
			Int("Id").Min(0).Max(1000),
			String("Name").Min(2).Max(20),
			String("Email").Email(),
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
			Int("Id").Min(0).Max(1000),
			String("Name").Min(2).Max(20),
			String("Email").Email(),
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
			Int("Id").Min(0).Max(1000),
			String("Name").Min(2).Max(20),
			String("Email").Email(),
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
		Int("Id").Min(0).Max(1000),
		String("Name").Min(2).Max(20),
		String("Email").Email(),
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
			Int("Id").Min(0).Max(1000),
			String("Name").Min(2).Max(20),
			String("Email").Email(),
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
