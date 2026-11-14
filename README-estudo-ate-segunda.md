# Plano de estudo até segunda-feira

O objetivo deste plano não é dominar toda a stack. É conseguir resolver uma
tarefa pequena em Go + Vue, explicar o raciocínio e pedir ajuda de forma
objetiva.

## Prioridade

Distribua o tempo disponível aproximadamente assim:

- **50% — Go:** ler código, corrigir erros e criar endpoints;
- **25% — Vue + TypeScript:** formulário e consumo de API;
- **15% — testes, depuração e explicação em voz alta;
- **10% — SQL, AWS e autenticação em nível conceitual.**

Deixe Flutter, integração real com ERP e deploy completo na AWS para depois.

## Hoje — entender o projeto atual

Reserve aproximadamente duas horas, com uma pausa.

### 1. Faça o projeto funcionar

- [ ] Backend compila.
- [ ] Frontend abre.
- [ ] Lista tarefas.
- [ ] Cria tarefa.
- [ ] Altera o estado da tarefa.
- [ ] Mostra uma mensagem quando ocorre um erro.

Execute o backend na raiz:

```powershell
go run .
```

Em outro terminal, execute o frontend:

```powershell
cd frontend
npm run dev
```

Valide o backend com:

```powershell
go test ./...
go vet ./...
```

Se algo falhar, leia o primeiro erro, localize o trecho, formule uma hipótese
e teste uma alteração por vez.

### 2. Explique o fluxo

Pratique explicar sem consultar:

```text
Usuário envia formulário
        ↓
Vue faz uma requisição
        ↓
Handler recebe e valida
        ↓
Função altera os dados
        ↓
API responde
        ↓
Vue atualiza a tela
```

Ao ler a função `Create`, responda:

- O que são `w` e `r`?
- O que `Decode(&body)` faz?
- Por que existe um `return` depois do erro?
- Onde a validação acontece?
- Onde a tarefa é armazenada?
- O que ocorre quando o título está vazio?

Anote as perguntas que não conseguir responder. Elas serão o foco da revisão.

## Domingo — simulado de 90 minutos

### Enunciado

Adicione prioridade às tarefas. A prioridade pode ser `baixa`, `média` ou
`alta`. A API deve rejeitar valores inválidos. A interface deve permitir
escolher a prioridade e filtrar a listagem.

### Divisão sugerida

- **10 min:** entender requisitos e planejar;
- **30 min:** implementar o backend;
- **25 min:** implementar o frontend;
- **15 min:** testar e corrigir;
- **10 min:** explicar a solução e suas limitações.

### Requisitos

- [ ] Adicionar o campo de prioridade no backend.
- [ ] Validar valores recebidos.
- [ ] Retornar a prioridade no JSON.
- [ ] Criar um `select` no formulário.
- [ ] Exibir a prioridade na lista.
- [ ] Implementar filtro por prioridade.
- [ ] Testar um caso válido e um inválido.

Na primeira tentativa, evite pedir código completo à IA. Faça perguntas
específicas e implemente a solução você mesmo.

## Como investigar e pedir ajuda

Use este padrão:

1. diga o que entendeu;
2. diga o que já verificou;
3. formule uma dúvida específica;
4. aplique a sugestão;
5. repita o caso que falhou.

Frases úteis:

> “Entendi a regra que preciso implementar, mas não lembro a função de Go para
> isso. Posso consultar a documentação?”

> “Já confirmei que a requisição chega ao handler. Minha dúvida agora é nesta
> conversão. Vocês podem me orientar sobre o tipo esperado?”

> “O compilador está apontando esta chamada. Vou conferir se a função existe e
> se o nome corresponde ao declarado.”

## Revisão rápida

### Go

- diferença entre declarar uma variável e chamar uma função;
- retorno de resultado e erro;
- acesso a campos de uma `struct`;
- iteração sobre um `slice`;
- separação em pacotes;
- leitura de JSON e resposta de erro HTTP;
- uso conceitual de `context` e goroutines.

### Vue e TypeScript

- onde fica o estado do formulário;
- como enviar dados para a API;
- como tratar carregamento e falha;
- como tipar uma tarefa;
- como passar dados para um componente;
- quais comportamentos testar com Vitest.

### SQL e integrações

- relacionar cliente e título a receber;
- buscar somente títulos em aberto;
- decidir quando usar uma transação;
- evitar duplicação durante uma importação.

Não tente decorar serviços AWS que ainda não aplicou. Explique o que conhece e
admita com clareza o que ainda precisa investigar.

## Segunda-feira antes da entrevista

- [ ] Rodar o projeto uma vez.
- [ ] Revisar os erros corrigidos e o motivo de cada correção.
- [ ] Reler as anotações de Go.
- [ ] Conferir horário, endereço e deslocamento.
- [ ] Testar o notebook, se for levá-lo.
- [ ] Perguntar quais consultas e ferramentas são permitidas.

## Apresentação possível

> “Minha experiência principal é com TypeScript, NestJS, APIs e banco de
> dados. Estou aprofundando Go e Vue com um projeto prático. Em Go ainda tenho
> pontos para consolidar, mas consigo trabalhar a partir do problema,
> investigar os erros e validar a solução. Quando encontro algo que não
> conheço, procuro entender o comportamento esperado e esclarecer a dúvida
> antes de implementar.”

Adapte essa apresentação para refletir exatamente sua experiência.
