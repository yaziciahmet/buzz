package buzz

import "testing"

func Benchmark_SimpleStruct(b *testing.B) {
	schema := Schema(
		User{},
		Int("Id").Min(0).Max(1000),
		String("Name").Min(2).Max(20),
		String("Email").Email(),
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
