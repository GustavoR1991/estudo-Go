<script setup>
import { computed, onMounted, ref } from "vue";

const tasks = ref([]);
const title = ref("");
const loading = ref(false);
const saving = ref(false);
const error = ref("");
const filter = ref("all");

// Exercício: a API chama esse campo de "done". Por que a aba de concluídas não funciona?
const completedTasks = computed(() => tasks.value.filter((task) => task.done));
const pendingTasks = computed(() => tasks.value.filter((task) => !task.done));
const progress = computed(() => {
  if (!tasks.value.length) return 0;
  return Math.round((completedTasks.value.length / tasks.value.length) * 100);
});

const visibleTasks = computed(() => {
  if (filter.value === "pending") return pendingTasks.value;
  if (filter.value === "completed") return completedTasks.value;
  return tasks.value;
});

async function loadTasks() {
  loading.value = true;
  error.value = "";

  try {
    const response = await fetch("/api/tasks");
    if (!response.ok) throw new Error("Falha ao carregar tarefas");
    tasks.value = await response.json();
  } catch {
    error.value = "Não foi possível carregar as tarefas.";
  } finally {
    loading.value = false;
  }
}

async function createTask() {
  const taskTitle = title.value.trim();

  if (!taskTitle) {
    error.value = "Digite um título para continuar.";
    return;
  }

  const taskAlreadyExists = tasks.value.some(
    (task) => task.title.toLowerCase() === taskTitle.toLowerCase(),
  );

  if (taskAlreadyExists) {
    error.value = "Tarefa já existe.";
    return;
  }

  saving.value = true;
  error.value = "";
  try {
    const response = await fetch("/api/tasks", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: taskTitle }),
    });

    if (!response.ok) {
      const body = await response.json().catch(() => ({}));
      throw new Error(body.error || "Não foi possível criar a tarefa.");
    }

    tasks.value.push(await response.json());
    title.value = "";
  } catch (err) {
    error.value = err.message;
  } finally {
    saving.value = false;
  }
}

async function toggleTask(task) {
  const response = await fetch(`/api/tasks/${task.id}`, {
    method: "PATCH",
  });

  const updatedTask = await response.json();
  const index = tasks.value.findIndex((item) => item.id === updatedTask.id);

  if (index !== -1) {
    tasks.value[index] = updatedTask;
  }

  if (!response.ok) {
    throw new Error("Não foi possível atualizar a tarefa.");
  }
}

onMounted(loadTasks);
</script>

<template>
  <main class="shell">
    <section class="hero">
      <div class="brand">
        <div class="brand-mark">✓</div>
        <span>Flowboard</span>
      </div>
      <div class="hero-copy">
        <p class="eyebrow">Seu espaço de foco</p>
        <h1>Pequenos passos.<br /><em>Grandes entregas.</em></h1>
        <p class="subtitle">
          Organize o que importa e acompanhe seu ritmo sem complicação.
        </p>
      </div>
      <div class="progress-card">
        <div class="progress-heading">
          <span>Progresso geral</span>
          <strong>{{ progress }}%</strong>
        </div>
        <div class="progress-track">
          <div class="progress-value" :style="{ width: `${progress}%` }"></div>
        </div>
        <small
          >{{ completedTasks.length }} de {{ tasks.length }} tarefas
          concluídas</small
        >
      </div>
    </section>

    <section class="workspace">
      <div class="section-heading">
        <div>
          <p class="eyebrow">Visão geral</p>
          <h2>Minhas tarefas</h2>
        </div>
        <div class="stats">
          <span
            ><b>{{ pendingTasks.length }}</b> pendentes</span
          >
          <span
            ><b>{{ completedTasks.length }}</b> concluídas</span
          >
        </div>
      </div>

      <form class="composer" @submit.prevent="createTask">
        <span class="plus">+</span>
        <input
          v-model="title"
          type="text"
          placeholder="O que você quer realizar hoje?"
        />
        <button type="submit" :disabled="saving">
          {{ saving ? "Salvando..." : "Adicionar tarefa" }}
        </button>
      </form>

      <div class="toolbar">
        <div class="tabs">
          <button :class="{ active: filter === 'all' }" @click="filter = 'all'">
            Todas <span>{{ tasks.length }}</span>
          </button>
          <button
            :class="{ active: filter === 'pending' }"
            @click="filter = 'pending'"
          >
            Pendentes <span>{{ pendingTasks.length }}</span>
          </button>
          <button
            :class="{ active: filter === 'completed' }"
            @click="filter = 'completed'"
          >
            Concluídas <span>{{ completedTasks.length }}</span>
          </button>
        </div>
        <button class="refresh" @click="loadTasks" :disabled="loading">
          ↻ Atualizar
        </button>
      </div>

      <p v-if="error" class="alert">{{ error }}</p>
      <div v-if="loading" class="empty-state">
        <span class="spinner"></span> Carregando seu espaço...
      </div>
      <div v-else-if="!visibleTasks.length" class="empty-state">
        <span class="empty-icon">✦</span>
        <strong>Nada por aqui ainda</strong>
        <span
          >Adicione uma tarefa e comece a transformar seus planos em ação.</span
        >
      </div>
      <ul v-else class="task-list">
        <li
          v-for="task in visibleTasks"
          :key="task.id"
          class="task-item"
          :class="{ completed: task.done }"
        >
          <button
            type
            class="check"
            :aria-label="task.done ? 'Reabrir tarefa' : 'Concluir tarefa'"
            @click="toggleTask(task)"
          >
            <span v-if="task.done">✓</span>
          </button>
          <span class="task-title">{{ task.title }}</span>
          <span class="task-status">{{
            task.done ? "Concluída" : "Em andamento"
          }}</span>
        </li>
      </ul>
    </section>
    <footer>Flowboard <span>•</span> construído com Go + Vue</footer>
  </main>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=DM+Sans:wght@400;500;600;700&family=Space+Grotesk:wght@500;600;700&display=swap");

