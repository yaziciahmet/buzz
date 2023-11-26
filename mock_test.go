package buzz

type TestStruct struct {
	String string
}

type User struct {
	Id    int
	Name  string
	Email string
}

type Id struct {
	Id int
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

type UserWithUnexportedFields struct {
	Id   int
	name string
}

type StructWithSlice struct {
	Id    int
	List  []string
	List2 []int
}
