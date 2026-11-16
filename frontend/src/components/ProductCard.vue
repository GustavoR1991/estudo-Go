<script setup>
import { formatPrice } from "../utils";

defineProps({ product: { type: Object, required: true } });
defineEmits(["add"]);
</script>

<template>
  <article class="product">
    <div class="product-image">
      <span v-if="product.featured" class="badge">Destaque</span>
      <img :src="product.image" :alt="product.name" />
      <button v-if="product.stock" class="quick-add" @click="$emit('add', product)">
        + Adicionar
      </button>
    </div>
    <div class="product-info">
      <div>
        <small>{{ product.category }}</small>
        <span v-if="!product.stock" class="sold-out">Esgotado</span>
      </div>
      <h3>{{ product.name }}</h3>
      <p>{{ product.description }}</p>
      <div class="price">
        <strong>{{ formatPrice(product.price) }}</strong>
        <del v-if="product.oldPrice">{{ formatPrice(product.oldPrice) }}</del>
      </div>
    </div>
  </article>
</template>
