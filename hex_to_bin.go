package reload

// import "fmt"

func AtoiBase(s string, base string) int {
	baseRunes := []rune(base)
	sRunes := reverse([]rune(s))
	baseLen := len(baseRunes)
	sLen := len(sRunes)
	outnum := 0

	for i := 0; i < baseLen; i++ {
		for j := 0; j < baseLen; j++ {
			if baseRunes[i] == baseRunes[j] && i != j {
				return 0
			}
		}
	}

	if base[0] == '-' || base[0] == '+' || len(baseRunes) < 2 {
		return 0
	}

	for i := sLen - 1; i >= 0; i-- {
		outnum += findIn(baseRunes, sRunes[i]) * powerOf(baseLen, i)
	}

	return outnum
}

func reverse(s []rune) []rune {
	// from stackOverflow
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func findIn(arr []rune, bnum rune) int {
	for i, run := range arr {
		switch {
		case run == bnum:
			return i
		}
	}
	return 0
}

func powerOf(num, pow int) int {
	if pow == 0 {
		return 1
	}
	store := num
	for i := 1; i < pow; i++ {
		store *= num
	}
	return store
}
