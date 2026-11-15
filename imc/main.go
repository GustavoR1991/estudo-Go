package main

// Perguntar o nome da pessoa.

// Perguntar o peso em quilogramas.

// Perguntar a altura em metros.

// Validar os dados:

// Peso deve ser maior que zero.

// Altura deve ser maior que zero.

// Calcular o IMC.

// Mostrar o IMC com duas casas decimais.

// Classificar o resultado:

import (
	"fmt"
)

func main () {
	var peso, altura float64

	var nome string

	fmt.Println("Digite o nome do paciente:")
	fmt.Scanln(&nome)

	fmt.Println("Digite a altura do paciente (em m):")
	fmt.Scanln(&altura,)

	if altura <= 0 { 
		fmt.Println("Altura inválida")
		return
	}

	fmt.Println("Digite o peso do paciente (em kg):")
	fmt.Scanln(&peso)

  if peso <= 0 {
		fmt.Println(" peso invalido")
		return
	}

	imc := peso / (altura * altura)
	fmt.Println()
	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Altura: %.2f m\n", altura)
	fmt.Printf("Peso: %.2f kg\n", peso)
	fmt.Printf("IMC: %.2f\n", imc)


	if imc < 18.5 {
		fmt.Println("abaixo do peso")
	} else if imc >= 18.5 && imc <25 {
		fmt.Println("peso normal")
	} else if imc >= 25 && imc < 30 {
		fmt.Println("Sobrepeso")
	} else if imc >= 30 {
		fmt.Println("Obesidade")
	 }


}