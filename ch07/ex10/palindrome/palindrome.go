package palindrome

import "sort"

func IsPalindrome(s sort.Interface) bool {
	l := s.Len()
	h := l >> 1
	for i := 0; i < h; i++ {
		if s.Less(i, l-i-1) || s.Less(l-i-1, i) {
			return false
		}
	}
	return true
}
