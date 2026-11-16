<script setup>
import { computed, onMounted, ref } from "vue";

const products = ref([]);
const query = ref("");
const category = ref("Todos");
const categoryOptions = ref(["Todos"]);
const cart = ref([]);
const loading = ref(false);
const error = ref("");
const showCart = ref(false);

const categories = computed(() => categoryOptions.value);
const filteredProducts = computed(() => products.value);
const cartCount = computed(() =>
  cart.value.reduce((sum, item) => sum + item.quantity, 0),
);
const cartTotal = computed(() =>
  cart.value.reduce((sum, item) => sum + item.price * item.quantity, 0),
);
const formatPrice = (value) =>
  new Intl.NumberFormat("pt-BR", { style: "currency", currency: "BRL" }).format(
    value,
  );

async function loadProducts() {
  loading.value = true;
  error.value = "";
  const params = new URLSearchParams();
  if (query.value.trim()) params.set("q", query.value.trim());
  if (category.value !== "Todos") params.set("category", category.value);
  try {
    const response = await fetch(`/api/products?${params}`);
    if (!response.ok) throw new Error("Não foi possível carregar os produtos.");
    products.value = await response.json();
    if (category.value === "Todos" && !query.value.trim()) {
      categoryOptions.value = [
        "Todos",
        ...new Set(products.value.map((product) => product.category)),
      ];
    }
  } catch (err) {
    error.value = err.message;
  } finally {
    loading.value = false;
  }
}

function addToCart(product) {
  if (!product.stock) return;
  const item = cart.value.find((entry) => entry.id === product.id);
  if (item) {
    if (item.quantity < product.stock) item.quantity += 1;
    return;
  }
  cart.value.push({ ...product, quantity: 1 });
}

function removeFromCart(item) {
  const index = cart.value.findIndex((entry) => entry.id === item.id);
  if (index !== -1) cart.value.splice(index, 1);
}

function changeQuantity(item, quantity) {
  const nextQuantity = item.quantity + amount;

  if (nextQuantity <= 0) {
    removeFromCart(item);
    return;
  }
  if (nextQuantity > item.stock) {
    item.quantity = item.stock;
  }
}

function checkout() {
  if (!cart.value.length) return;
  window.alert("Pedido demonstrativo criado com sucesso!");
  cart.value = [];
  showCart.value = false;
}

onMounted(loadProducts);
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
        Sacola <b>{{ cartCount }}</b>
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
        <div v-else-if="!filteredProducts.length" class="state">
          Nenhum produto encontrado. Tente outra busca.
        </div>
        <div v-else class="grid">
          <article
            v-for="product in filteredProducts"
            :key="product.id"
            class="product"
          >
            <div class="product-image">
              <span v-if="product.featured" class="badge">Destaque</span>
              <img :src="product.image" :alt="product.name" />
              <button
                v-if="product.stock"
                class="quick-add"
                @click="addToCart(product)"
              >
                + Adicionar
              </button>
            </div>
            <div class="product-info">
              <div>
                <small>{{ product.category }}</small
                ><span v-if="!product.stock" class="sold-out">Esgotado</span>
              </div>
              <h3>{{ product.name }}</h3>
              <p>{{ product.description }}</p>
              <div class="price">
                <strong>{{ formatPrice(product.price) }}</strong
                ><del v-if="product.oldPrice">{{
                  formatPrice(product.oldPrice)
                }}</del>
              </div>
            </div>
          </article>
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

    <div v-if="showCart" class="overlay" @click.self="showCart = false">
      <aside class="cart">
        <div class="cart-heading">
          <h2>Sua sacola</h2>
          <button @click="showCart = false">×</button>
        </div>
        <div v-if="!cart.length" class="cart-empty">Sua sacola está vazia.</div>
        <div v-else class="cart-items">
          <div v-for="item in cart" :key="item.id" class="cart-item">
            <img :src="item.image" :alt="item.name" />
            <div>
              <strong>{{ item.name }}</strong>

              <div class="quantity">
                <button @click="changeQuantity(item, -1)">−</button>
                <span>{{ item.quantity }}</span>
                <button @click="changeQuantity(item, 1)">+</button>
              </div>

              <span>{{ formatPrice(item.price * item.quantity) }}</span>
            </div>
            <button @click="removeFromCart(item)">Remover</button>
          </div>
          <div class="cart-total">
            <span>Total</span><strong>{{ formatPrice(cartTotal) }}</strong>
          </div>
          <button class="checkout" @click="checkout">Finalizar pedido</button>
        </div>
      </aside>
    </div>
  </div>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=DM+Sans:wght@400;500;600;700&family=Playfair+Display:ital,wght@0,600;1,600&display=swap");
