package buzz

import (
	"errors"
	"testing"
	"time"
)

func Test_SchemaBasicStructSuccess(t *testing.T) {
	if err := Schema(
		User{},
		Field("Id", Number[int]().Min(0).Max(1000)),
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
		Field("Id", Number[int]().Min(0).Max(1000)),
		Field("Name", String().Min(2).Max(20)),
		Field("Address", Schema[any](
			Address{},
			Field("ZipCode", Number[int]().Min(0).Max(100000)),
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
		Field("Id", Number[int]().Min(0).Max(1000)),
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
		Field("Id", Number[int]()),
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
			Field("Id", Number[int]().Min(0).Max(1000)),
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
			Field("Id", Number[int]().Min(0).Max(1000)),
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
			Field("Id", Number[int]().Min(0).Max(1000)),
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
			Field("Id", Number[int]().Min(0).Max(1000)),
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
			Field("Id", Number[int]().Min(0).Max(1000)),
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
			Field("Id", Number[int]().Min(0).Max(1000)),
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
		Field("Id", Number[int]().Min(0).Max(1000)),
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
			Field("Id", Number[int]().Min(0).Max(1000)),
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

func Test_SchemaStructWithPointerFieldSuccess(t *testing.T) {
	schema := Schema(
		StructWithPointer{},
		Field("Id", Number[int]()),
		Field("Name", Ptr(String().Min(4))),
	)

	name := "ahmet"

	if err := schema.Validate(StructWithPointer{
		Id:   100,
		Name: &name,
	}); err != nil {
		t.FailNow()
	}
}

func Test_SchemaComplexStruct(t *testing.T) {
	now := time.Now()

	schema := Schema(
		ComplexStruct{},
		Field("Id", Number[int]()),
		Field("Email", String()),
		Field("Spouse", Ptr(Schema(
			User{},
			Field("Id", Number[int]().Min(1)),
			Field("Email", String().Email()),
			Field("Name", String()),
		))),
		Field("Int64", Number[int64]()),
		Field("Comments", Slice[string]().Max(10)),
		Field("Friends", Slice[User]().ForEach(func(v User) error {
			if v.Id > 100 {
				return errors.New("too big id bro")
			}
			return nil
		})),
		Field("FriendsWithPtrUsers", Slice[*User]().Min(1)),
		Field("Admin", Bool().True()),
		Field("CreatedAt", Time().After(now.Add(-1))),
		Field("UpdatedAt", Ptr(Time()).Nonnil()),
		Field("LastErrorMsg", Ptr(String())),
	).Custom(func(cs ComplexStruct) error {
		if cs.LastErrorMsg != nil {
			return errors.New(*cs.LastErrorMsg)
		}
		return nil
	})

	complexStruct1 := ComplexStruct{
		Id:    100,
		Email: "ahmet@mail.com",
		Spouse: &User{
			Id:    1,
			Email: "spouse@mail.com",
			Name:  "spouse",
		},
		Int64:    241321,
		Comments: []string{"no comment"},
		Friends: []User{{
			Id: 2,
		}, {
			Id: 99,
		}},
		FriendsWithPtrUsers: []*User{{
			Id: 22,
		}},
		Admin:        true,
		CreatedAt:    now,
		UpdatedAt:    &now,
		LastErrorMsg: nil,
	}
	if err := schema.Validate(complexStruct1); err != nil {
		t.FailNow()
	}
}

func Test_SchemaReuseFieldInDifferentSchemas(t *testing.T) {
	numberField := Number[int]().Min(1)

	schema1 := Schema(
		Id{},
		Field("Id", numberField),
	)

	schema2 := Schema(
		Id2{},
		Field("Id2", numberField),
	)

	if err := schema1.Validate(Id{Id: 1}); err != nil {
		t.FailNow()
	}

	if err := schema2.Validate(Id2{Id2: 1}); err != nil {
		t.FailNow()
	}
}
