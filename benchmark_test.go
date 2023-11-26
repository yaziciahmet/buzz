package buzz

import (
	"errors"
	"testing"
	"time"
)

func Benchmark_SimpleStruct(b *testing.B) {
	schema := Schema(
		User{},
		Field("Id", Number[int]().Min(0).Max(1000)),
		Field("Name", String().Min(2).Max(20)),
		Field("Email", String().Email()),
	)

	user := User{
		Id:    100,
		Name:  "ahmet",
		Email: "ahmet@gmail.com",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		schema.Validate(user)
	}
}

func Benchmark_SimpleStructParallel(b *testing.B) {
	schema := Schema(
		User{},
		Field("Id", Number[int]().Min(0).Max(1000)),
		Field("Name", String().Min(2).Max(20)),
		Field("Email", String().Email()),
	)

	user := User{
		Id:    100,
		Name:  "ahmet",
		Email: "ahmet@gmail.com",
	}

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			schema.Validate(user)
		}
	})
}

func Benchmark_ComplexStruct(b *testing.B) {
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
		Field("CreatedAt", Time().After(now.Add(-1))),
		Field("UpdatedAt", Ptr(Time())),
		Field("LastError", Ptr(String()).Nullable()),
	).Custom(func(cs ComplexStruct) error {
		if cs.LastError != nil {
			return errors.New(*cs.LastError)
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
		CreatedAt: now,
		UpdatedAt: &now,
		LastError: nil,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		schema.Validate(complexStruct1)
	}
}
