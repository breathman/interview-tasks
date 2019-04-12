package arrays

func Merge(ar1, ar2 []int) []int {
	lar1, lar2 := len(ar1), len(ar2)
	var a int
	for i := lar2 - 1; i >= 0; i-- {
		last := ar1[lar1-1]

		if ar2[i] < last {
			a = ar2[i]
			ar2[i] = last
		} else {
			continue
		}

		for j := lar1 - 1; j >= 0; j-- {
			if a >= ar1[j] {
				copy(ar1[j+2:], ar1[j+1:])
				ar1[j+1] = a
				break
			}
		}
	}

	return append(ar1, ar2...)
}
