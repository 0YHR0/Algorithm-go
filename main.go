package main

import (
	"algorithm-go/sorts"
	"fmt"
)

func main() {
	//fmt.Println("test")
	//arr := [...]int{9, 6, 5, 4, 8, 3, 2, 19, 19, 66, 0, -1}
	arr_2 := make([]int, 0)
	arr_2 = append(arr_2, []int{9, 6, 5, 4, 8, 3, 2, 19, 19, 66, 0, -1}...)
	//sorts.BubbleSort(&arr)
	//sorts.SelectSort(&arr)
	fmt.Println(arr_2)
	//sorts.InsertSort(&arr)
	//sorts.ShellSort(&arr)
	//sorts.QuickSort(&arr, 0, 11)
	//sorts.DoublePivotQuickSort(&arr, 0, 11)

	fmt.Println(sorts.MergeSort(arr_2))
}
