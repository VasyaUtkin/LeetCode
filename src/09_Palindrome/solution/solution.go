package solution

import "strconv"

func IsPalindrome(x int) bool {

	str := strconv.Itoa(x)

	arrBytes := []byte(str)
	if len(arrBytes) < 2 {
		return true
	}

	for i, j := 0, len(arrBytes)-1; i < j; i, j = i+1, j-1 {
		if arrBytes[i] != arrBytes[j] {
			return false
		}
	}

	return true
}
