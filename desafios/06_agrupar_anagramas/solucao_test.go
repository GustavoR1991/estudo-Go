package anagramas

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

// normaliza ordena as palavras de cada grupo e os grupos entre si,
// para comparar sem depender da ordem.
func normaliza(g [][]string) [][]string {
	out := make([][]string, 0, len(g))
	for _, grupo := range g {
		c := slices.Clone(grupo)
		slices.Sort(c)
		out = append(out, c)
	}
	slices.SortFunc(out, func(a, b []string) int {
		return strings.Compare(strings.Join(a, ","), strings.Join(b, ","))
	})
	return out
}

func TestAgrupar(t *testing.T) {
	casos := []struct {
		in   []string
		quer [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
		{[]string{"ab", "ba", "abc", "cab", "bca", "x"}, [][]string{{"ab", "ba"}, {"abc", "bca", "cab"}, {"x"}}},
	}
	for _, c := range casos {
		got := normaliza(Agrupar(c.in))
		if !reflect.DeepEqual(got, normaliza(c.quer)) {
			t.Errorf("Agrupar(%v) = %v; quer %v", c.in, got, c.quer)
		}
	}
}
