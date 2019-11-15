package main

import "fmt"

func main() {
	fmt.Println(onePass(76, 4)) // 1030
}

func onePass(number int, numSys int) []int {
	var result []int
	for number >= numSys {
		result = append([]int{number % numSys}, result...)
		number = number / numSys
	}

	result = append([]int{number}, result...)
	return result
}
