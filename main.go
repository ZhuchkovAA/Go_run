package main 

import ( 
	"fmt"
 )

func main() {
	user := &user.UserType{ "Alex", "Zhuchkov", 21 }
	age := user.getAge(user)
	fmt.Println(age)
}