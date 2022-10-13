package sort

func ShellSort(num []int) {
	length := len(num)

	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			tmp := num[i]

			var j int
			for j = i; j >= gap && num[j-gap] > tmp; j -= gap {
				num[j] = num[j-gap]
			}
			num[j] = tmp
		}
	}
}
