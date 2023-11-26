package buzz

import "time"

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

type Id2 struct {
	Id2 int
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

type StructWithPointer struct {
	Id   int
	Name *string
}

type ComplexStruct struct {
	Id                  int
	Email               string
	Spouse              *User
	Int64               int64
	Comments            []string
	Friends             []User
	FriendsWithPtrUsers []*User
	Admin               bool
	CreatedAt           time.Time
	UpdatedAt           *time.Time
	LastError           *string
}

type MyError struct{}

func (e *MyError) Error() string {
	return "error brother"
}

type MyInterface interface {
	Method() int
}

type MyInterfaceStruct struct{}

func (s *MyInterfaceStruct) Method() int { return 5 }
