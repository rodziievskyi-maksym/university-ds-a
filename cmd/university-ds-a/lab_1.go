package main

import (
	"fmt"
	"log"
)

var (
	n, m, algoNum int
	allowedMinNum = 0
	allowedMaxNum = 1000
	amountOfAlgos = 2
)

func RunLab1() {
	ScanUserInput()

	if !ValidateUserInput(n, m) {
		log.Fatalf("Please enter nums within allowed range from %d to %d", allowedMinNum, allowedMaxNum)
	}

	if !ValidateUserAlgoInput(algoNum) {
		log.Fatalf("Please enter a valid algo num")
	}

	switch algoNum {
	case 1:
		fmt.Println(GCDEuclid(n, m))
	case 2:
		fmt.Println(GCDSequential(n, m))
	}
}

func ScanUserInput() {
	fmt.Println("Enter n:")
	fmt.Scan(&n)

	fmt.Println("Enter m:")
	fmt.Scan(&m)

	fmt.Println("Enter num of algorithm \n 1: Euclid \n 2: Sequential Search")
	fmt.Scan(&algoNum)
}

func ValidateUserInput(n, m int) bool {
	if n < allowedMinNum || n > allowedMaxNum {
		return false
	}

	if m < allowedMinNum || m > allowedMaxNum {
		return false
	}

	return true
}

func ValidateUserAlgoInput(algoNum int) bool {
	if algoNum <= 0 || algoNum > amountOfAlgos {
		return false
	}

	return true
}

func GCDEuclid(n, m int) int {
	//Обчислення НСД чисел m і n за допомогою алгоритму Евкліда

	if n == 0 {
		return m
	}

	iterations := 0

	for n != 0 {
		r := m % n
		m = n
		n = r

		iterations++
	}

	fmt.Println("iterations:", iterations)
	return m
}

func GCDSequential(n, m int) int {
	//Обчислення НСД чисел m і n за допомогою послідовного перебору з кінця
	iterations := 0

	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	minNum := m
	if n < m {
		minNum = n
	}

	for i := minNum; i >= 1; i-- {
		if m%i == 0 && n%i == 0 {
			iterations++
			fmt.Println("iterations:", iterations)
			return i
		}
		iterations++
	}

	return 1
}
