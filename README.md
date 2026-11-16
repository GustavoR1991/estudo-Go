# Desafios práticos — Go, Vue e APIs

Este arquivo contém somente desafios para você resolver. Consulte
[perguntas e respostas.md](<./perguntas%20e%20respostas.md>) apenas quando quiser revisar a teoria.

## Como iniciar

### Backend

```powershell
go run .
```

Servidor:

```text
http://localhost:8080
```

### Frontend

```powershell
cd frontend
npm run dev
```

### Validação

```powershell
go test ./...
go vet ./...
```

Porta diferente (se a 8080 estiver ocupada): `$env:PORT="8081"; go run .`

### API do e-commerce

Dados ficam em memória durante a execução. Erros sempre vêm como `{"error": "mensagem"}`.

```text
GET    /api/categories                        lista de categorias
GET    /api/products?q=fone&category=Áudio&featured=true
POST   /api/products
GET    /api/products/{id}
PUT    /api/products/{id}
DELETE /api/products/{id}

POST   /api/orders                            checkout (cria pedido e baixa estoque)
GET    /api/orders?email=ana@email.com
GET    /api/orders/{id}
```

Um produto precisa de `name`, `category`, `price` maior que zero e `stock` maior
ou igual a zero.

Exemplo de checkout — o front manda só `productId` e `quantity`; preço, nome e
frete são calculados no servidor (frete R$ 19,90, grátis a partir de R$ 199):

```json
POST /api/orders
{
  "customer": { "name": "Ana", "email": "ana@email.com", "address": "Rua A, 10" },
  "items": [{ "productId": 1, "quantity": 2 }]
}
```

Respostas: `201` pedido criado · `400` dados inválidos · `404` produto não existe ·
`409` estoque insuficiente (nenhum estoque é alterado).

### Organização do código

```text
main.go                         rotas + serve o front compilado (frontend/dist)
internal/httpx/                 helpers JSON / erro / leitura de {id}
internal/products/              catálogo (handler.go = HTTP, store.go = dados + estoque)
internal/orders/                checkout e pedidos (+ handler_test.go)

frontend/src/
  services/api.js               ÚNICO lugar que chama fetch("/api/...")
  composables/useCart.js        estado do carrinho (salvo no localStorage)
  components/ProductCard.vue    card de produto
  components/CartDrawer.vue     sacola → formulário → pedido confirmado
  App.vue                       página: busca, filtros, grid
  style.css                     estilos globais
```

### Sua tarefa: integrar a API no front

O código já está escrito **comentado**, com explicação em cada passo. Você
reescreve à mão embaixo do comentário e apaga a linha marcada com `⛔ APAGUE`.
Siga a ordem:

| Passo | Arquivo → função | O que faz |
|---|---|---|
| 1 | `services/api.js` → `request()` | função base do `fetch` + tratamento de erro |
| 2 | `services/api.js` → `listCategories()` | `GET /api/categories` |
| 3 | `services/api.js` → `listProducts()` | `GET /api/products?q=&category=` |
| 4 | `App.vue` → `loadProducts()` | mostra os produtos na tela |
| 5 | `App.vue` → `loadCategories()` | mostra os filtros |
| 6 | `services/api.js` → `createOrder()` | `POST /api/orders` |
| 7 | `components/CartDrawer.vue` → `submitOrder()` | finaliza o pedido |
| extra | `services/api.js` → `getOrder()` | sem resposta pronta, tente sozinho |

Teste cada rota primeiro no navegador/Postman (ex.: `http://localhost:8080/api/products`)
para ver o JSON antes de escrever o `fetch`.

O front-end em desenvolvimento (`npm run dev`) encaminha automaticamente `/api`
para o Go em `http://localhost:8080` (ver `frontend/vite.config.js`). Depois de
mexer no front, rode `npm run build` para o `go run .` servir a versão nova.

---

## Desafio 10 — API de produtos

Crie um novo recurso `products`.

Campos:

```json
{
  "id": 1,
  "name": "Cabo elétrico",
  "price": 25.9,
  "active": true
}
```

Rotas:

```text
GET    /api/products
POST   /api/products
GET    /api/products/{id}
PUT    /api/products/{id}
DELETE /api/products/{id}
```

Regras:

