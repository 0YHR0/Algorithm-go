package sorts

import "fmt"

func BubbleSort(arr *[12]int) {
	flag := false // 优化，若某次发现不用交换，则提前结束排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				flag = true
			}
		}
		if flag == false { // 如果在某一次循环中没有交换过，则直接break
			break
		} else {
			flag = false // 下次判断的初始条件
		}
	}
	fmt.Println("this is bubble sort:")
}
