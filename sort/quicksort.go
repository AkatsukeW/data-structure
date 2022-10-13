package sort

func QuickSort(nums []int, left, end int) {
	if left > end {
		return
	}

	tmp := nums[left]
	k, v := left, end
	for k != v {
		for nums[v] >= tmp && v > k {
			v--
		}
		for nums[k] <= tmp && v > k {
			k++
		}
		if k < v {
			nums[k], nums[v] = nums[v], nums[k]
		}
	}

	nums[left] = nums[k]
	nums[k] = tmp
	QuickSort(nums, left, k-1)
	QuickSort(nums, k+1, end)
	return
}
