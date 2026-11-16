<script setup>
import { reactive, ref } from "vue";
import { api } from "../services/api";
import { FREE_SHIPPING_FROM, useCart } from "../composables/useCart";
import { formatPrice } from "../utils";

const emit = defineEmits(["close", "ordered"]);
const cart = useCart();

// "cart" -> lista de itens | "checkout" -> formulário | "done" -> pedido criado
const step = ref("cart");
const sending = ref(false);
const error = ref("");
const order = ref(null);
const customer = reactive({ name: "", email: "", address: "" });

// ---------------------------------------------------------------------
// PASSO 7 — Finalizar o pedido.
//
// Roda quando o formulário é enviado (botão "Finalizar pedido").
// Usa a api.createOrder() do PASSO 6, mandando:
//   { ...customer }    -> cópia dos dados do formulário { name, email, address }
//   cart.items.value   -> os itens da sacola
//
// Se a API aceitar (201), ela devolve o pedido pronto. Então:
//   1. guardamos o pedido em order.value (a tela de "Pedido #X confirmado" usa ele)
//   2. limpamos a sacola
//   3. trocamos a tela para "done"
//   4. avisamos o App.vue com emit("ordered") para ele recarregar os produtos
//      (o estoque diminuiu!)
//
// Se a API recusar (400, 404, 409), o request() lança um erro, o código pula
// direto para o catch e a mensagem do Go aparece em vermelho na sacola.
// sending desabilita o botão enquanto espera, para não criar pedido duplicado.
// ---------------------------------------------------------------------
async function submitOrder() {
  sending.value = true;
  error.value = "";
  try {
    order.value = await api.createOrder({ ...customer }, cart.items.value);
    cart.clear();
    step.value = "done";
    emit("ordered", order.value);

    throw new Error("PASSO 7 pendente: finalize o pedido em CartDrawer.vue"); // ⛔ APAGUE depois do PASSO 7
  } catch (err) {
    error.value = err.message;
  } finally {
    sending.value = false;
  }
}

</script>

<template>
  <div class="overlay" @click.self="emit('close')">
    <aside class="cart">
      <div class="cart-heading">
        <h2>{{ step === "checkout" ? "Seus dados" : "Sua sacola" }}</h2>
        <button aria-label="Fechar" @click="emit('close')">×</button>
      </div>

      <!-- Pedido confirmado -->
      <div v-if="step === 'done'" class="order-success">
        <span class="logo-mark">✓</span>
        <h2>Pedido #{{ order.id }} confirmado!</h2>
        <p>
          Obrigado, {{ order.customer.name }}. Total pago:
          <strong>{{ formatPrice(order.total) }}</strong
          ><br />Enviaremos as novidades para {{ order.customer.email }}.
        </p>
        <button class="checkout" @click="emit('close')">
          Continuar comprando
        </button>
      </div>

      <div v-else-if="!cart.items.value.length" class="cart-empty">
        Sua sacola está vazia.
      </div>

      <template v-else>
        <!-- Itens -->
        <div v-if="step === 'cart'" class="cart-items">
          <div
            v-for="item in cart.items.value"
            :key="item.id"
            class="cart-item"
          >
            <img :src="item.image" :alt="item.name" />
            <div>
              <strong>{{ item.name }}</strong>
              <div class="quantity">
                <button @click="cart.changeQuantity(item, -1)">−</button>
                <span>{{ item.quantity }}</span>
                <button
                  :disabled="item.quantity >= item.stock"
                  @click="cart.changeQuantity(item, 1)"
                >
                  +
                </button>
              </div>
              <span>{{ formatPrice(item.price * item.quantity) }}</span>
            </div>
            <button @click="cart.remove(item.id)">Remover</button>
          </div>
        </div>

        <!-- Formulário de checkout -->
        <form
          v-else
          id="checkout-form"
          class="checkout-form"
          @submit.prevent="submitOrder"
        >
          <label
            >Nome<input v-model="customer.name" required autocomplete="name"
          /></label>
          <label
            >E-mail<input
              v-model="customer.email"
              type="email"
              required
              autocomplete="email"
          /></label>
          <label
            >Endereço<input
              v-model="customer.address"
              required
              autocomplete="street-address"
          /></label>
        </form>

        <div class="cart-summary">
          <div>
            <span>Subtotal</span
            ><span>{{ formatPrice(cart.subtotal.value) }}</span>
          </div>
          <div>
            <span>Frete</span>
            <span>{{
              cart.shipping.value ? formatPrice(cart.shipping.value) : "Grátis"
            }}</span>
          </div>
          <small v-if="cart.shipping.value">
            Faltam
            {{ formatPrice(FREE_SHIPPING_FROM - cart.subtotal.value) }} para
            frete grátis
          </small>
          <div class="grand">
            <span>Total</span
            ><strong>{{ formatPrice(cart.total.value) }}</strong>
          </div>
        </div>

        <p v-if="error" class="alert">{{ error }}</p>

        <button
          v-if="step === 'cart'"
          class="checkout"
          @click="step = 'checkout'"
        >
          Continuar para pagamento
        </button>
        <template v-else>
          <button
            class="checkout"
            type="submit"
            form="checkout-form"
            :disabled="sending"
          >
            {{ sending ? "Enviando..." : "Finalizar pedido" }}
          </button>
          <button class="link-button" @click="step = 'cart'">
            ← Voltar para a sacola
          </button>
        </template>
      </template>
    </aside>
  </div>
</template>
