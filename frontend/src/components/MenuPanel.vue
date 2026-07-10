<script setup lang="ts">
import { ref, computed, onMounted } from "vue";

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
const activeCategory = ref("All");
const showForm = ref(false);
const form = ref({ name: "", description: "", category: "Drinks", price: 0 });
const saving = ref(false);

const categories = computed(() => {
  const cats = new Set(items.value.map((i) => i.category));
  return ["All", ...Array.from(cats)];
});

const filtered = computed(() =>
  activeCategory.value === "All"
    ? items.value
    : items.value.filter((i) => i.category === activeCategory.value),
);

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
  <div class="max-w-4xl mx-auto px-4 py-16">
    <!-- Header -->
    <div class="flex flex-col sm:flex-row sm:items-end justify-between gap-4 mb-10">
      <div>
        <p class="text-xs font-medium uppercase tracking-[0.18em] text-primary/70 mb-3">Menu</p>
        <h1 class="text-3xl md:text-5xl font-bold tracking-tight">What we're serving</h1>
      </div>
      <button class="btn btn-primary" @click="showForm = !showForm">
        {{ showForm ? "Cancel" : "Add Item" }}
      </button>
    </div>

    <!-- Category tabs -->
    <div v-if="!loading && !error" class="flex gap-2 flex-wrap mb-8">
      <button
        v-for="cat in categories"
        :key="cat"
        class="btn btn-sm transition-all duration-200"
        :class="activeCategory === cat ? 'btn-primary' : 'btn-ghost'"
        @click="activeCategory = cat"
      >
        {{ cat }}
      </button>
    </div>

    <!-- Add form -->
    <div v-if="showForm" class="card bg-base-200 border border-base-300 mb-8">
      <div class="card-body gap-3">
        <h3 class="font-semibold text-lg">New Menu Item</h3>
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
          <span v-if="saving" class="loading loading-spinner loading-xs"></span>
          {{ saving ? "Adding..." : "Add to Menu" }}
        </button>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center py-20">
      <span class="loading loading-dots loading-lg text-primary"></span>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="card bg-error/10 border border-error/20">
      <div class="card-body text-center py-12">
        <p class="text-error font-medium">{{ error }}</p>
        <button class="btn btn-ghost btn-sm mt-2" @click="fetchMenu">Retry</button>
      </div>
    </div>

    <!-- Menu grid -->
    <div v-else class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
      <div
        v-for="item in filtered"
        :key="item.id"
        class="card bg-base-100 border border-base-200 shadow-sm hover:shadow-md hover:-translate-y-0.5 transition-all duration-200"
      >
        <div class="card-body gap-2 p-5">
          <div class="flex items-start justify-between gap-2">
            <h3 class="font-semibold text-base-content">{{ item.name }}</h3>
            <span class="badge badge-ghost badge-sm shrink-0">{{ item.category }}</span>
          </div>
          <p class="text-sm text-base-content/50 leading-relaxed">{{ item.description }}</p>
          <p class="text-lg font-semibold text-primary tabular-nums mt-1">
            ${{ item.price.toFixed(2) }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
