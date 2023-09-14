package reload

func Atoi(s string) int {
	runeS := []rune(s)
	intArr := []int{}
	var output int = 0
	is := false
	neg := false
	if len(s) > 0 {
		if s[0] == '-' || s[0] == '+' {
			if len(runeS) > 1 {
				if s[1] <= 57 && 48 <= s[1] {
					if s[0] == '-' {
						neg = true
					}

					runeS = runeS[1:]
				}
			} else {
				is = true
			}
		}
	}

	for _, run := range runeS {
		if run <= 57 && 48 <= run && is == false {
			intArr = append(intArr, int(run-'0'))
		} else {
			intArr = []int{}
			output = 0
			break
		}
	}

	for _, num := range intArr {
		output *= 10
		output += num
	}
	if neg {
		output = -output
	}

	return output
}
