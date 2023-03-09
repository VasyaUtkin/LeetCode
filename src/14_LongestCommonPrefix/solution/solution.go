package solution

func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ {
		for j, word := range strs {
			if j == 0 {
				continue
			}
			if i == len(word) || word[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
