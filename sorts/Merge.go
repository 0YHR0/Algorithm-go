package sorts

// 归并排序
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 获取分区位置
	p := len(nums) / 2
	// 通过递归分区
	left := MergeSort(nums[0:p])
	right := MergeSort(nums[p:])
	// 排序后合并
	return merge(left, right)
}

// 排序合并
func merge(left []int, right []int) []int {
	i, j := 0, 0
	m, n := len(left), len(right)
	// 用于存放结果集
	var result []int
	for {
		// 任何一个区间遍历完，则退出
		if i >= m || j >= n {
			break
		}
		// 对所有区间数据进行排序
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 如果左侧区间还没有遍历完，将剩余数据放到结果集
	if i != m {
		for ; i < m; i++ {
			result = append(result, left[i])
		}
	}

	// 如果右侧区间还没有遍历完，将剩余数据放到结果集
	if j != n {
		for ; j < n; j++ {
			result = append(result, right[j])
		}
	}

	// 返回排序后的结果集
	return result
}
