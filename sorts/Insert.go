package sorts

func InsertSort(arr *[12]int) {
	for i := 1; i < len(arr); i++ {
		insertIndex := i
		j := 0
		// find the position in the sorted array
		for j = 0; j < i; j++ {
			if arr[insertIndex] < arr[j] {
				//find the postion j
				break
			}
		}
		insertNum := arr[insertIndex]
		// move all the larger ones to the right
		for k := insertIndex; k > j; k-- {
			arr[k] = arr[k-1]
		}
		arr[j] = insertNum
	}
}
