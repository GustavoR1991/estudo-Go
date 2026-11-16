package main

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"
)

func comissoesHandler(w http.ResponseWriter, r *http.Request) {
	vendedorTexto := r.URL.Query().Get("vendedor")
	mes := r.URL.Query().Get("mes")

	if vendedorTexto == "" || mes == "" {
		http.Error(w, "informe vendedor e mes (ex.: ?vendedor=12&mes=2026-10)", http.StatusBadRequest)
		return
	}

	// TODO 1: converta vendedorTexto em int com strconv.Atoi.
	// Se der erro, responda 400 com "vendedor deve ser um número".

	vendedorID, err := strconv.Atoi(vendedorTexto)
	if err != nil {
		http.Error(w, "vendedor deve ser um número", http.StatusBadRequest)
		return
	}

	inicio, err := time.Parse("2006-01", mes)
	if err != nil {
		http.Error(w, "mes inválido, use AAAA-MM", http.StatusBadRequest)
		return
	}
	fim := inicio.AddDate(0, 1, 0)

	// TODO 2: use buscarPercentual. Se não achar, responda 404 com
	// "vendedor não encontrado" (http.StatusNotFound).

	percentual, achou := buscarPercentual(vendedorID, configsMock)
	if !achou {
		http.Error(w, "vendedor nao encontrado", http.StatusNotFound)
	}

	// TODO 3: calcule com calcularComissao(vendasMock, ...) e devolva o JSON
	// com a struct RespostaComissao.

	total := calcularComissao(vendasMock, vendedorID, percentual, inicio, fim)

	resp := RespostaComissao{VendedorID: vendedorID, Mes: mes, Total: total}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func buscarPercentual(vendedorID int, configs []ComissaoConfig) (float64, bool) {
	for _, c := range configs {
		if c.VendedorID == vendedorID {
			return c.Percentual, true
		}
	}
	return 0, false // não achou
}

func calcularComissao(vendas []Venda, vendedorID int, percentual float64, inicio, fim time.Time) float64 {
	var total float64
	for _, v := range vendas {
		if v.VendedorID != vendedorID {
			continue // nao e desse vendedor
		}

		if v.Status != "faturada" {
			continue
		}

		if v.DataVenda.Before(inicio) || !v.DataVenda.Before(fim) {
			continue
		}

		total += v.Valor * percentual / 100

	}

	return math.Round(total*100) / 100
}
