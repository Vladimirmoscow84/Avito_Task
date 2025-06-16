package main

import "fmt"

func main() {
	var arr1 []int = []int{6, 0, 0}
	var arr2 []int = []int{4, 5, 6}
	if len(arr1) == 0 {
		fmt.Println(arr2)
		return
	} else if len(arr2) == 0 {
		fmt.Println(arr1)
		return
	}

	answer := getSumTwoSlices(arr1, arr2)
	fmt.Println(answer)
}

func getSumTwoSlices(arr1, arr2 []int) []int {

	// длина результирующего слайса
	length := max(len(arr1), len(arr2)) + 1

	//емкость результирующего слайса
	capacity := length

	//создаем результирующий слайс
	resSlice := make([]int, length, capacity)

	//создаем два слайса на основе заданных, но с длиной, равной результирующему, а значения по лишним индексам принимаем 0
	slice1 := make([]int, length-len(arr1))
	slice2 := make([]int, length-len(arr2))
	fmt.Println(slice1)
	fmt.Println(slice2)

	//заполняем slice1
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 0
	}
	slice1 = append(slice1, arr1...)
	fmt.Println(slice1)

	//заполняем slice2
	for i := 0; i < len(slice2); i++ {
		slice2[i] = 0
	}
	slice2 = append(slice2, arr2...)
	fmt.Println(slice2)

	//заполняем результирующий слайс
	n := 0
	for i := len(resSlice) - 1; i >= 0; i-- {
		resSlice[i] = (slice1[i] + slice2[i] + n) % 10
		if slice1[i]+slice2[i]+n > 9 {
			n = 1
		} else {
			n = 0
		}

	}
	return resSlice
}
