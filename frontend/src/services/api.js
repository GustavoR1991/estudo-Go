// =====================================================================
//  INTEGRAÇÃO COM A API GO — PASSO A PASSO
//
//  Como usar este arquivo:
//    - Cada PASSO tem a explicação e, logo abaixo, o código comentado.
//    - Leia, entenda e ESCREVA o código à mão embaixo do comentário
//      (ou apague o "//" das linhas). Depois apague a linha marcada
//      com "⛔ APAGUE" daquele passo.
//    - Faça na ordem: PASSO 1 → 2 → 3 aqui, depois vá para App.vue
//      (PASSOS 4 e 5), volte aqui no PASSO 6 e termine no CartDrawer.vue
//      (PASSO 7).
//
//  Antes de começar, em dois terminais:
//    go run .                   (API em http://localhost:8080)
//    cd frontend && npm run dev (front em http://localhost:5173)
//
//  O Vite repassa tudo que começa com /api para o Go (vite.config.js),
//  por isso usamos URL relativa: "/api/products" e não "http://localhost:8080/...".
// =====================================================================

const BASE_URL = "/api";

// ---------------------------------------------------------------------
// PASSO 1 — Função request(): o "carteiro" que fala com a API.
//
// Todas as outras funções vão usar esta, assim o fetch fica num lugar só.
// Ela recebe:
//   path    -> o pedaço da URL depois de /api, ex.: "/products"
//   options -> opções do fetch, ex.: { method: "POST", body: "..." }
//
// O que ela faz:
//   1. chama fetch juntando BASE_URL + path
//   2. avisa a API que vamos mandar/receber JSON (header Content-Type)
//   3. transforma a resposta em objeto JavaScript (response.json())
//   4. se o status NÃO for 2xx (response.ok === false), lança um erro
//      com a mensagem que o Go mandou. Nosso Go sempre manda { "error": "..." }
//   5. se deu tudo certo, devolve os dados
// ---------------------------------------------------------------------
async function request(path, options = {}) {
  const response = await fetch(`${BASE_URL}${path}`, {
    headers: { "Content-type": "application/json" },
    ...options,
  });

  const data = await response.json().catch(() => null);

  if (!response.ok) {
    throw new Error(data?.error || `Erro ${response.status}`);
  }

  return data;
}

export const api = {
  // -------------------------------------------------------------------
  // PASSO 2 — Listar categorias.
  //
  // Rota no Go: GET /api/categories  (main.go -> products.Categories)
  // Teste antes no navegador: http://localhost:8080/api/categories
  // Resposta: ["Acessórios", "Casa", "Moda", "Tecnologia", "Áudio"]
  //
  // GET é o padrão do fetch, então basta passar o caminho.
  // -------------------------------------------------------------------
  listCategories() {
    return request("/categories");
  },

  // -------------------------------------------------------------------
  // PASSO 3 — Listar produtos (com busca e categoria).
  //
  // Rota no Go: GET /api/products?q=fone&category=Áudio
  // Teste antes:  http://localhost:8080/api/products
  //               http://localhost:8080/api/products?q=fone
  // Resposta: [{ id, name, description, category, price, oldPrice, image, stock, featured }]
  //
  // URLSearchParams monta a "query string" (?q=...&category=...) e já
  // trata acentos e espaços. Só adicionamos o filtro se ele não estiver vazio.
  // -------------------------------------------------------------------
  listProducts({ q = "", category = "" } = {}) {
    const params = new URLSearchParams();
    if (q) params.set("q", q);
    if (category) params.set("category", category);

    return request(`/products?${params}`);
  },

  // ✋ Agora vá para src/App.vue e faça os PASSOS 4 e 5.

  // -------------------------------------------------------------------
  // PASSO 6 — Criar pedido (checkout).
  //
  // Rota no Go: POST /api/orders  (internal/orders/handler.go)
  //
  // Recebe:
  //   customer -> { name, email, address }   (vem do formulário)
  //   items    -> itens do carrinho: [{ id, name, price, image, stock, quantity }]
  //
  // Mas a API quer SÓ o id e a quantidade (preço quem calcula é o servidor,
  // assim ninguém compra mais barato mexendo no JSON). Formato esperado:
  //   {
  //     "customer": { "name": "Ana", "email": "ana@email.com", "address": "Rua A, 10" },
  //     "items": [{ "productId": 1, "quantity": 2 }]
  //   }
  //
  // Por isso usamos items.map(...) para converter cada item.
  // JSON.stringify transforma o objeto JS em texto JSON para enviar no body.
  //
  // Respostas possíveis:
  //   201 -> pedido criado: { id, customer, items, subtotal, shipping, total, status, createdAt }
  //   400 -> dados inválidos (ex.: e-mail errado)
  //   404 -> produto não existe
  //   409 -> sem estoque
  // Os erros viram exceção no request() do PASSO 1 — quem chama trata no catch.
  // -------------------------------------------------------------------
  createOrder(customer, items) {
    const body = {
      customer,
      items: items.map((item) => ({
        productId: item.id,
        quantity: item.quantity,
      })),
    };
    return request("/orders", {
      method: "POST",
      body: JSON.stringify(body),
    });
    // const body = {
    //   customer,
    //   items: items.map((item) => ({
    //     productId: item.id,
    //     quantity: item.quantity,
    //   })),
    // };
    //
    // return request("/orders", {
    //   method: "POST",
    //   body: JSON.stringify(body),
    // });
  },

  // ✋ Agora vá para src/components/CartDrawer.vue e faça o PASSO 7.

  // -------------------------------------------------------------------
  // DESAFIO EXTRA — Buscar um pedido pelo id.
  // Rota: GET /api/orders/{id}
  // Tente sozinho! Dica: é parecido com o PASSO 2, mas com o id na URL
  // usando template string: `/orders/${id}`
  // -------------------------------------------------------------------
  getOrder(id) {
    return request(`/orders/${id}`);
  },
};
