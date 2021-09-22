package _3_instructor_optimal_solution

func backspaceCompare(s string, t string) bool {
	p1 := len(s) - 1
	p2 := len(t) - 1

	for p1 >= 0 || p2 >= 0 {
		skipS := 0

		for p1 >= 0 {
			if s[p1] == '#' {
				skipS++
				p1--
			} else if skipS > 0 {
				skipS--
				p1--
			} else {
				break
			}
		}
		skipT := 0

		for p2 >= 0 {
			if t[p2] == '#' {
				skipT++
				p2--
			} else if skipT > 0 {
				skipT--
				p2--
			} else {
				break
			}
		}
		if (p1 >= 0) != (p2 >= 0) {
			return false
		}
		if p1 >= 0 && p2 >= 0 && s[p1] != t[p2] {
			return false
		}
		p1--
		p2--
	}
	return true
}
