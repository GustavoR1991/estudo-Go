# Perguntas e respostas — preparação Dev Jr

## Fundamentos de programação

### O que é polimorfismo?

Polimorfismo significa que diferentes tipos podem ser tratados por uma mesma abstração, mas cada tipo pode executar o comportamento de uma forma diferente.

Em Go, isso normalmente é feito com interfaces:

```go
type Animal interface {
	Falar() string
}

type Cachorro struct{}

func (Cachorro) Falar() string {
	return "Au"
}

type Gato struct{}

func (Gato) Falar() string {
	return "Miau"
}

func emitirSom(animal Animal) string {
	return animal.Falar()
}
```

`emitirSom` recebe qualquer tipo que implemente `Falar`.

### O que é encapsulamento?

É proteger os detalhes internos de uma estrutura e permitir acesso somente por operações controladas.

Em Go, nomes iniciados com letra maiúscula são exportados:

```go
type User struct {
	Name string
}
```

Nomes iniciados com letra minúscula não são exportados para outros pacotes:

```go
type user struct {
	name string
}
```

### O que é abstração?

É mostrar somente o que é necessário para usar algo e esconder sua implementação interna.

Uma interface é um exemplo:

```go
type Repository interface {
	FindByID(id int) (Project, error)
}
```

Quem usa a interface não precisa saber se os dados vêm de memória, PostgreSQL ou uma API externa.

### O que é herança?

Herança é quando uma classe recebe características de outra classe. Go não possui herança tradicional como Java ou C#.

Em Go, composição é preferida:

```go
type Address struct {
	City string
}

type User struct {
	Name string
	Address
}
```

`User` reutiliza `Address` por composição.

### O que é uma interface?

É um contrato formado por métodos. Um tipo implementa a interface implicitamente quando possui esses métodos.

```go
type Logger interface {
	Log(message string)
}
```

Não é necessário declarar explicitamente que um tipo implementa `Logger`.

### O que é uma struct?

É um tipo que agrupa campos relacionados:

```go
type Project struct {
	ID     int
	Name   string
	Status string
}
```

### Qual a diferença entre variável e constante?

Variável pode mudar:

```go
var status = "active"
status = "paused"
```

Constante não pode ser alterada:

```go
const maxProjects = 100
```

### O que é uma função?

É um bloco reutilizável de código que pode receber parâmetros e retornar valores:

```go
func sum(a int, b int) int {
	return a + b
}
```

### O que é tratamento de erro em Go?

Go retorna erros explicitamente:

```go
result, err := doSomething()
if err != nil {
	return err
}
```

O erro não deve ser ignorado sem uma razão clara.

### O que significa `defer`?

Agenda uma chamada para ser executada quando a função terminar:

```go
file, err := os.Open("data.txt")
if err != nil {
	return err
}
defer file.Close()
```

É muito usado para fechar arquivos, liberar conexões e desbloquear mutexes.

### O que é um ponteiro?

Ponteiro guarda o endereço de uma variável:

```go
name := "Gustavo"
pointer := &name
```

`&name` obtém o endereço. `*pointer` acessa o valor naquele endereço.

### Por que usar `&project` no decoder JSON?

```go
json.NewDecoder(r.Body).Decode(&project)
```

O decoder precisa receber o endereço da variável para preencher o objeto original. Sem `&`, ele receberia uma cópia.

### O que significa `*http.Request`?

É um ponteiro para a requisição HTTP. A requisição contém método, URL, headers, body e parâmetros.

```go
func Create(w http.ResponseWriter, r *http.Request)
```

`w` escreve a resposta e `r` representa a requisição recebida.

### Qual a diferença entre slice e array?

Array possui tamanho fixo:

```go
var numbers [3]int
```

Slice é uma lista dinâmica:

```go
numbers := []int{1, 2, 3}
numbers = append(numbers, 4)
```

### O que é um map?

É uma estrutura de chave e valor:

```go
users := map[string]int{
	"Gustavo": 1,
}
```

Um map `nil` pode ser lido, mas não pode receber escrita até ser inicializado com `make`.