:root {
  color: #e8e7f2;
  background: #11121a;
  font-family: "DM Sans", sans-serif;
  font-synthesis: none;
}
* {
  box-sizing: border-box;
}
body {
  margin: 0;
  min-width: 320px;
  background:
    radial-gradient(circle at 85% 5%, #293146 0, transparent 30%), #11121a;
}
button,
input {
  font: inherit;
}
button {
  cursor: pointer;
}
.shell {
  width: min(1050px, calc(100% - 40px));
  margin: 0 auto;
  padding: 38px 0 28px;
}
.hero {
  position: relative;
  min-height: 350px;
  padding: 3px 4px 70px;
}
.brand {
  display: flex;
  align-items: center;
  gap: 10px;
  color: #faf9ff;
  font-weight: 700;
  letter-spacing: -0.3px;
}
.brand-mark {
  display: grid;
  place-items: center;
  width: 31px;
  height: 31px;
  color: #14151d;
  background: #c9f27a;
  border-radius: 9px;
  font-size: 19px;
  font-weight: 700;
}
.hero-copy {
  margin-top: 68px;
  max-width: 600px;
}
.eyebrow {
  margin: 0 0 12px;
  color: #a6b0c5;
  text-transform: uppercase;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 1.8px;
}
h1,
h2 {
  margin: 0;
  color: #f9f8ff;
  font-family: "Space Grotesk", sans-serif;
  letter-spacing: -1.7px;
}
h1 {
  font-size: clamp(42px, 6vw, 66px);
  line-height: 1.02;
}
h1 em {
  color: #c9f27a;
  font-style: normal;
}
.subtitle {
  max-width: 410px;
  margin: 22px 0 0;
  color: #9da2b7;
  font-size: 16px;
  line-height: 1.6;
}
.progress-card {
  position: absolute;
  right: 3px;
  bottom: 69px;
  width: 255px;
  padding: 19px;
  background: #1b1d29cc;
  border: 1px solid #ffffff0d;
  border-radius: 16px;
  box-shadow: 0 20px 50px #00000025;
}
.progress-heading {
  display: flex;
  justify-content: space-between;
  margin-bottom: 15px;
  color: #b6b9c9;
  font-size: 12px;
}
.progress-heading strong {
  color: #c9f27a;
  font-size: 18px;
}
.progress-track {
  height: 7px;
  overflow: hidden;
  background: #303342;
  border-radius: 99px;
}
.progress-value {
  height: 100%;
  background: #c9f27a;
  border-radius: inherit;
  transition: width 0.3s ease;
}
.progress-card small {
  display: block;
  margin-top: 12px;
  color: #777d94;
  font-size: 11px;
}
.workspace {
  padding: 36px 40px 42px;
  background: #f8f8fc;
  border-radius: 26px;
  color: #212331;
  box-shadow: 0 25px 70px #00000033;
}
.section-heading {
  display: flex;
  align-items: end;
  justify-content: space-between;
  margin-bottom: 28px;
}
.workspace .eyebrow {
  color: #9699a8;
  margin-bottom: 8px;
}
h2 {
  color: #20212c;
  font-size: 31px;
  letter-spacing: -1px;
}
.stats {
  display: flex;
  gap: 23px;
  color: #9496a5;
  font-size: 12px;
}
.stats b {
  color: #242633;
  font-size: 18px;
  margin-right: 5px;
}
.composer {
  display: flex;
  align-items: center;
  gap: 13px;
  padding: 7px 8px 7px 17px;
  background: #fff;
  border: 1px solid #e8e8f0;
  border-radius: 13px;
  box-shadow: 0 8px 20px #2729420a;
}
.plus {
  color: #8e93a4;
  font-size: 25px;
  font-weight: 300;
}
.composer input {
  flex: 1;
  min-width: 0;
  border: 0;
  outline: 0;
  color: #272936;
  background: transparent;
  font-size: 14px;
}
.composer input::placeholder {
  color: #a5a7b3;
}
.composer button {
  padding: 12px 18px;
  border: 0;
  border-radius: 9px;
  color: #20241d;
  background: #c9f27a;
  font-size: 12px;
  font-weight: 700;
}
.composer button:disabled,
.refresh:disabled {
  opacity: 0.6;
  cursor: wait;
}
.toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 29px 0 15px;
  border-bottom: 1px solid #e8e8ef;
}
.tabs {
  display: flex;
  gap: 18px;
}
.tabs button,
.refresh {
  padding: 0 0 12px;
  border: 0;
  background: transparent;
  color: #999baa;
  font-size: 12px;
}
.tabs button span {
  margin-left: 4px;
  color: #b6b8c2;
}
.tabs button.active {
  color: #262833;
  border-bottom: 2px solid #7d9e3f;
  font-weight: 700;
}
.refresh {
  padding-bottom: 12px;
  color: #777b8c;
}
.alert {
  padding: 11px 14px;
  border: 1px solid #efcaca;
  border-radius: 8px;
  color: #b94a4a;
  background: #fff4f4;
  font-size: 13px;
}
.task-list {
  padding: 0;
  margin: 0;
  list-style: none;
}
.task-item {
  display: flex;
  align-items: center;
  gap: 14px;
  min-height: 68px;
  border-bottom: 1px solid #ededf2;
}
.check {
  display: grid;
  place-items: center;
  width: 21px;
  height: 21px;
  padding: 0;
  border: 1.5px solid #c8cbd5;
  border-radius: 7px;
  color: #20251b;
  background: transparent;
  font-size: 13px;
  font-weight: 700;
}
.completed .check {
  border-color: #b2d66d;
  background: #c9f27a;
}
.task-title {
  flex: 1;
  color: #30313d;
  font-size: 14px;
}
.completed .task-title {
  color: #a1a3ae;
  text-decoration: line-through;
}
.task-status {
  padding: 5px 9px;
  border-radius: 20px;
  color: #79838e;
  background: #f0f1f5;
  font-size: 10px;
}
.completed .task-status {
  color: #73963d;
  background: #edf6dc;
}
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 7px;
  padding: 52px 20px 30px;
  color: #999baa;
  text-align: center;
  font-size: 13px;
}
.empty-state strong {
  color: #40424e;
}
.empty-icon {
  color: #a7c965;
  font-size: 28px;
}
.spinner {
  width: 20px;
  height: 20px;
  border: 2px solid #dde4d0;
  border-top-color: #7d9e3f;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}
footer {
  padding-top: 23px;
  color: #62677b;
  text-align: center;
  font-size: 11px;
}
footer span {
  padding: 0 5px;
  color: #9ab75e;
}
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
@media (max-width: 700px) {
  .shell {
    width: min(100% - 24px, 550px);
    padding-top: 24px;
  }
  .hero {
    min-height: 490px;
  }
  .hero-copy {
    margin-top: 57px;
  }
  .progress-card {
    right: 0;
    bottom: 65px;
    width: 100%;
  }
  .workspace {
    padding: 28px 20px 35px;
  }
  .section-heading {
    align-items: start;
    flex-direction: column;
    gap: 15px;
  }
  .composer button {
    padding: 11px 10px;
  }
  .task-status {
    display: none;
  }
}
</style>
