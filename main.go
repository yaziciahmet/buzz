package main

type Bro struct {
	Id   int
	Sis  Sis
	Name string
}

type Sis struct {
	Lol int
}

func main() {
	sisSchema := Schema(
		"Sis",
		Sis{},
		Int("Lol").Max(15),
	)
	broSchema := Schema(
		"bro",
		Bro{},
		Int("Id").Min(15),
		String("Name").Min(0),
		sisSchema,
	)

	err := broSchema.Validate(Bro{
		Id: 17,
		Sis: Sis{
			Lol: 20,
		},
	})
	if err != nil {
		panic(err)
	}
}
