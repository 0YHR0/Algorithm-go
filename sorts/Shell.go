package sorts

func ShellSort(arr *[12]int) {

	for gap := len(arr) / 2; gap > 0; gap = gap / 2 {
		// back pointer
		for back := gap; back < len(arr); back++ {
			front := back - gap
			// Insert sort
			tmp := arr[back]
			for ; front >= 0 && arr[front] > tmp; front = front - gap {
				arr[front+gap] = arr[front]
			}
			arr[front+gap] = tmp
		}
	}
}
