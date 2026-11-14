// import sem uso é erro de compilação
// Removeu o import os, que não era usado. Em Go, import sem uso é erro de compilação.
// Adicionou o return resultado. A função declarava que devolvia int mas não devolvia nada.
// Trocou "3" (texto) por 3 (número), porque soma espera int.

package main

// EXERCICIO 1

// import (
// 	"fmt"
// )

// func soma(a int, b int) int {
//     resultado := a + b
// 		return resultado
// }

// func main() {
//     total := soma(2, 3)
//     fmt.Println("Total:", total)
// }

// make faz o que ?

// var totais map[string]int64 só declara a variável. O valor dela é nil, ou seja, um map que ainda não existe.
// Ler de um map nil funciona e devolve o valor zero. Escrever nele causa panic: assignment to entry in nil map, que era a mensagem que aparecia ao rodar.
// make(map[string]int64) cria o map de verdade, vazio, e aí pode escrever.

// totais := map[string]int64{} = var totais = make(map[string]int64)

//o map não estava “sendo salvo”, ele nunca tinha sido criado. var totais map[string]int64 deixa a variável como nil, e o Go não deixa escrever em um map que não existe. Dizer “o map estava nil porque faltou o make” é a explicação mais precisa para um entrevistador.
//Você fez o caminho que a entrevista avalia:

// Rodou e leu o erro.
// Achou a linha.
// Entendeu a causa.
// Corrigiu.

// EXERCICIO 2

// type Titulo struct {
// 	Cliente  string
// 	Centavos int64
// }

// func totalPorCliente(titulos []Titulo) map[string]int64 {
// 	var totais = make(map[string]int64)
// 	for _, t := range titulos {
// 		totais[t.Cliente] += t.Centavos
// 	}
// 	return totais
// }

// func main() {
// 	titulos := []Titulo{{"Ana", 1000}, {"Beto", 500}, {"Ana", 250}}
// 	fmt.Println(totalPorCliente(titulos))
// }

// EXERCICIO 3

// Todo programa executável em Go começa dizendo a que pacote pertence.
// "main" é o pacote especial que o Go roda quando você faz "go run ."

// import: traz pacotes (bibliotecas) que o código usa.
// Se importar e não usar, o Go dá erro de compilação.
import (
	"fmt"     // fmt: imprimir na tela e montar mensagens (Println, Errorf)
	"strconv" // strconv: converter texto em número (ParseInt)
	"strings" // strings: mexer em texto (Split)
)

// func: define uma função. Esta se chama parseValor.
// (s string)        -> recebe UM parâmetro chamado s, do tipo texto
// (int64, error)    -> devolve DOIS valores: um número e um possível erro
// Em Go é comum devolver "resultado, erro". Se deu tudo certo, o erro é nil.
func parseValor(s string) (int64, error) {

	// strings.Split corta o texto em pedaços usando o separador ".".
	// "10.10" vira ["10", "10"]   (um slice com 2 pedaços)
	// "abc"   vira ["abc"]        (um slice com 1 pedaço)
	// ":=" declara a variável "partes" e já dá um valor a ela.
	partes := strings.Split(s, ".")

	// len(partes) é a quantidade de pedaços.
	// Mais de 2 pedaços ("1.2.3") significa que o valor não faz sentido.
	if len(partes) > 2 {
		// return devolve os dois valores da função: 0 e um erro.
		// fmt.Errorf cria um erro com mensagem formatada.
		// %q coloca o texto entre aspas ("1.2.3"), bom para ver espaços e lixo.
		return 0, fmt.Errorf("valor inválido %q: mais de um ponto", s)
	}

	// strconv.ParseInt converte texto em número inteiro.
	// Argumentos: (texto, base, bits) -> base 10 (decimal), 64 bits.
	// Devolve DOIS valores: o número (reais) e um erro (err).
	reais, err := strconv.ParseInt(partes[0], 10, 64)

	// err != nil significa "deu erro". Se o texto era "abc", cai aqui.
	// Conferir isso é obrigatório: se você esquece, o erro passa batido.
	if err != nil {
		// %w "embrulha" o erro original dentro do novo, mantendo a causa.
		return 0, fmt.Errorf("valor inválido %q: %w", s, err)
	}

	// var declara uma variável sem valor inicial.
	// Todo int64 começa valendo 0, então centavos começa em 0.
	// Para valores sem parte decimal ("5"), ele continua 0.
	var centavos int64

	// Só entra aqui se houver parte depois do ponto ("10.10" tem; "5" não).
	// Isso evita o erro "index out of range" ao acessar partes[1].
	if len(partes) == 2 {
		c := partes[1] // c = o texto dos centavos, por exemplo "1" ou "10"

		// "10.1" significa 10 reais e 10 centavos, não 1 centavo.
		// Então, se veio só 1 dígito, acrescentamos um "0" no fim.
		if len(c) == 1 {
			c += "0" // += junta o texto: "1" vira "10"
		}

		// Depois do ajuste, precisa ter exatamente 2 dígitos.
		// "7.999" tem 3 dígitos nos centavos, então é inválido.
		if len(c) != 2 {
			return 0, fmt.Errorf("valor inválido %q: centavos devem ter 1 ou 2 dígitos", s)
		}

		// Aqui usamos "=" e não ":=", porque centavos e err JÁ existem.
		// ":=" criaria variáveis novas só dentro deste bloco.
		centavos, err = strconv.ParseInt(c, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("valor inválido %q: %w", s, err)
		}
	}

	// Tudo certo: converte para centavos totais.
	// 10 reais e 10 centavos = 10*100 + 10 = 1010.
	// nil no lugar do erro significa "não houve erro".
	return reais*100 + centavos, nil
}

// main é a função onde o programa começa a executar.
// func main() {
// 	// []string{...} cria um slice (lista) de textos para testar.
// 	entradas := []string{"10.1", "5", "1.2.3", "7.999", "abc"}

// 	// Acumulador do total, em centavos. Começa em 0.
// 	var total int64

// 	// for ... range percorre a lista.
// 	// "_" ignora o índice (posição), que não usamos.
// 	// "e" recebe cada texto da lista, um por vez.
// 	for _, e := range entradas {
// 		// Chama a função e guarda os DOIS retornos.
// 		valor, err := parseValor(e)

// 		if err != nil {
// 			fmt.Println("aviso:", err) // mostra o problema na tela
// 			continue                   // pula para a próxima entrada, sem somar esta
// 		}

// 		total += valor // += soma e guarda: total = total + valor
// 	}

// 	// Imprime o resultado final, em centavos.
// 	fmt.Println("total:", total)
// }
