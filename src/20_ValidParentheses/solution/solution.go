package solution

func IsValid(s string) bool {

	Pars := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []rune

	for _, v := range s {

		i, ok := Pars[v]
		if ok {
			stack = append(stack, i)
			continue
		}

		l := len(stack)
		if l == 0 || stack[l-1] != v {
			return false
		}

		stack = stack[:l-1]
	}

	return len(stack) == 0

}
