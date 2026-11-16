package main

import "testing"

func TestCalcularComissao(t *testing.T) {
	inicio := dataVenda(2026, 10, 1)
	fim := dataVenda(2026, 11, 1)

	casos := []struct {
		nome       string
		vendedorID int
		percentual float64
		esperado   float64
	}{
		{"vendedor 12 em outubro", 12, 5.00, 90.00},
		{"vendedor 15 em outubro", 15, 4.00, 0 /* TODO: quanto deve dar? */},
		{"vendedor sem vendas", 99, 5.00, 0},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			obtido := calcularComissao(vendasMock, c.vendedorID, c.percentual, inicio, fim)
			if obtido != c.esperado {
				t.Errorf("esperado %.2f, obtido %.2f", c.esperado, obtido)
			}
		})
	}
}
