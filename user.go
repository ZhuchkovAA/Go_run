package user

type UserType struct {
	Firstname string
	Lastname string
	Age int
}

func getAge(user *UserType) (int) {
	return user.Age
}
