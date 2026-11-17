package parenteses

import "fmt"

// Parênteses Válidos (Fácil)
//
// Dada uma string contendo apenas '(', ')', '{', '}', '[' e ']',
// diga se ela é válida:
//   - todo bloco aberto deve ser fechado pelo mesmo tipo
//   - blocos devem ser fechados na ordem correta
//
// Exemplos: "()[]{}" -> true   "(]" -> false   "([)]" -> false   "{[]}" -> true
//
// Dica: em Go uma pilha é só um slice: append para empilhar, s[:len(s)-1] para desempilhar.
func EhValido(s string) bool {
	pares := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var pilha []rune

	for _, c := range s {
		if abre, ehFechamento := pares[c]; ehFechamento {
			if len(pilha) == 0 || pilha[len(pilha)-1] != abre {
				return false
			}
			pilha = pilha[:len(pilha)-1] // desempilha
		} else {
			pilha = append(pilha, c) // c é uma abertura: empilha
		}
	}

	return len(pilha) == 0 // sobrou algo aberto? então é inválido

}
func main() {
	for _, s := range []string{"()[]{}", "(]", "([)]", "{[]}", "", "(("} {
		fmt.Printf("%q -> %v\n", s, EhValido(s))
	}
}
