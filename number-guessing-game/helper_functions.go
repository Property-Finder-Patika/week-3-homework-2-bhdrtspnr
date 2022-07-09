package main

//how is this not a default function?
func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//how is this not a default function?
func parseDigits(n int) []int { //parsing results in reversed slice so reverted it back with the func above
	var ret []int
	for n != 0 { //parsing integer requires you to add the remainder of 10 then divide number by 10 and %10 again
		ret = append(ret, n%10)
		n /= 10
	}
	reverseInt(ret)
	return ret
}

//how is this not a default function?
func containsInt(s []int, str int) bool { //check if a int slice contains given integer
	for _, v := range s { //_ for index v for item
		if v == str {
			return true
		}
	}
	return false
}
