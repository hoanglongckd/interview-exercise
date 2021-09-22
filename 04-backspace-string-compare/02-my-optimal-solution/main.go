package _2_my_optimal_solution

func backspaceCompare(s string, t string) bool {
	numSharpOfS := 0
	numSharpOfT := 0
	p1 := len(s) - 1
	p2 := len(t) - 1

	for p1 >= -1 && p2 >= -1 {
		if p1 > -1 {
			if s[p1] == '#' {
				numSharpOfS++
				p1--
				continue
			} else if numSharpOfS > 0 {
				numSharpOfS--
				p1--
				continue
			}
		}

		if p2 > -1 {
			if t[p2] == '#' {
				numSharpOfT++
				p2--
				continue
			} else if numSharpOfT > 0 {
				numSharpOfT--
				p2--
				continue
			}
		}
		if p1 == -1 && p2 == -1 {
			return true
		}
		if p1 == -1 || p2 == -1 {
			return false
		}
		if s[p1] != t[p2] {
			return false
		}
		p1--
		p2--
	}
	if p1 != p2 {
		return false
	}
	return true
}
