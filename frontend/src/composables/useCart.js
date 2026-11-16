import { computed, ref, watch } from "vue";

// Mesmas regras do backend (internal/orders/store.go). O valor final oficial
// é sempre o que a API devolve no pedido; aqui é só a prévia.
export const FREE_SHIPPING_FROM = 199;
export const SHIPPING_PRICE = 19.9;

const STORAGE_KEY = "casanova:cart";

function loadSaved() {
  try {
    return JSON.parse(localStorage.getItem(STORAGE_KEY)) || [];
  } catch {
    return [];
  }
}

// Estado fora da função = compartilhado entre todos os componentes que usam o carrinho.
const items = ref(loadSaved());
watch(items, (value) => localStorage.setItem(STORAGE_KEY, JSON.stringify(value)), {
  deep: true,
});

export function useCart() {
  const count = computed(() => items.value.reduce((sum, item) => sum + item.quantity, 0));
  const subtotal = computed(() =>
    items.value.reduce((sum, item) => sum + item.price * item.quantity, 0),
  );
  const shipping = computed(() =>
    !items.value.length || subtotal.value >= FREE_SHIPPING_FROM ? 0 : SHIPPING_PRICE,
  );
  const total = computed(() => subtotal.value + shipping.value);

  function add(product) {
    if (!product.stock) return;
    const item = items.value.find((entry) => entry.id === product.id);
    if (item) {
      item.quantity = Math.min(item.quantity + 1, item.stock);
      return;
    }
    const { id, name, price, image, stock } = product;
    items.value.push({ id, name, price, image, stock, quantity: 1 });
  }

  function remove(id) {
    items.value = items.value.filter((item) => item.id !== id);
  }

  function changeQuantity(item, amount) {
    const next = item.quantity + amount;
    if (next <= 0) return remove(item.id);
    item.quantity = Math.min(next, item.stock);
  }

  function clear() {
    items.value = [];
  }

  return { items, count, subtotal, shipping, total, add, remove, changeQuantity, clear };
}
