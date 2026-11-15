package main

import "fmt"

func main() {
	var nome string
	var preco float64
	var quantidade int
	var estoque int = 10

	fmt.Print("Digite o nome do produto: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite o preço do produto: ")
	fmt.Scanln(&preco)

	fmt.Print("Digite a quantidade do produto: ")
	fmt.Scanln(&quantidade)

	if quantidade <= 0 {
		fmt.Println("Quantidade inválida")
		return
	}

	if quantidade > estoque {
		fmt.Println("Estoque insuficiente")
		return
	}

	subtotal := preco * float64(quantidade)

	var desconto float64
	var percentualDesconto float64

	if quantidade >= 10 {
		percentualDesconto = 0.10
	} else if quantidade >= 5 {
		percentualDesconto = 0.05
	}

	desconto = subtotal * percentualDesconto
	total := subtotal - desconto

	fmt.Println("Produto:", nome)
	fmt.Printf("Preço unitário: R$ %.2f\n", preco)
	fmt.Println("Quantidade:", quantidade)
	fmt.Printf("Subtotal: R$ %.2f\n", subtotal)
	fmt.Printf("Desconto aplicado: %.0f%%\n", percentualDesconto*100)
	fmt.Printf("Total final: R$ %.2f\n", total)
}