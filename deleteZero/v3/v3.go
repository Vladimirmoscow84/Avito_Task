package main

import "fmt"

func main() {

	slice := []int{3, 5, 0, 4, 0, 0, 7, 5}
	fmt.Println(run(slice))

}

func run(slice []int) []int {
	nonZeroIndex := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] != 0 {
			slice[nonZeroIndex] = slice[i]
			nonZeroIndex++
		}
	}
	for i := nonZeroIndex; i < len(slice); i++ {
		slice[i] = 0
	}
	return slice

}