### O que é uma goroutine?

É uma função executada concorrentemente:

```go
go process()
```

O acesso a dados compartilhados precisa ser protegido quando houver risco de condição de corrida.

### Para que serve um mutex?

Mutex controla o acesso simultâneo a dados compartilhados:

```go
mu.Lock()
defer mu.Unlock()
```

Assim, duas goroutines não alteram o mesmo slice ao mesmo tempo.

---

## HTTP e APIs REST

### O que é uma API?

É um contrato que permite que sistemas conversem. Uma API HTTP recebe requisições e retorna respostas.

### O que é REST?

É um estilo de arquitetura que organiza recursos por URLs e usa os métodos HTTP de forma previsível.

### Qual a diferença entre GET, POST, PUT, PATCH e DELETE?

- `GET`: consulta dados.
- `POST`: cria um recurso.
- `PUT`: substitui ou atualiza um recurso completo.
- `PATCH`: altera parte de um recurso.
- `DELETE`: remove um recurso.

### O que significa o status 200?

A requisição foi processada com sucesso.

### O que significa o status 201?

Um novo recurso foi criado com sucesso.

### O que significa o status 204?

A operação teve sucesso, mas não possui corpo de resposta.

### O que significa o status 400?

O cliente enviou dados inválidos, como JSON malformado ou campo obrigatório ausente.

### O que significa o status 401?

O cliente não está autenticado ou não enviou credenciais válidas.

### O que significa o status 403?

O cliente foi identificado, mas não possui permissão.

### O que significa o status 404?

O recurso ou a rota não foi encontrado.

### O que significa o status 409?

Existe conflito com o estado atual, como tentativa de criar um registro duplicado.

### O que significa o status 500?

O servidor encontrou um erro inesperado.

### O que significa o status 503?

O serviço está temporariamente indisponível, por exemplo, porque uma dependência caiu.

### O que é JSON?

É um formato textual para troca de dados:

```json
{
  "name": "Projeto Go",
  "active": true
}
```

### O que é `Content-Type`?

É o header que informa o formato do corpo:

```text
Content-Type: application/json
```

### O que são parâmetros de rota?

São valores dentro da URL:

```text
GET /api/projects/10
```

O `10` pode ser lido com:

```go
r.PathValue("id")
```

### O que são query parameters?

São parâmetros depois de `?`:

```text
GET /api/projects?status=active
```

Em Go:

```go
status := r.URL.Query().Get("status")
```

### Como investigar um 404?

1. Conferir método HTTP.
2. Conferir URL.
3. Conferir prefixo `/api`.
4. Conferir se a rota foi registrada.
5. Reiniciar o servidor.
6. Confirmar se a requisição chegou ao processo correto.

### Como investigar um 500?

1. Ler o log do backend.
2. Reproduzir com o mesmo body.
3. Identificar a linha que falhou.
4. Validar entradas antes de usar.
5. Adicionar tratamento de erro.
6. Criar um teste para o cenário.

---

## Go HTTP

### Como registrar uma rota?

```go
mux.HandleFunc("GET /api/projects", projects.List)
mux.HandleFunc("POST /api/projects", projects.Create)
```

### Como ler JSON do body?

```go
var project Project
err := json.NewDecoder(r.Body).Decode(&project)
```

### Como responder JSON?

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(project)
```

### Por que criar um helper `writeJSON`?

Para evitar repetir headers, status e encode em todos os handlers.

### Onde deve ficar a regra de persistência?

O handler deve cuidar do HTTP. O store ou repository deve cuidar de buscar, criar, atualizar e remover dados.

---

## JavaScript e Vue

### O que é `ref`?

`ref` cria um valor reativo:

```js
const count = ref(0);
count.value++;
```

No template, o Vue acessa o valor sem `.value`.

### O que é `computed`?

É um valor calculado automaticamente a partir de dados reativos:

```js
const completed = computed(() =>
  tasks.value.filter((task) => task.done),
);
```

### Para que serve `onMounted`?

Executa código quando o componente é montado:

```js
onMounted(loadTasks);
```

### O que significa `v-model`?

Cria ligação entre o input e uma variável:

```vue
<input v-model="title" />
```

### Para que serve `v-for`?

Renderiza uma lista:

```vue
<li v-for="task in tasks" :key="task.id">
  {{ task.title }}
