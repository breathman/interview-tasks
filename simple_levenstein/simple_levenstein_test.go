package simple_levenstein

import "testing"

var posCases = map[string]string{
	"bb":  "bb",
	"baa": "bba",

	"ab":    "aa",
	"ba":    "aa",
	"abbaa": "aabaa",

	"aaa":     "aa",
	"aa":      "aaa",
	"b":       "",
	"aaab":    "aaa",
	"aab":     "aa",
	"abb":     "bb",
	"aabbb":   "abbb",
	"abbb":    "aabb",
	"aacbb":   "aabb",
	"abcdefg": "abcefg",
}

var negCases = map[string]string{
	"aa":   "aaaa",
	"aaaa": "aa",
	"bb":   "",

	"ab":  "ba",
	"aab": "baa",

	"aaa":  "ab",
	"baaa": "bab",
}

func TestIsValid(t *testing.T) {
	for k, v := range posCases {
		res := IsValid(k, v)
		if !res {
			t.Errorf("check case %q - %q result must be true", k, v)
		}
	}

	for k, v := range negCases {
		res := IsValid(k, v)
		if res {
			t.Errorf("check case %q - %q result must be false", k, v)
		}
	}
}

func TestIsValidRec(t *testing.T) {
	for k, v := range posCases {
		res := IsValidRec(k, v)
		if !res {
			t.Errorf("check case %q - %q result must be true", k, v)
		}
	}

	for k, v := range negCases {
		res := IsValidRec(k, v)
		if res {
			t.Errorf("check case %q - %q result must be false", k, v)
		}
	}
}
