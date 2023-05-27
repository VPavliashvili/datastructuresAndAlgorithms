package sorts

func QuickSort(arr []int) []int {
	left := 0
	right := len(arr) - 1

	qs(&arr, left, right)

	return arr
}

func qs(arr *[]int, l int, r int) {
	if l >= r {
		return
	}

	p := partition(arr, l, r)

	qs(arr, l, p-1)
	qs(arr, p+1, r)
}

func partition(arr *[]int, l int, r int) int {
	pivot := (*arr)[r]
	i := l - 1
	for j := l; j < r; j++ {
		if (*arr)[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}

	swap(arr, i+1, r)
	return i + 1
}

func swap(arr *[]int, i int, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}
