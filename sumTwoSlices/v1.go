package main

import "fmt"

func main() {
	var arr1 []int = []int{5}
	var arr2 []int = []int{4, 5, 6}

	answer := getSumTwoSlices(arr1, arr2)
	fmt.Println(answer)
}

func getSumTwoSlices(arr1, arr2 []int) []int {

	// длина результирующего слайса
	length := max(len(arr1), len(arr2)) + 1

	//создаем результирующий слайс
	resSlice := make([]int, length)

	//создаем два слайса, у которых в последствии будут длины как у результирующего для упращения складывания элементов
	slice1 := make([]int, length-len(arr1))
	slice2 := make([]int, length-len(arr2))

	//заполняем slice1
	slice1 = append(slice1, arr1...)

	//заполняем slice2
	slice2 = append(slice2, arr2...)

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
