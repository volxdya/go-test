package models

import (
	"fmt"
)

func TestTypes() {
	testBool := true
	testInt := 15

	// два способа объявления массива (фиксированная длина)
	testArray := [4]int{1, 2, 3, 4}
	testArray2 := [...]int{1, 2, 34, 5}

	// слайсы (динамический массив)

	/*
		Вместимость слайса - 3,
		Мы добавляем до вместимости, далее у нас расширяется до 6, под капотом наш слайс copy -> в новый слайс, где будет вместимость в 2 раза больше
		Параметры make(интерабле, длина, вместимость)
	*/
	s2 := make([]int, 2, 3)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	s2[0] = 10
	s2 = append(s2, 5)
	s2 = append(s2, 5)

	fmt.Println(s2)

	type Person struct {
		Name    string
		Age     int
		Address string
	}

	var person1 = Person{Name: "John", Age: 25, Address: "Doe"}

	//for i := 1; i <= len(testArray); i++ {
	//	fmt.Println(testArray[i])
	//}

	fmt.Println(testInt, testBool, testArray, person1, testArray2)
}
