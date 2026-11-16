package main

import (
	"log"
	"net/http"
	"time"
)

type Venda struct {
	ID         int64
	VendedorID int
	Cliente    string
	Valor      float64
	DataVenda  time.Time
	Status     string
}

type ComissaoConfig struct {
	VendedorID int
	Percentual float64 // ex.: 5.00 significa 5%
}

type RespostaComissao struct {
	VendedorID int     `json:"vendedor_id"`
	Mes        string  `json:"mes"`
	Total      float64 `json:"comissao_total"`
}

func dataVenda(ano int, mes time.Month, dia int) time.Time {
	return time.Date(ano, mes, dia, 0, 0, 0, 0, time.UTC)
}

var vendasMock = []Venda{
	{101, 12, "Alfa Ltda", 1000.00, dataVenda(2026, 10, 3), "faturada"},
	{102, 12, "Beta SA", 2500.00, dataVenda(2026, 10, 10), "cancelada"},
	{103, 12, "Gama ME", 800.00, dataVenda(2026, 10, 15), "faturada"},
	{104, 12, "Delta Ltda", 1500.00, dataVenda(2026, 10, 31), "pendente"},
	{105, 12, "Alfa Ltda", 600.00, dataVenda(2026, 11, 1), "faturada"},
	{106, 15, "Epsilon SA", 2000.00, dataVenda(2026, 10, 5), "faturada"},
}

var configsMock = []ComissaoConfig{
	{12, 5.00},
	{15, 4.00}}

func main() {
	http.HandleFunc("/comissoes", comissoesHandler)

	log.Println("servidor em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
