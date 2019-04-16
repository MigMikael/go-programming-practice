package main

import "fmt"
import "os"

func tableProductionDay() {
	var workerBigTable, workerSmallTable, bigTableOrder, smallTableOrder int
	fmt.Scan(&workerBigTable, &workerSmallTable)
	fmt.Scan(&bigTableOrder, &smallTableOrder)

	// worker BigTable 6 table/day
	// worker SmallTable 10 table/day
	bigTableManPower := workerBigTable * 6
	smallTableManPower := workerSmallTable * 10

	if bigTableManPower == 0 && bigTableOrder > 0 {
		fmt.Println("Unable to finish order")
		os.Exit(0)
	}
	if smallTableManPower == 0 && smallTableOrder > 0 {
		fmt.Println("Unable to finish order")
		os.Exit(0)
	}

	var bigTableFinishDay int
	if bigTableManPower > 0 {
		bigTableFinishDay = bigTableOrder / bigTableManPower
		if bigTableOrder%bigTableManPower > 0 {
			bigTableFinishDay++
		}
	} else {
		bigTableFinishDay = 0
	}

	var smallTableFinishDay int
	if smallTableManPower > 0 {
		smallTableFinishDay = smallTableOrder / smallTableManPower
		if smallTableOrder%smallTableManPower > 0 {
			smallTableFinishDay++
		}
	} else {
		smallTableFinishDay = 0
	}

	if bigTableFinishDay >= smallTableFinishDay {
		fmt.Println(bigTableFinishDay)
	} else {
		fmt.Println(smallTableFinishDay)
	}
}
