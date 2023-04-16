package sorts

func QuickSort(arr *[12]int, leftBound int, rightBound int) {
	// quit condition
	if leftBound >= rightBound {
		return
	}
	// use the most right one as the pivot
	pivot := arr[rightBound]
	left := leftBound
	right := rightBound - 1
	for left <= right {
		for left <= right && arr[left] <= pivot {
			left++
		}
		for left <= right && arr[right] > pivot {
			right--
		}
		// attention! IF here
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[left], arr[rightBound] = arr[rightBound], arr[left]
	QuickSort(arr, leftBound, left-1)
	QuickSort(arr, left+1, rightBound)
}
