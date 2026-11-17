package substring

import "testing"

func TestMaiorSubstring(t *testing.T) {
	casos := []struct {
		s    string
		quer int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"dvdf", 3},
		{"abba", 2},
		{"tmmzuxt", 5},
	}
	for _, c := range casos {
		if got := MaiorSubstring(c.s); got != c.quer {
			t.Errorf("MaiorSubstring(%q) = %d; quer %d", c.s, got, c.quer)
		}
	}
}
