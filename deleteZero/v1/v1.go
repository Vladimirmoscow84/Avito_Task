package main

import "fmt"

func main() {

	slice := []int{3, 5, 0, 4, 0, 0, 7, 5}
	fmt.Println(run(slice))

}

func run(slice []int) []int {
	var sl1, sl2 []int

	for _, v := range slice {
		if v == 0 {
			sl1 = append(sl1, 0)
		} else {
			sl2 = append(sl2, v)
		}
	}
	return append(sl2, sl1...)

}
