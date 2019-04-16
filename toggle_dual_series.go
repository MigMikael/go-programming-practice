package main

import "fmt"

func toggleDualSeries() {
	var m, n int
	fmt.Scan(&m, &n)

	var sumFirst, sumSecond int
	var isPlus = true
	for index := 1; index <= m; index += 2 {
		//fmt.Println(index)
		if isPlus {
			sumFirst += index
			isPlus = false
		} else {
			sumFirst -= index
			isPlus = true
		}
	}
	//fmt.Println()
	isPlus = true
	for index := 1; index <= n; index *= 2 {
		//fmt.Println(index)
		if isPlus {
			sumSecond += index
			isPlus = false
		} else {
			sumSecond -= index
			isPlus = true
		}
	}

	fmt.Println(sumFirst * sumSecond)
}
