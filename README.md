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

---

## Desafio 1 — CRUD de projetos

Implemente completamente estas rotas:

```text
GET    /api/projects
POST   /api/projects
GET    /api/projects/{id}
PUT    /api/projects/{id}
DELETE /api/projects/{id}
```

O projeto deve possuir:

```json
{
  "id": 1,
  "name": "Projeto Go",
  "description": "Estudo de API",
  "status": "active"
}
```

Regras:

- `name` é obrigatório;
- `status` aceita somente `active`, `paused` ou `completed`;
- IDs devem ser únicos;
- projeto inexistente deve retornar `404`;
- JSON inválido deve retornar `400`;
- criação deve retornar `201`;
- atualização deve retornar `200`;
- exclusão deve retornar `200` ou `204`.

Teste todas as rotas no Postman.

---

## Desafio 2 — Criar projeto pelo Postman

Faça:

```text
POST http://localhost:8080/api/projects
```

Header:

```text
Content-Type: application/json
```

Body:

```json
{
  "name": "Projeto de compras",
  "description": "Conferência de pedidos",
  "status": "active"
}
```

Teste também:

```json
{}
```

```json
{
  "name": ""
}
```

```json
{
  "name": "Projeto",
  "status": "invalido"
}
```

```json
{
  "name": "Projeto"
```

Para cada caso, confirme o status HTTP e o corpo da resposta.

---

## Desafio 3 — Buscar projeto

Teste:

```text
GET http://localhost:8080/api/projects/1
```

Depois teste:

```text
GET http://localhost:8080/api/projects/999
GET http://localhost:8080/api/projects/abc
```

Defina respostas diferentes para:

- ID existente;
- ID inexistente;
- ID inválido.

---

## Desafio 4 — Atualizar projeto

Faça:

```text
PUT http://localhost:8080/api/projects/1
```

Body:

```json
{
  "name": "Projeto atualizado",
  "description": "Descrição nova",
  "status": "completed"
}
```

Depois confirme com:

```text
GET http://localhost:8080/api/projects/1
```

Teste atualização de:

- projeto inexistente;
- ID inválido;
- JSON inválido;
- nome vazio;
- status inválido.

---

## Desafio 5 — Excluir projeto

Faça:

```text
DELETE http://localhost:8080/api/projects/1
```

Depois confirme que o projeto não existe mais:

```text
GET http://localhost:8080/api/projects/1
```

Teste também:

```text
DELETE http://localhost:8080/api/projects/999
DELETE http://localhost:8080/api/projects/abc
```

---

## Desafio 6 — IDs e memória

Investigue e corrija:

- IDs repetidos;
- `nextID` iniciado com valor incorreto;
- comportamento após excluir um projeto;
- comportamento após reiniciar o servidor;
- criação simultânea de projetos.

Execute:

```powershell
go test -race ./...
```

---

## Desafio 7 — Testes automatizados em Go

Crie testes para:

- listar projetos;
- criar projeto válido;
- rejeitar JSON inválido;
- rejeitar nome vazio;
- rejeitar status inválido;
- buscar projeto existente;
- retornar `404` para projeto inexistente;
- atualizar projeto;
- excluir projeto;
- tentar excluir projeto inexistente.

Use:

```go
net/http/httptest
```

Execute:

```powershell
go test ./... -v
```

---

## Desafio 8 — Frontend de projetos

Crie uma tela Vue para projetos com:

- listagem;
- formulário de criação;
- edição;
- exclusão;
- filtro por status;
- estado de carregamento;
- estado vazio;
- mensagem de erro;
- confirmação antes de excluir.

Use as APIs:

```text
GET    /api/projects
POST   /api/projects
PUT    /api/projects/{id}
DELETE /api/projects/{id}
```

---

## Desafio 9 — Bugs no frontend

Crie e depois corrija estes problemas:

- lista não atualiza após criar;
- botão permite dois submits;
- erro da API não aparece;
- filtro mostra dados errados;
- exclusão remove o item visual errado;
- loading não termina quando ocorre erro;
- frontend usa `title` e API usa `name`;
- frontend usa `completed` e API usa `done`.

Para cada bug:

1. reproduza;
2. descreva o sintoma;
3. encontre a causa;
4. corrija;
5. teste novamente.

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
GET /api/projects?page=1&limit=10
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
GET /api/projects?status=active
GET /api/projects?search=go
```

Faça a busca funcionar sem diferenciar letras maiúsculas e minúsculas.

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
