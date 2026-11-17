# Desafios estilo LeetCode em Go

Cada pasta tem:
- `solucao.go` — enunciado + função vazia para você implementar
- `solucao_test.go` — testes que validam sua resposta

## Como usar

```bash
go test ./01_two_sum/          # testa um desafio
go test -v ./01_two_sum/       # mostra cada caso
go test ./...                  # testa todos
go test -bench=. ./01_two_sum/ # benchmark (quando houver)
```

Todos os testes começam falhando. Implemente a função até ficar verde.

| #  | Desafio                    | Nível  | Conceito             |
|----|----------------------------|--------|----------------------|
| 01 | Two Sum                    | Fácil  | map                  |
| 02 | Parênteses Válidos         | Fácil  | pilha (slice)        |
| 03 | Palíndromo Válido          | Fácil  | dois ponteiros, rune |
| 04 | Merge de Intervalos        | Médio  | sort.Slice           |
| 05 | Maior Substring sem Repetição | Médio | janela deslizante  |
| 06 | Agrupar Anagramas          | Médio  | map de chave composta|
