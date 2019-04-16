package main

import (
	"fmt"
)

func whatOperator() {
	var first, second, result int
	fmt.Scan(&first, &second, &result)

	if first+second == result {
		fmt.Println("+")
	} else if first-second == result {
		fmt.Println("-")
	} else if first*second == result {
		fmt.Println("*")
	} else if first/second == result {
		fmt.Println("/")
	} else if first%second == result {
		fmt.Println("%")
	}
}