</li>
```

### Por que usar `:key`?

Permite que o Vue identifique cada item corretamente ao atualizar a lista.

### Como tratar um `fetch`?

```js
const response = await fetch("/api/projects");

if (!response.ok) {
  throw new Error("Falha na API");
}

const data = await response.json();
```

### Como evitar dois submits?

Use um estado de carregamento:

```js
if (saving.value) return;
saving.value = true;
```

E desabilite o botão:

```vue
<button :disabled="saving">Salvar</button>
```

---

## SQL

### O que é uma tabela?

É uma estrutura formada por linhas e colunas para armazenar dados.

### O que é uma chave primária?

É uma coluna que identifica cada registro de forma única.

### O que é uma chave estrangeira?

É uma coluna que referencia outra tabela e ajuda a manter relacionamentos válidos.

### O que é `JOIN`?

Combina dados de tabelas relacionadas:

```sql
SELECT p.name, t.title
FROM projects p
JOIN tasks t ON t.project_id = p.id;
```

### Diferença entre `INNER JOIN` e `LEFT JOIN`?

`INNER JOIN` retorna apenas registros que possuem correspondência nos dois lados.

`LEFT JOIN` mantém todos os registros da tabela da esquerda, mesmo sem correspondência.

### Para que serve `WHERE`?

Filtra linhas antes da agregação:

```sql
SELECT *
FROM projects
WHERE status = 'active';
```

### Para que serve `GROUP BY`?

Agrupa registros para usar funções como `COUNT`, `SUM` e `AVG`.

### Diferença entre `WHERE` e `HAVING`?

`WHERE` filtra linhas antes do agrupamento. `HAVING` filtra grupos depois do `GROUP BY`.

### O que é SQL injection?

É quando entrada do usuário altera indevidamente uma consulta SQL. Deve ser evitada com queries parametrizadas.

### O que é uma transação?

É um conjunto de operações tratado como uma unidade: todas são confirmadas ou todas são desfeitas.

### Como investigar uma divergência financeira?

Conferir período, filtros, duplicidades, status, timezone, arredondamento, casas decimais e fonte oficial dos dados. Depois comparar uma amostra com o ERP ou planilha.

---

## Testes e debugging

### O que é teste unitário?

Testa uma pequena parte isolada do código, como uma função.

### O que é teste de integração?

Testa a interação entre partes, como handler, store e banco.

### Como testar um handler Go?

Use `net/http/httptest` para criar uma requisição e capturar a resposta sem iniciar um servidor real.

### Para que serve `go test`?

Executa os testes do projeto.

```powershell
go test ./...
```

### Para que serve `go vet`?

Encontra construções suspeitas e possíveis problemas no código.

### Para que serve o race detector?

Detecta acessos concorrentes perigosos:

```powershell
go test -race ./...
```

### Como investigar um bug?

1. Reproduzir.
2. Registrar entrada e saída.
3. Ler o erro.
4. Isolar a camada: frontend, rota, handler, store ou banco.
5. Criar uma hipótese.
6. Fazer um teste pequeno.
7. Corrigir.
8. Repetir o teste original.
9. Criar um teste automatizado para não regredir.

---

## Git, Docker e infraestrutura

### O que é uma branch?

É uma linha independente de desenvolvimento para trabalhar sem alterar diretamente a branch principal.

### O que é um commit?

É um registro de uma mudança no histórico do projeto.

### O que é um Pull Request?

É uma proposta para revisar e integrar alterações em outra branch.

### Como resolver um conflito?

1. Ler os trechos conflitantes.
2. Escolher ou combinar a versão correta.
3. Remover marcadores do Git.
4. Testar.
5. Adicionar e continuar o processo.

### O que é Docker?

É uma forma de empacotar a aplicação e suas dependências em uma imagem reproduzível.

### O que é Terraform?

É uma ferramenta de infraestrutura como código. A infraestrutura é descrita em arquivos e pode ser revisada e aplicada de forma controlada.

### O que é GitHub Actions?

É uma ferramenta de automação para executar build, testes, lint e deploy.

### Onde guardar secrets?

Em secret managers ou variáveis protegidas do ambiente. Nunca no código, no Dockerfile ou em commits.

---

## AWS e integrações

### O que é uma Lambda?

É uma função executada sob demanda sem a necessidade de administrar um servidor diretamente.

### O que é API Gateway?

É uma porta de entrada para APIs, com roteamento, autenticação, limites e observabilidade.

### O que é SQS?

É uma fila de mensagens usada para desacoplar produtores e consumidores e processar tarefas de forma assíncrona.

### O que é S3?

É um serviço de armazenamento de objetos.

### O que é CloudWatch?

É usado para logs, métricas, alarmes e observabilidade na AWS.

### O que fazer quando o ERP está fora do ar?

Informar uma mensagem clara, não mostrar dados parciais como definitivos, registrar o erro, permitir nova tentativa quando apropriado e monitorar a dependência.

---

## Perguntas comportamentais

### O que você faz quando não sabe algo?

> Primeiro reproduzo o problema e leio a mensagem de erro. Depois separo o fluxo em partes, consulto a documentação e faço um teste pequeno para validar minha hipótese. Se continuar bloqueado, peço ajuda mostrando o que tentei, o resultado e a dúvida específica.

### Como você lida com regra de negócio ambígua?

> Eu não escolho sozinho uma regra que possa alterar valores ou processos. Apresento as opções, explico o impacto e confirmo com a pessoa responsável pelo negócio.

### Como você valida uma entrega?

> Testo o caminho principal, entradas inválidas, erros de dependências e o impacto em funcionalidades existentes. Quando envolve números, comparo com a fonte oficial e valido com a área.

### Conte sobre um projeto seu

Use esta estrutura:

1. qual problema existia;
2. qual foi sua responsabilidade;
3. qual tecnologia usou;
4. qual dificuldade apareceu;
5. como investigou;
6. qual resultado obteve;
7. o que faria diferente hoje.

### Por que quer esta vaga?

Uma resposta possível:

> Tenho experiência full-stack com interfaces, APIs e integração com dados. Quero evoluir em Go, Vue e sistemas internos, entendendo melhor o impacto do software no negócio. Gosto de investigar problemas e conversar com quem usa o sistema, não apenas implementar uma tela isolada.

### Qual é seu ponto forte?

Escolha algo comprovável, como:

- investigação de problemas;
- aprendizado rápido;
- comunicação com áreas;
- reutilização de componentes;
- cuidado com experiência do usuário.

### Qual é um ponto a desenvolver?

Escolha algo real e mostre ação:

> Ainda estou aprofundando Go e testes automatizados. Por isso montei este laboratório, implementei um CRUD, testei as rotas no Postman e estou adicionando testes para entender melhor os erros e os contratos da API.

### O que você faz se perceber que o prazo não será cumprido?

> Aviso cedo, explico o que já foi feito, o que bloqueou e quais são as opções. Posso reduzir escopo, dividir a entrega ou renegociar o prazo, mas não escondo o risco.

---

## Checklist de revisão

- [ ] Sei explicar polimorfismo com uma interface.
- [ ] Sei explicar `*`, `&`, ponteiro e `defer`.
- [ ] Sei diferenciar encapsulamento, abstração e herança.
- [ ] Sei criar e testar um CRUD.
- [ ] Sei explicar `200`, `201`, `400`, `404`, `409`, `500` e `503`.
- [ ] Sei investigar um `404` e um `500`.
- [ ] Sei usar `JOIN`, `GROUP BY`, `WHERE` e `HAVING`.
- [ ] Sei explicar `ref`, `computed` e `v-model`.
- [ ] Sei escrever um teste básico em Go.
- [ ] Sei explicar como valido dados financeiros.
- [ ] Sei falar sobre um projeto do meu currículo com contexto e resultado.
- [ ] Sei dizer o que não sei e como aprenderia.
