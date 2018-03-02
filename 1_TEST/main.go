package main

import "fmt"

// use brute force to read from std in to std out as long as input is not 42

func main() {

	var num int

	fmt.Scan(&num)

	for num != 42 {

		fmt.Println(num)
		fmt.Scan(&num)
	}
}
