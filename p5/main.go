package main

func isPalindromic(s []rune) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// isPalindromicView checks if a slice view [start, end) is a palindromic string
func isPalindromicView(s []rune, start int, end int) bool {
	length := end + start
	for i := start; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func getAllSubstrings(runicString []rune) [][]rune {
	solutions := make([][]rune, 0)
	for i := range runicString {
		for j := i + 1; j <= len(runicString); j++ {
			solutions = append(solutions, runicString[i:j])
		}
	}

	// for _, v := range solutions {
	// 	fmt.Println(string(v))
	// }

	return solutions
}

func longestPalindrome1(s string) string {
	var sol []rune
	for _, substr := range getAllSubstrings([]rune(s)) {
		if len(substr) > len(sol) && isPalindromic(substr) {
			sol = substr
		}
	}
	return string(sol)
}

func longestPalindrome2(s string) string {
	runicString := []rune(s)
	var sol []rune
	for i := range runicString {
		for j := i + 1; j <= len(runicString); j++ {
			substr := runicString[i:j]
			if isPalindromic(substr) && len(sol) < len(substr) {
				sol = substr
			}
		}
	}
	return string(sol)
}

func longestPalindrome3(s string) string {
	runicString := []rune(s)
	var sol []rune

	for i := range runicString {
		for j := 0; j <= i; j++ {
			// center in `i`
			left := i - j
			right := i + j + 1
			if right > len(runicString) {
				break
			}

			if right-left > len(sol) && isPalindromicView(runicString, left, right) {
				sol = runicString[left:right]
				//fmt.Println("sol: ", string(sol))
			}

			// centered in i-1, i
			if left > 0 {
				left = left - 1
				if right-left > len(sol) && isPalindromicView(runicString, left, right) {
					sol = runicString[left:right]
					//fmt.Println("sol: ", string(sol))
				}
			}
		}
	}
	return string(sol)
}

func longestPalindrome(s string) string {
	return longestPalindrome3(s)
}
