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

func DoublePivotQuickSort(arr *[12]int, leftBound int, rightBound int) {
	if leftBound >= rightBound {
		return
	}

	if arr[leftBound] > arr[rightBound] {
		arr[leftBound], arr[rightBound] = arr[rightBound], arr[leftBound]
	}

	pivot1 := arr[leftBound]
	pivot2 := arr[rightBound]
	left := leftBound
	right := rightBound
	k := leftBound + 1

OutLoop:
	for k < right {
		if arr[k] < pivot1 {
			left++
			arr[left], arr[k] = arr[k], arr[left]
			k++
		} else if arr[k] <= pivot2 && arr[k] >= pivot1 {
			k++
		} else {
			right--
			for arr[right] > pivot2 {
				if right <= k {
					break OutLoop
				}
				right--
			}
			if arr[right] <= pivot2 && arr[right] >= pivot1 {
				arr[k], arr[right] = arr[right], arr[k]
				k++
			} else {
				arr[right], arr[k] = arr[k], arr[right]
				left++
				arr[k], arr[left] = arr[left], arr[k]
				k++
			}
		}
	}

	arr[leftBound], arr[left] = arr[left], arr[leftBound]
	arr[right], arr[rightBound] = arr[rightBound], arr[right]

	//一次双轴快排至少确定两个轴（元素）的位置
	DoublePivotQuickSort(arr, leftBound, left-1)
	DoublePivotQuickSort(arr, left+1, right-1)
	DoublePivotQuickSort(arr, right+1, rightBound)

}
