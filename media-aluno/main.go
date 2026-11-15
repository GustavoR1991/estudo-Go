package main

import "fmt"

func main() {
	var n1 , n2, n3 float64

	var media float64

	var nome string

	fmt.Println("Digite o nome do aluno:")
	fmt.Scanln(&nome)
	fmt.Println("Digite a primeira nota:")
	fmt.Scanln(&n1)

	if n1 < 0 || n1 > 10 {
		fmt.Println("O aluno", nome, "tem uma nota inválida:", n1	)
		return
	}


	fmt.Println("Digite a segunda nota:")
	fmt.Scanln(&n2)

	if n2 < 0 || n2 > 10 {
		fmt.Println("O aluno", nome, "tem uma nota inválida:", n2)
		return
	}

	fmt.Println("Digite a terceira nota:")
	fmt.Scanln(&n3)

	if n3 < 0 || n3 > 10 {
		fmt.Println("O aluno", nome, "tem uma nota inválida:", n3)
		return
	}

	media = float64(n1 + n2 + n3) / 3

	fmt.Printf("Aluno: %s\n", nome)
	fmt.Printf("Notas: %.2f, %.2f, %.2f\n", n1, n2, n3)
	fmt.Printf("Média final: %.2f\n", media)
	
if media >= 7 {
	fmt.Println("Aluno aprovado")
} else if media >= 5 {
	fmt.Println("Aluno em recuperacao")
} else {
	fmt.Println("Aluno reprovado")
}


}