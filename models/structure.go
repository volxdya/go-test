package models

import (
	"f/structs"
	"fmt"
)

func TestStructs() {
	p := structs.Point{
		X: 2,
		Y: 3,
	}

	user1 := structs.Developer{
		User: structs.User{
			FirstName: "John",
			LastName:  "Doe",
			Password:  "123456",
		},
		PL: "Golang",
	}

	if user1.PL == "" {
		fmt.Println("Пусто. Автмоатически заполнилось")
	} else {
		fmt.Println(user1.PL)
	}

	fmt.Println(p.Square())
	p.SetXY(5, 6)
	fmt.Println(p)
	fmt.Println(p.Square())

}
