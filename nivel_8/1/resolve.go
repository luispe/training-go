package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u1 := User{
		Name: "Homer",
		Age:  42,
	}

	u2 := User{
		Name: "Barnie",
		Age:  39,
	}

	u3 := User{
		Name: "Moe",
		Age:  54,
	}

	Users := []User{u1, u2, u3}

	b, err := json.Marshal(Users)
	if err != nil {
		fmt.Printf(err.Error(), err)
		return
	}

	fmt.Println(string(b))
}
