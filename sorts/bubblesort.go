package sorts

// 5, 2, 10, 7, 220, 22, 1
// starts with first and comparing it to the next
// if current value is greater then next
// both switches their indexes in the array
func BubbleSort(list []int) []int {
	for i, k := 0, len(list)-1; i < len(list)-1; i++ {
		for j := 0; j < k; j++ {
			cur := list[j]
			next := list[j+1]
			if cur > next {
				list[j] = next
				list[j+1] = cur
			}
		}

		k--
	}

	return list
}
