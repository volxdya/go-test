package main

import (
	"fmt"
	"math"
)

// test1
func main() {
	for i := 0; true; i++ {
		var h, w float64
		fmt.Println("Input ur height: ")
		_, err := fmt.Scan(&h)

		if err != nil {
			fmt.Printf("I NEED A NUMBER, U WROTE NOT A NUMBER")
		}

		fmt.Println("Input ur height: ")
		_, err = fmt.Scan(&w)
		if err != nil {
			fmt.Printf("I NEED A NUMBER, U WROTE NOT A NUMBER")
		}

		indexMassBody := w / math.Pow(h, 2)

		fmt.Println(indexMassBody)
	}
}
