package main

import (
	"fmt"
)

func considerOrder() {
	var stockPlate, stockStick, bigTable, smallTable int
	fmt.Scan(&stockPlate, &stockStick)
	fmt.Scan(&bigTable, &smallTable)

	//fmt.Println(stockPlate, stockStick, bigTable, smallTable)
	// BigTable => 2 wood plate 6 wood stick
	// SmallTable => 1 wood plate 4 wood stick

	var plateUsed, stickUsed int
	plateUsed += (bigTable * 2) + (smallTable * 1)
	stickUsed += (bigTable * 6) + (smallTable * 4)

	plateLeft := stockPlate - plateUsed
	stickLeft := stockStick - stickUsed

	if plateLeft < 0 || stickLeft < 0 {
		fmt.Print("No ")
		if plateLeft < 0 {
			fmt.Printf("%d ", plateLeft*-1)
		} else {
			fmt.Print("0 ")
		}

		if stickLeft < 0 {
			fmt.Printf("%d ", stickLeft*-1)
		} else {
			fmt.Print("0")
		}
	} else {
		fmt.Println("Yes", plateLeft, stickLeft)
	}
}
