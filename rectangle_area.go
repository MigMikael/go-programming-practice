package main

import (
	"fmt"
)

func rectangleArea() {
	fmt.Print("Enter width height : ")
	var w, h float32
	fmt.Scan(&w, &h)
	//fmt.Printf("You input width is %d and %d", w, h)

	if w <= 0 && h <= 0 {
		fmt.Println("invalid width and height")
	} else if w <= 0 {
		fmt.Println("invalid width")
	} else if h <= 0 {
		fmt.Println("invalid height")
	} else {
		fmt.Printf("%.2f", w*h)
	}
}
