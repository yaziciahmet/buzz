package buzz

import (
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
	if err := Schema(
		User{},
		Int("Id").Min(0).Max(1000),
		String("Name").Min(2).Max(20),
		String("Email").Email(),
	).Extend(
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
	if err := Schema(
		User{},
		Int("Id").Min(0).Max(1000),
		String("Name").Min(2).Max(20),
		String("Email").Email(),
	).Extend(
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
	if err := Schema(
		User{},
		Int("Id").Min(0).Max(1000),
		String("Name").Min(2).Max(20),
		String("Email").Email(),
	).Pick(
		Id{},
		"Id",
	).Validate(Id{
		Id: 100,
	}); err != nil {
		t.FailNow()
	}
}

func Test_SchemaPickFail(t *testing.T) {
	if err := Schema(
		User{},
		Int("Id").Min(0).Max(1000),
		String("Name").Min(2).Max(20),
		String("Email").Email(),
	).Pick(
		Id{},
		"Id",
	).Validate(Id{
		Id: 10000,
	}); err == nil {
		t.FailNow()
	}
}

func Test_SchemaOmitSuccess(t *testing.T) {
	if err := Schema(
		User{},
		Int("Id").Min(0).Max(1000),
		String("Name").Min(2).Max(20),
		String("Email").Email(),
	).Omit(
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
	if err := Schema(
		User{},
		Int("Id").Min(0).Max(1000),
		String("Name").Min(2).Max(20),
		String("Email").Email(),
	).Omit(
		Id{},
		"Name",
		"Email",
	).Validate(Id{
		Id: 10000,
	}); err == nil {
		t.FailNow()
	}
}
