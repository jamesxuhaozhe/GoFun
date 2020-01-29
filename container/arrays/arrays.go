package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [5]int{1, 4, 6, 9, 10}
	arr3 := [...]int{2, 4, 5, 6, 7, 8, 8}
	var grid [4][5]bool

	fmt.Println(arr1, arr2, arr3, grid)

	printArray(arr1)
	printArray(arr2)

	printArrayByRef(&arr1)
	fmt.Println(arr1)
}

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrayByRef(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
