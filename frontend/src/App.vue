<script setup>
import { onMounted, ref } from "vue";
import { api } from "./services/api";
import { useCart } from "./composables/useCart";
import ProductCard from "./components/ProductCard.vue";
import CartDrawer from "./components/CartDrawer.vue";

const cart = useCart();
const products = ref([]);
const categories = ref(["Todos"]);
const query = ref("");
const category = ref("Todos");
const loading = ref(false);
const error = ref("");
const showCart = ref(false);

// ---------------------------------------------------------------------
// PASSO 4 — Carregar os produtos na tela.
//
// Esta função é chamada:
//   - quando a página abre (onMounted, lá embaixo)
//   - quando aperta Enter na busca, clica num filtro ou no botão ↻
//   - depois que um pedido é finalizado (o estoque mudou)
//
// Ela usa a api.listProducts() que você fez no PASSO 3.
// "await" espera a resposta da API chegar antes de seguir.
// O que vier da API vai para products.value, e o Vue redesenha o grid sozinho.
//
// Os filtros vêm da tela:
//   query.value    -> texto digitado na busca (trim() tira espaços das pontas)
//   category.value -> categoria clicada. "Todos" significa SEM filtro, então manda "".
//
// loading mostra "Carregando produtos..." e error mostra a mensagem se der erro.
// ---------------------------------------------------------------------
async function loadProducts() {
  loading.value = true;
  error.value = "";
  try {
    products.value = await api.listProducts({
      q: query.value.trim(),
      category: category.value === "Todos" ? "" : category.value,
    });
  } catch (err) {
    error.value = `Não foi possível carregar os produtos: ${err.message}`;
  } finally {
    loading.value = false;
  }
}

// ---------------------------------------------------------------------
// PASSO 5 — Carregar as categorias (botões de filtro).
//
// Usa a api.listCategories() do PASSO 2.
// A API devolve ["Acessórios", "Casa", ...] e nós colocamos "Todos" na frente.
// Os "..." (spread) espalham os itens do array dentro do novo array:
//   ["Todos", ...["Casa", "Moda"]]  ->  ["Todos", "Casa", "Moda"]
//
// Se der erro, não tem problema: o filtro só mostra "Todos".
// ---------------------------------------------------------------------
async function loadCategories() {
  try {
    const lista = await api.listCategories();
    categories.value = ["Todos", ...lista];
  } catch {
    categories.value = ["todos"];
  }
}

// ✅ Teste: salve, abra http://localhost:5173 — os produtos e filtros devem aparecer.
//    Teste a busca (digite "fone" + Enter) e clique nas categorias.
// ✋ Depois volte para services/api.js e faça o PASSO 6.

function addToCart(product) {
  cart.add(product);
  showCart.value = true;
}

onMounted(() => {
  loadCategories();
  loadProducts();
});
</script>

<template>
  <div class="page">
    <header class="topbar">
      <a
        class="logo"
        href="#"
        @click.prevent="
          category = 'Todos';
          loadProducts();
        "
      >
        <span class="logo-mark">✦</span> casa<span>nova</span>
      </a>
      <nav>
        <a href="#catalogo">Catálogo</a>
        <a href="#beneficios">Por que Casa Nova?</a>
      </nav>
      <button
        class="cart-button"
        @click="showCart = true"
        aria-label="Abrir carrinho"
      >
        Sacola <b>{{ cart.count.value }}</b>
      </button>
    </header>

    <main>
      <section class="hero">
        <div>
          <p class="eyebrow">Coleção essencial 2026</p>
          <h1>Escolhas que<br /><i>fazem sentido.</i></h1>
          <p class="hero-copy">
            Produtos bem pensados para deixar seus dias mais leves, bonitos e
            práticos.
          </p>
          <a class="hero-link" href="#catalogo"
            >Explorar produtos <span>↘</span></a
          >
        </div>
        <div class="hero-art">
          <div class="sun"></div>
          <div class="hero-card">
            <img
              src="https://images.unsplash.com/photo-1494438639946-1ebd1d20bf85?auto=format&fit=crop&w=900&q=80"
              alt="Ambiente aconchegante"
            />
            <span>Design para viver bem</span>
          </div>
        </div>
      </section>

      <section id="catalogo" class="catalog">
        <div class="section-title">
          <div>
            <p class="eyebrow">O que você procura?</p>
            <h2>Encontre seu próximo favorito</h2>
          </div>
          <div class="search">
            <span>⌕</span>
            <input
              v-model="query"
              placeholder="Buscar produtos..."
              @keyup.enter="loadProducts"
            />
            <button
              v-if="query"
              @click="
                query = '';
                loadProducts();
              "
            >
              ×
            </button>
          </div>
        </div>
        <div class="filters">
          <button
            v-for="item in categories"
            :key="item"
            :class="{ active: category === item }"
            @click="
              category = item;
              loadProducts();
            "
          >
            {{ item }}
          </button>
          <button class="refresh" @click="loadProducts" :disabled="loading">
            ↻
          </button>
        </div>
        <p v-if="error" class="alert">{{ error }}</p>
        <div v-if="loading" class="state">Carregando produtos...</div>
        <div v-else-if="!products.length" class="state">
          Nenhum produto encontrado. Tente outra busca.
        </div>
        <div v-else class="grid">
          <ProductCard
            v-for="product in products"
            :key="product.id"
            :product="product"
            @add="addToCart"
          />
        </div>
      </section>

      <section id="beneficios" class="benefits">
        <div>
          <span>01</span><strong>Curadoria real</strong>
          <p>Só entra o que a gente usaria em casa.</p>
        </div>
        <div>
          <span>02</span><strong>Entrega simples</strong>
          <p>Frete grátis para compras acima de R$ 199.</p>
        </div>
        <div>
          <span>03</span><strong>Compra segura</strong>
          <p>Seus dados protegidos em todas as etapas.</p>
        </div>
      </section>
    </main>

    <footer>
      <span class="logo-mark">✦</span> casa nova
      <small>feito para viver melhor</small>
    </footer>

    <CartDrawer
      v-if="showCart"
      @close="showCart = false"
      @ordered="loadProducts"
    />
  </div>
</template>
