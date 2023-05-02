package binarySearch

/*
[lo, hi)
low is always inclusive
hi is always exclusive

 lets say we are searching 3

needle := 3

 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
 m = math.Floor(0 + (10 - 0) / 2) = 5
 m = math.Floor(lo + (hi - lo) / 2)
 if m == needle {
    return m
} else if needle > m {
    lo += m + 1
} else {
    hi -= m
}
*/
// one important note is that lo, hi and m all are indexes not values
func bs_list(list []int, needle int) bool {
	lo := 0
	hi := len(list)

	for {

		// no need for math floor since
		// go automatically rounds to floor
		// when dividing odd integer by 2
		m := lo + (hi-lo)/2
        v := list[m]

		if needle == v {
			return true
		} else if needle > v {
			lo = m + 1
		} else {
			hi = m
		}

        if lo >= hi {
            break
        }
	}

	return false
}
