<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useMenuStore } from "../stores/menuStore";

const store = useMenuStore();

// Form fields for creating a new menu item
const name = ref("");
const description = ref("");
const category = ref("Coffee");
const price = ref(0);

onMounted(() => {
  // Load menu items when the component first renders
  store.fetchItems();
});

async function handleCreate() {
  await store.createItem({
    Name: name.value,
    Description: description.value,
    Category: category.value,
    Price: price.value,
    Available: true,
  });
  // Reset form fields after creating
  name.value = "";
  description.value = "";
  price.value = 0;
}
</script>

<template>
  <div class="p-6">
    <h2 class="text-2xl font-bold mb-4">Menu Management</h2>

    <!-- New item form using DaisyUI form controls -->
    <div class="card bg-base-200 p-4 mb-6">
      <h3 class="text-lg font-semibold mb-2">Add New Item</h3>
      <div class="flex flex-wrap gap-2">
        <input
          v-model="name"
          type="text"
          placeholder="Item name"
          class="input input-bordered w-48"
        />
        <input
          v-model="description"
          type="text"
          placeholder="Description"
          class="input input-bordered w-48"
        />
        <select v-model="category" class="select select-bordered">
          <option>Coffee</option>
          <option>Tea</option>
          <option>Pastry</option>
          <option>Sandwich</option>
        </select>
        <input
          v-model.number="price"
          type="number"
          step="0.01"
          placeholder="Price"
          class="input input-bordered w-24"
        />
        <button class="btn btn-primary" @click="handleCreate">Add</button>
      </div>
    </div>

    <!-- Loading spinner -->
    <div v-if="store.loading" class="flex justify-center p-8">
      <span class="loading loading-spinner loading-lg"></span>
    </div>

    <!-- Menu items table using DaisyUI table classes -->
    <div v-else class="overflow-x-auto">
      <table class="table">
        <thead>
          <tr>
            <th>Name</th>
            <th>Category</th>
            <th>Price</th>
            <th>Available</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in store.items" :key="item.ID">
            <td>
              <div class="font-bold">{{ item.Name }}</div>
              <div class="text-sm opacity-50">{{ item.Description }}</div>
            </td>
            <td>
              <span class="badge badge-outline">{{ item.Category }}</span>
            </td>
            <td>${{ item.Price.toFixed(2) }}</td>
            <td>
              <span :class="item.Available ? 'badge badge-success' : 'badge badge-error'">
                {{ item.Available ? "Yes" : "No" }}
              </span>
            </td>
            <td>
              <button class="btn btn-error btn-sm" @click="store.deleteItem(item.ID)">
                Delete
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
