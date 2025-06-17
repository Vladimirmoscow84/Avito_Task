//решение без сохранения порядка!!!!!

package main

import "fmt"

func main() {

	slice := []int{3, 5, 0, 4, 0, 0, 7, 9}
	fmt.Println(run(slice))

}

//решение без сохранения порядка!!!!!

func run(slice []int) []int {
	left := 0
	right := len(slice) - 1

	for left <= right {
		//if slice[left] == 0 {
		for left < right && slice[right] == 0 {
			right--
		}
		if slice[left] == 0 {
			slice[left], slice[right] = slice[right], slice[left]
			right--
		}
		//}
		left++
	}
	return slice

}

//решение без сохранения порядка!!!!!
//решение без сохранения порядка!!!!!
//решение без сохранения порядка!!!!!
