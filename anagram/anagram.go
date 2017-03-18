package anagram

import (
	"sort"
	"unicode"
)

func isAnagram(first, second string) bool {
	if len(first) != len(second) {
		return false
	}
	checkedIndexFlags := make([]bool, len(second))
	i := 0
	for _, firstStringChar := range first {
		j := 0
		didFindMatch := false
		for _, secondStringChar := range second {
			if !checkedIndexFlags[j] {
				if unicode.ToLower(secondStringChar) == unicode.ToLower(firstStringChar) {
					didFindMatch = true
					checkedIndexFlags[j] = true
					break
				}
			}
			j++
		}
		if !didFindMatch {
			return false
		}
		i++
	}
	return true
}

func strToLowerAndToInt(str string) []int {
	result := make([]int, len(str))
	i := 0
	for _, ch := range str {
		result[i] = int(unicode.ToLower(ch))
		i++
	}
	return result
}

func isAnagramWithSort(first, second string) bool {
	if len(first) != len(second) {
		return false
	}
	firstStringAsInt := strToLowerAndToInt(first)
	sort.Ints(firstStringAsInt)
	secondStringAsInt := strToLowerAndToInt(second)
	sort.Ints(secondStringAsInt)
	for i := 0; i < len(firstStringAsInt); i++ {
		if firstStringAsInt[i] != secondStringAsInt[i] {
			return false
		}
	}
	return true
}
