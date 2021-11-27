package main

import (
	"fmt"
)

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func findMaxValue(arr []int) int {
	maxValue := arr[0]
	for _, i := range arr {
		if i > maxValue {
			maxValue = i
		}
	}
	return maxValue
}

func main() {
	var myArr = []int{7, 1, 5}
	var s = make([]int, len(myArr))
	fmt.Println("nilai terbesarnya", findMaxValue(myArr), "dari", myArr)
	copy(s, myArr)
	reverseArray(myArr)
	s = append(s, myArr...)
	fmt.Println(s)
}
