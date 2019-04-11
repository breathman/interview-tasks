package brackets

import "testing"

func TestValid(t *testing.T) {
	posCases := []string{
		"()()()()",
		"()",
		"(())(())()",
		"((()))",
		"",
	}

	negCases := []string{
		"(",
		"()(",
		")",
		")(",
		"(()))",
		"((()",
	}

	for _, c := range posCases {
		res := Valid(c)
		if !res {
			t.Errorf("check case %q result must be true", c)
		}
	}

	for _, c := range negCases {
		res := Valid(c)
		if res {
			t.Errorf("check case %q result must be false", c)
		}
	}

}
