package answer3

import "fmt"

func HighestPalindrome(s string, k int) string {
	n := len(s)
	runes := []rune(s)
	left := make([]rune, n)
	copy(left, runes)
	right := make([]rune, n)
	copy(right, runes)

	var createPalindrome func(int, int, int) (string, int)
	createPalindrome = func(l, r, k int) (string, int) {
		if l >= r {
			return string(left), k
		}
		if left[l] != right[r] {
			if k <= 0 {
				return "-1", k
			}
			if left[l] > right[r] {
				right[r] = left[l]
			} else {
				left[l] = right[r]
			}
			k--
		}
		return createPalindrome(l+1, r-1, k)
	}

	// Ensure the string is a palindrome
	palindrome, k := createPalindrome(0, n-1, k)
	if palindrome == "-1" {
		return "-1"
	}
	runes = []rune(palindrome)
	fmt.Println(s, runes)
	var maximizePalindrome func(int, int, int) string
	maximizePalindrome = func(l, r, k int) string {
		if l > r {
			return string(runes)
		}
		if runes[l] < '9' {
			if l == r {
				if k > 0 {
					runes[l] = '9'
				}
			} else {
				if k >= 2 && runes[l] != '9' && runes[r] != '9' {
					runes[l], runes[r] = '9', '9'
					k -= 2
				} else if k >= 1 && (runes[l] != '9' || runes[r] != '9') && (runes[l] == runes[r]) {
					runes[l], runes[r] = '9', '9'
					k--
				}
			}
		}
		return maximizePalindrome(l+1, r-1, k)
	}

	return maximizePalindrome(0, n-1, k)
}
