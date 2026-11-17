package palindromo

import "testing"

func TestEhPalindromo(t *testing.T) {
	casos := []struct {
		s    string
		quer bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
		{"ab_a", true},
		{"Ana", true},
		{"Socorram-me, subi no ônibus em Marrocos", false}, // 'ô' != 'o' sem normalização
		{"ôvô", true},
	}
	for _, c := range casos {
		if got := EhPalindromo(c.s); got != c.quer {
			t.Errorf("EhPalindromo(%q) = %v; quer %v", c.s, got, c.quer)
		}
	}
}
