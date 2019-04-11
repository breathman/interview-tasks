package simple_levenstein

func IsValid(s1, s2 string) bool {

	if len(s1)-len(s2) >= 2 || len(s1)-len(s2) <= -2 {
		return false
	}

	flag := false
	if len(s1) == len(s2) {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if flag {
					return false
				}
				flag = true
			}
		}
	} else {
		var min, max string
		if len(s1) < len(s2) {
			min, max = s1, s2
		} else {
			min, max = s2, s1
		}
		if len(min) == 0 {
			return true
		}

		flag := false
		for i := 0; i < len(min); {
			if flag {
				if min[i] != max[i+1] {
					return false
				}
				i++
			} else {
				if min[i] != max[i] {
					flag = true
				} else {
					i++
				}
			}
		}
	}
	return true
}

func IsValidRec(s1, s2 string) bool {
	if len(s1)-len(s2) >= 2 || len(s1)-len(s2) <= -2 {
		return false
	}

	return validRec(s1, s2, false)
}

func validRec(s1, s2 string, flag bool) bool {
	if len(s1) == len(s2) {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if flag {
					return false
				}
				flag = true
			}
		}
	} else {
		var min, max string
		if len(s1) < len(s2) {
			min, max = s1, s2
		} else {
			min, max = s2, s1
		}
		if len(min) == 0 {
			return true
		}

		flag := false
		for i := 0; i < len(min); i++ {
			if min[i] != max[i] {
				if flag {
					return false
				}
				return validRec(min[i:], max[i+1:], true)
			}
		}
	}
	return true
}
