package answer2

func IsBalanced(s string) string {
	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}
