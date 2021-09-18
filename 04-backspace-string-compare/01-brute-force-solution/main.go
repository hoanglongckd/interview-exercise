package _1_brute_force_solution

func backspaceCompare(s string, t string) bool {
	finalStrOfS := getFinalString(s)
	finalStrOfT := getFinalString(t)

	if len(finalStrOfS) == len(finalStrOfT) {
		for idx, char := range finalStrOfS {
			if char != rune(finalStrOfT[idx]) {
				return false
			}
		}
		return true
	}
	return false
}

func getFinalString(original string) string {
	var finalStrOfS string
	for _, char := range original {
		lenOfFinalStrOfS := len(finalStrOfS)

		if char == '#' && lenOfFinalStrOfS > 0 {
			finalStrOfS = finalStrOfS[0 : lenOfFinalStrOfS-1]
		} else if char != '#' {
			finalStrOfS += string(char)
		}
	}
	return finalStrOfS
}
