package main

import "fmt"

func woodUsed() {
	var bigTable, smallTable int
	fmt.Scan(&bigTable, &smallTable)

	// BigTable => 2 wood plate 6 wood stick
	// SmallTable => 1 wood plate 4 wood stick
	var woodPlate, woodStick int
	woodPlate += (bigTable * 2) + (smallTable * 1)
	woodStick += (bigTable * 6) + (smallTable * 4)

	fmt.Println(woodPlate, woodStick)
}
