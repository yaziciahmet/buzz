package buzz

import "testing"

func Benchmark_SimpleStruct(b *testing.B) {
	schema := Schema(
		User{},
		Field("Id", Int().Min(0).Max(1000)),
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
		Field("Id", Int().Min(0).Max(1000)),
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
