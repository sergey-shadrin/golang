package anagram

import (
	"testing"
)

type AnagramTestCase struct {
	first, second string
	want          bool
}

var cases = []AnagramTestCase{
	{"", "", true},
	{"T", "T", true},
	{"T", "t", true},
	{"покраснение", "пенсионерка", true},
	{"T", "", false},
	{"", "T", false},
	{"G", "T", false},
	{"12345", "1234", false},
	{"12345", "12345", true},
	{"12345", "12245", false},
	{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non sunt proident, in culpa qui officia deserunt mollit anim id est laborum.", true},
	{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non sunt proident, in culpa qui officia deserunt mollix anim id est laborum.", false},
}

func TestIsAnagram(t *testing.T) {
	for _, testCase := range cases {
		got := IsAnagram(testCase.first, testCase.second)
		if got != testCase.want {
			t.Errorf("IsAnagram(%q, %q) == %v, want %v", testCase.first, testCase.second, got, testCase.want)
		}
	}
}

func TestIsAnagramWithSort(t *testing.T) {
	for _, testCase := range cases {
		got := IsAnagramWithSort(testCase.first, testCase.second)
		if got != testCase.want {
			t.Errorf("IsAnagramWithSort(%q, %q) == %v, want %v", testCase.first, testCase.second, got, testCase.want)
		}
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsAnagram("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non sunt proident, in culpa qui officia deserunt mollit anim id est laborum.")
	}
}

func BenchmarkIsAnagramWithSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsAnagramWithSort("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non sunt proident, in culpa qui officia deserunt mollit anim id est laborum.")
	}
}