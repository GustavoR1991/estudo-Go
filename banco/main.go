package main

import "fmt"

type Conta struct {
	saldo float64
}

func (c *Conta) depositar(valor float64) {
	if valor <= 0 {
		fmt.Println("Valor inválido")
		return
	}

	c.saldo += valor

	fmt.Printf("Deposito de R$ %.2f realizado com sucesso. Saldo atual: R$ %.2f\n", valor, c.saldo)

}

func (c *Conta) sacar(valor float64) bool {
	if valor <= 0 {
		fmt.Println("Valor inválido")
		return false
	}

	if valor > c.saldo {
		fmt.Println("Saldo insuficiente")
		return false
	}

	c.saldo -= valor

	fmt.Printf("Saque de R$ %.2f realizado com sucesso. Saldo atual: R$ %.2f\n", valor, c.saldo)

	return true
}

func (c *Conta) verSaldo( )float64 {
	return c.saldo
}

func main() {
	conta := Conta{}
	var opcao string

	for {
		fmt.Println()
		fmt.Println("Digite a operação desejada:")
		fmt.Println("depositar, sacar, saldo ou sair")
		fmt.Print("Opção: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "depositar":
			var valor float64

			fmt.Print("Digite o valor do depósito: ")
			fmt.Scanln(&valor)

			conta.depositar(valor)

		case "sacar":
			var valor float64

			fmt.Print("Digite o valor do saque: ")
			fmt.Scanln(&valor)

			conta.sacar(valor)

		case "saldo":
			fmt.Printf("Saldo atual: R$ %.2f\n", conta.verSaldo())

		case "sair":
			fmt.Println("Encerrando...")
			return

		default:
			fmt.Println("Opção inválida")
		}
	}
}