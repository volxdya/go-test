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

	fmt.Println(p.Square())

	p.SetXY(5, 6)
	fmt.Println(p)

	fmt.Println(p.Square())

}
