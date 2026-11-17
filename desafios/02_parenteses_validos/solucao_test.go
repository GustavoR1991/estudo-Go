package parenteses

import "testing"

func TestEhValido(t *testing.T) {
	casos := []struct {
		s    string
		quer bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"{[]}", true},
		{"", true},
		{"(]", false},
		{"([)]", false},
		{"(", false},
		{")", false},
		{"((()))[{}]", true},
		{"((", false},
	}
	for _, c := range casos {
		if got := EhValido(c.s); got != c.quer {
			t.Errorf("EhValido(%q) = %v; quer %v", c.s, got, c.quer)
		}
	}
}
