package main

import (
	"fmt"
	"log"
)

func RunLab1() {
	var n, m int
	fmt.Println("Enter n and m:")
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		log.Fatal(err)
	}

	if m == 0 && n == 0 {
		log.Fatal("Both m and n are zero")
	}

	if m < 0 {
		m = -m
	}
	if n < 0 {
		n = -n
	}

	var algorithm string
	fmt.Println("Please enter the algorithm [euclid, sequential]: ")
	_, err = fmt.Scan(&algorithm)
	if err != nil {
		log.Fatal(err)
	}

	switch algorithm {
	case "euclid":
		fmt.Println(GCDEuclid(n, m))
	case "sequential":
		fmt.Println(GCDSequential(n, m))
	}
}

func GCDEuclid(n, m int) int {
	//Обчислення НСД чисел m і n за допомогою алгоритму Евкліда.®

	if n == 0 {
		fmt.Println("GCDEuclid is zero")
		return m
	}

	for n != 0 {
		r := m % n
		m = n
		n = r
	}

	return m
}

func GCDSequential(n, m int) int {
	//Обчислення НСД чисел m і n за допомогою послідовного перебору

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
			return i
		}
	}

	return 1
}
