package huiwen

import (
	"unicode"
)

// IsPalindrome 检测字符串是否为回文
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// 跳过非字母数字字符
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		// 比较字符
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		left++
		right--
	}
	return true
}

// 辅助函数：检查是否为字母数字字符
func isAlphanumeric(char byte) bool {
	return unicode.IsLetter(rune(char)) || unicode.IsDigit(rune(char))
}
