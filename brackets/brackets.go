package brackets

func Valid(s string) bool {
	var i int
	for _, v := range s {
		if string(v) == "(" {
			i++
		} else {
			i--
		}
		if i < 0 {
			return false
		}
	}
	if i == 0 {
		return true
	}
	return false
}
