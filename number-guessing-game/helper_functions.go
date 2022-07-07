package main

func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func parseDigits(n int) []int { //parsing results in reversed slice so reverted it back with the func above
	var ret []int
	for n != 0 {
		ret = append(ret, n%10)
		n /= 10
	}
	reverseInt(ret)
	return ret
}

func containsInt(s []int, str int) bool { //check if a int slice contains given integer
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
