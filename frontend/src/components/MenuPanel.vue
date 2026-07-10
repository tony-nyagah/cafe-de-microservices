<script setup lang="ts">
import { ref, onMounted } from "vue";

interface MenuItem {
  id: number;
  name: string;
  description: string;
  category: string;
  price: number;
}

const items = ref<MenuItem[]>([]);
const loading = ref(true);
const error = ref("");
const showForm = ref(false);
const form = ref({ name: "", description: "", category: "Drinks", price: 0 });
const saving = ref(false);

async function fetchMenu() {
  loading.value = true;
  error.value = "";
  try {
    const res = await fetch("/api/menu");
    if (!res.ok) throw new Error("Failed to fetch menu");
    items.value = await res.json();
  } catch {
    error.value = "Could not load the menu.";
  } finally {
    loading.value = false;
  }
}

async function addItem() {
  saving.value = true;
  try {
    const res = await fetch("/api/menu", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(form.value),
    });
    if (!res.ok) throw new Error("Failed to add item");
    await fetchMenu();
    form.value = { name: "", description: "", category: "Drinks", price: 0 };
    showForm.value = false;
  } catch {
    error.value = "Could not add item.";
  } finally {
    saving.value = false;
  }
}

onMounted(fetchMenu);
</script>

<template>
  <div class="max-w-3xl mx-auto px-4 py-16">
    <div class="flex items-end justify-between mb-10">
      <div>
        <p class="text-sm font-medium uppercase tracking-widest text-primary/70 mb-4">Menu</p>
        <h1 class="text-3xl md:text-5xl font-bold tracking-tight">What we're serving</h1>
      </div>
      <button class="btn btn-primary btn-sm" @click="showForm = !showForm">
        {{ showForm ? "Cancel" : "Add Item" }}
      </button>
    </div>

    <!-- Add form -->
    <div v-if="showForm" class="card bg-base-200 shadow-sm mb-8">
      <div class="card-body gap-3">
        <h3 class="font-semibold">New Menu Item</h3>
        <div class="grid gap-3 sm:grid-cols-2">
          <input v-model="form.name" placeholder="Name" class="input input-bordered" />
          <select v-model="form.category" class="select select-bordered">
            <option>Drinks</option>
            <option>Pastries</option>
            <option>Food</option>
          </select>
          <input
            v-model.number="form.price"
            type="number"
            step="0.01"
            min="0"
            placeholder="Price"
            class="input input-bordered"
          />
          <input
            v-model="form.description"
            placeholder="Description"
            class="input input-bordered sm:col-span-2"
          />
        </div>
        <button
          class="btn btn-primary"
          :disabled="saving || !form.name || form.price <= 0"
          @click="addItem"
        >
          {{ saving ? "Adding..." : "Add to Menu" }}
        </button>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="text-center text-base-content/40 py-12">
      <span class="loading loading-dots loading-lg"></span>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="text-center text-error py-12">{{ error }}</div>

    <!-- Empty -->
    <div v-else-if="items.length === 0" class="text-center text-base-content/40 py-12">
      Nothing on the menu yet. Add something!
    </div>

    <!-- Menu list -->
    <div v-else class="space-y-1">
      <div
        v-for="item in items"
        :key="item.id"
        class="flex items-center justify-between py-4 border-b border-base-200 last:border-b-0"
      >
        <div>
          <h3 class="font-semibold text-base-content">
            {{ item.name }}
            <span class="badge badge-ghost badge-sm align-middle ml-1">{{ item.category }}</span>
          </h3>
          <p class="text-sm text-base-content/50">{{ item.description }}</p>
        </div>
        <span class="text-base-content/70 font-medium tabular-nums"
          >${{ item.price.toFixed(2) }}</span
        >
      </div>
    </div>
  </div>
</template>
