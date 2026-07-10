import { defineStore } from "pinia";
import { ref } from "vue";

// Define the shape of a menu item from the Go service
export interface MenuItem {
  ID: number;
  Name: string;
  Description: string;
  Category: string;
  Price: number;
  Available: boolean;
}

export const useMenuStore = defineStore("menu", () => {
  // Reactive list of all menu items
  const items = ref<MenuItem[]>([]);
  const loading = ref(false);
  const error = ref<string | null>(null);

  // Fetch all menu items from the Go Menu Service
  async function fetchItems() {
    loading.value = true;
    error.value = null;
    try {
      const response = await fetch("/api/menu");
      if (!response.ok) throw new Error("Failed to fetch menu items");
      items.value = await response.json();
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  }

  // Create a new menu item
  async function createItem(item: Omit<MenuItem, "ID">) {
    const response = await fetch("/api/menu", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(item),
    });
    if (!response.ok) throw new Error("Failed to create item");
    await fetchItems();
  }

  // Delete a menu item by ID
  async function deleteItem(id: number) {
    const response = await fetch(`/api/menu/${id}`, { method: "DELETE" });
    if (!response.ok) throw new Error("Failed to delete item");
    await fetchItems();
  }

  return { items, loading, error, fetchItems, createItem, deleteItem };
});
