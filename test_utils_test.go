package buzz

type TestStruct struct {
	String string
}

type User struct {
	Id    int
	Name  string
	Email string
}

type UserExtended struct {
	Id     int
	Name   string
	Email  string
	String string
}

type UserWithAddress struct {
	Id      int
	Name    string
	Address Address
}

type Address struct {
	ZipCode int
	Text    string
}