- nome obrigatório;
- preço maior ou igual a zero;
- preço com no máximo duas casas decimais;
- ID automático;
- produto inexistente retorna `404`;
- JSON inválido retorna `400`.

---

## Desafio 11 — Paginação

Adicione paginação em:

```text
GET /api/products?page=1&limit=10
```

Resposta esperada:

```json
{
  "items": [],
  "page": 1,
  "limit": 10,
  "total": 0
}
```

Teste:

- página inexistente;
- `limit` negativo;
- `limit` muito alto;
- parâmetros ausentes;
- valores que não são números.

---

## Desafio 12 — Busca e filtros

Adicione:

```text
GET /api/products?minPrice=100&maxPrice=400
GET /api/products?inStock=true
```

Os filtros devem funcionar junto com `q` e `category`, que já existem. Responda `400`
se `minPrice` ou `maxPrice` não forem números.

---

## Desafio 13 — SQL

Crie as tabelas:

```sql
projects(id, name, status, created_at)
tasks(id, project_id, title, done)
```

Escreva consultas para:

- listar projetos ativos;
- contar tarefas por projeto;
- contar tarefas concluídas;
- encontrar projetos sem tarefas;
- encontrar tarefas duplicadas;
- calcular percentual de conclusão;
- buscar projetos por nome;
- ordenar pelo mais recente.

---

## Desafio 14 — Dados financeiros

Simule uma divergência entre API e planilha.

Investigue:

- período usado;
- filtros;
- registros duplicados;
- status;
- arredondamento;
- casas decimais;
- timezone;
- fonte oficial;
- amostra manual.

Produza uma explicação curta contendo:

- valor da API;
- valor da planilha;
- diferença;
- causa;
- correção;
- como evitar regressão.

---

## Desafio 15 — ERP indisponível

Simule uma dependência externa fora do ar.

Implemente:

- timeout;
- mensagem amigável;
- log técnico;
- status adequado;
- tentativa novamente;
- nenhum dado parcial apresentado como definitivo.

---

## Desafio 16 — Git

Faça o fluxo:

```powershell
git switch -c feat/project-validation
git status
git add .
git commit -m "feat: validate project status"
git log --oneline -5
```

Depois pratique:

- criar conflito;
- resolver conflito;
- revisar diff;
- abrir uma branch para bug;
- criar um commit pequeno;
- escrever uma descrição de Pull Request.

---

## Desafio 17 — Docker

Crie um `Dockerfile` para o backend.

Requisitos:

- build em múltiplas etapas;
- imagem final pequena;
- porta documentada;
- configuração por variável de ambiente;
- aplicação iniciando com um comando simples.

---

## Desafio 18 — GitHub Actions

Crie um workflow que execute:

```text
go test ./...
go vet ./...
npm run build
```

Faça o workflow falhar quando qualquer etapa falhar.

---

## Desafio 19 — Terraform

Crie uma configuração simples para representar:

- aplicação;
- banco;
- bucket;
- variável de ambiente;
- ambiente de desenvolvimento.

Documente como inicializar, validar e planejar:

```powershell
terraform init
terraform validate
terraform plan
```

---

## Desafio 20 — Simulação de teste ao vivo

Implemente, sem copiar uma solução pronta, uma API de pedidos.

Pedido:

```json
{
  "customer": "Cliente A",
  "total": 150.5,
  "status": "pending"
}
```

Requisitos:

- CRUD;
- validação;
- total não negativo;
- status válido;
- erros HTTP corretos;
- testes automatizados;
- teste no Postman;
- documentação das rotas;
- tratamento de concorrência.

Durante a execução, explique:

1. o que está fazendo;
2. qual hipótese está testando;
3. o que a mensagem de erro significa;
4. o que já foi validado;
5. o que ainda falta validar.

---

## Checklist final

- [ ] Implementei um CRUD completo.
- [ ] Testei todas as rotas no Postman.
- [ ] Testei casos de erro.
- [ ] Escrevi testes automatizados.
- [ ] Executei `go test -race ./...`.
- [ ] Criei uma tela Vue consumindo a API.
- [ ] Tratei loading e erro.
- [ ] Pratiquei SQL com `JOIN` e `GROUP BY`.
- [ ] Simulei divergência de dados.
- [ ] Pratiquei Git.
- [ ] Entendi o fluxo de Docker e CI.
- [ ] Consigo explicar meu raciocínio durante um teste ao vivo.
