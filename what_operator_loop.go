package main

import (
	"fmt"
)

func whatOperatorLoop() {
	var first, second, result int

	for {
		fmt.Scan(&first, &second, &result)

		if first == 0 && second == 0 && result == 0 {
			break
		}

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
}