:root {
  color: #24211e;
  background: #f7f5f0;
  font-family: "DM Sans", sans-serif;
  font-synthesis: none;
}
* {
  box-sizing: border-box;
}
body {
  margin: 0;
  min-width: 320px;
}
button,
input {
  font: inherit;
}
button {
  cursor: pointer;
}
.page {
  max-width: 1240px;
  margin: auto;
  padding: 0 42px;
}
.topbar {
  height: 92px;
  display: flex;
  align-items: center;
  gap: 42px;
  border-bottom: 1px solid #dedbd4;
}
.logo {
  color: #24211e;
  font-size: 20px;
  font-weight: 700;
  letter-spacing: -1px;
  text-decoration: none;
}
.logo span:not(.logo-mark) {
  font-weight: 400;
}
.logo-mark {
  display: inline-grid;
  place-items: center;
  width: 28px;
  height: 28px;
  margin-right: 7px;
  color: #f7f5f0;
  background: #b5583d;
  border-radius: 50%;
  font-size: 14px;
}
nav {
  display: flex;
  gap: 27px;
  flex: 1;
}
nav a {
  color: #817a72;
  font-size: 12px;
  text-decoration: none;
}
nav a:hover,
.hero-link {
  color: #b5583d;
}
.cart-button {
  border: 0;
  background: none;
  color: #24211e;
  font-size: 12px;
}
.cart-button b {
  display: inline-grid;
  place-items: center;
  width: 20px;
  height: 20px;
  margin-left: 6px;
  color: #fff;
  background: #b5583d;
  border-radius: 50%;
  font-size: 10px;
}
.hero {
  display: grid;
  grid-template-columns: 1fr 1fr;
  min-height: 470px;
  align-items: center;
  gap: 50px;
}
.eyebrow {
  margin: 0 0 15px;
  color: #b5583d;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 2px;
  text-transform: uppercase;
}
h1 {
  margin: 0;
  font-size: clamp(44px, 6vw, 75px);
  line-height: 0.98;
  letter-spacing: -3px;
  font-weight: 500;
}
h1 i {
  color: #b5583d;
  font-family: "Playfair Display", serif;
  font-weight: 600;
}
.hero-copy {
  max-width: 330px;
  margin: 25px 0 28px;
  color: #817a72;
  font-size: 14px;
  line-height: 1.65;
}
.hero-link {
  font-size: 12px;
  font-weight: 700;
  text-decoration: none;
}
.hero-link span {
  display: inline-block;
  margin-left: 10px;
  font-size: 18px;
}
.hero-art {
  position: relative;
  min-height: 340px;
}
.sun {
  position: absolute;
  top: 0;
  right: 3%;
  width: 220px;
  height: 220px;
  background: #ead9bd;
  border-radius: 50%;
}
.hero-card {
  position: absolute;
  right: 13%;
  bottom: 5px;
  width: min(335px, 80%);
  transform: rotate(3deg);
  background: #fff;
  padding: 12px 12px 36px;
  box-shadow: 20px 20px 0 #e5dfd3;
}
.hero-card img {
  display: block;
  width: 100%;
  height: 270px;
  object-fit: cover;
  filter: saturate(0.72);
}
.hero-card span {
  position: absolute;
  bottom: 12px;
  left: 17px;
  color: #817a72;
  font-size: 11px;
}
.catalog {
  padding: 65px 0 75px;
  border-top: 1px solid #dedbd4;
}
.section-title {
  display: flex;
  align-items: end;
  justify-content: space-between;
  gap: 30px;
}
h2 {
  margin: 0;
  font-size: 31px;
  letter-spacing: -1.5px;
  font-weight: 500;
}
.search {
  display: flex;
  align-items: center;
  width: 245px;
  border-bottom: 1px solid #beb9b0;
}
.search span {
  color: #b5583d;
  font-size: 24px;
}
.search input {
  width: 100%;
  padding: 10px;
  border: 0;
  outline: 0;
  color: #24211e;
  background: transparent;
  font-size: 12px;
}
.search button {
  border: 0;
  background: none;
  color: #817a72;
}
.filters {
  display: flex;
  gap: 24px;
  margin: 34px 0 25px;
}
.filters button {
  padding: 0 0 10px;
  border: 0;
  border-bottom: 1px solid transparent;
  color: #999189;
  background: none;
  font-size: 12px;
}
.filters button.active {
  color: #b5583d;
  border-color: #b5583d;
  font-weight: 700;
}
.filters .refresh {
  margin-left: auto;
  color: #b5583d;
  font-size: 18px;
}
.grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 30px 18px;
}
.product-image {
  position: relative;
  overflow: hidden;
  background: #e8e4dc;
}
.product-image img {
  display: block;
  width: 100%;
  aspect-ratio: 1 / 1.08;
  object-fit: cover;
  transition: transform 0.3s;
}
.product:hover .product-image img {
  transform: scale(1.04);
}
.badge {
  position: absolute;
  z-index: 1;
  top: 12px;
  left: 12px;
  padding: 6px 9px;
  color: #fff;
  background: #b5583d;
  font-size: 9px;
}
.quick-add {
  position: absolute;
  right: 12px;
  bottom: 12px;
  padding: 10px 12px;
  border: 0;
  color: #24211e;
  background: #fff;
  font-size: 11px;
  opacity: 0;
  transition: opacity 0.2s;
}
.product:hover .quick-add {
  opacity: 1;
}
.product-info {
  padding-top: 15px;
}
.product-info > div:first-child {
  display: flex;
  justify-content: space-between;
}
.product-info small {
  color: #a49d94;
  font-size: 10px;
  text-transform: uppercase;
  letter-spacing: 1px;
}
.sold-out {
  color: #b5583d;
  font-size: 10px;
}
h3 {
  margin: 7px 0 5px;
  font-size: 16px;
  font-weight: 600;
}
.product-info p {
  min-height: 38px;
  margin: 0 0 12px;
  color: #918a82;
  font-size: 11px;
  line-height: 1.5;
}
.price {
  display: flex;
  align-items: center;
  gap: 8px;
}
.price strong {
  font-size: 14px;
}
.price del {
  color: #aaa39a;
  font-size: 11px;
}
.benefits {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 25px;
  padding: 45px 0;
  border-top: 1px solid #dedbd4;
}
.benefits div {
  padding-right: 25px;
  border-right: 1px solid #dedbd4;
}
.benefits div:last-child {
  border: 0;
}
.benefits span {
  display: block;
  margin-bottom: 19px;
  color: #b5583d;
  font-size: 11px;
}
.benefits strong {
  font-size: 14px;
}
.benefits p {
  margin: 8px 0 0;
  color: #918a82;
  font-size: 12px;
}
footer {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 25px 0;
  border-top: 1px solid #dedbd4;
  color: #24211e;
  font-size: 12px;
}
footer .logo-mark {
  width: 21px;
  height: 21px;
  margin: 0;
  font-size: 11px;
}
footer small {
  margin-left: auto;
  color: #a49d94;
}
.alert,
.state {
  padding: 20px;
  color: #b5583d;
  font-size: 13px;
}
.overlay {
  position: fixed;
  z-index: 5;
  inset: 0;
  display: flex;
  justify-content: flex-end;
  background: #24211e66;
}
.cart {
  width: min(420px, 100%);
  padding: 32px;
  background: #f7f5f0;
  box-shadow: -10px 0 30px #0002;
}
.cart-heading {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #dedbd4;
  padding-bottom: 20px;
}
.cart-heading button {
  border: 0;
  background: none;
  font-size: 25px;
}
.cart-empty {
  padding: 50px 0;
  color: #918a82;
  font-size: 13px;
}
.cart-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px 0;
  border-bottom: 1px solid #dedbd4;
}
.cart-item img {
  width: 60px;
  height: 60px;
  object-fit: cover;
}
.cart-item div {
  flex: 1;
}
.cart-item strong,
.cart-item span {
  display: block;
  font-size: 12px;
}
.cart-item span {
  margin-top: 5px;
  color: #918a82;
}
.cart-item > button {
  border: 0;
  background: none;
  color: #b5583d;
  font-size: 10px;
}
.cart-total {
  display: flex;
  justify-content: space-between;
  padding: 25px 0;
}
.checkout {
  width: 100%;
  padding: 14px;
  border: 0;
  color: #fff;
  background: #b5583d;
  font-weight: 700;
}
@media (max-width: 700px) {
  .page {
    padding: 0 20px;
  }
  .topbar {
    height: 72px;
    gap: 15px;
  }
  nav {
    display: none;
  }
  .cart-button {
    margin-left: auto;
  }
  .hero {
    display: block;
    padding: 70px 0;
  }
  .hero-art {
    margin-top: 55px;
  }
  .section-title {
    display: block;
  }
  .search {
    width: 100%;
    margin-top: 25px;
  }
  .filters {
    gap: 13px;
    overflow-x: auto;
  }
  .grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 24px 12px;
  }
  h2 {
    font-size: 25px;
  }
  .product-info p {
    display: none;
  }
  .benefits {
    display: block;
  }
  .benefits div {
    padding: 20px 0;
    border-right: 0;
    border-bottom: 1px solid #dedbd4;
  }
  .hero-card {
    right: 4%;
  }
}

.quantity {
  display: flex;
  align-items: center;
  gap: 8px;
  margin: 8px 0;
}

.quantity button {
  width: 24px;
  height: 24px;
  border: 1px solid #dedbd4;
  background: #fff;
  border-radius: 50%;
}
</style>
