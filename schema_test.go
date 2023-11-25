package buzz

import (
	"testing"
)

func Test_SchemaBasicStructSuccess(t *testing.T) {
	Schema(
		"User",
		User{},
		Int("Id").Min(0).Max(1000),
		String("Name").Min(2).Max(20),
		String("Email").Email(),
	).Validate(User{
		Id:    100,
		Name:  "ah",
		Email: "ahmet@mail.com",
	})
}

func Test_SchemaNestedStructSuccess(t *testing.T) {
	if err := Schema(
		"UserWithAddress",
		UserWithAddress{},
		Int("Id").Min(0).Max(1000),
		String("Name").Min(2).Max(20),
		Schema(
			"Address",
			Address{},
			Int("ZipCode").Min(0).Max(100000),
			String("Text").Min(5).Max(100),
		),
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
		"User",
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
		"User",
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
