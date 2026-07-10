<script setup lang="ts">
import { ref } from "vue";

const messages = ref<{ role: string; content: string }[]>([]);
const input = ref("");
const loading = ref(false);

async function send() {
  const text = input.value.trim();
  if (!text) return;
  messages.value.push({ role: "user", content: text });
  input.value = "";
  loading.value = true;

  try {
    const res = await fetch("/api/chat", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ message: text }),
    });
    const data = await res.json();
    messages.value.push({ role: "assistant", content: data.reply ?? "No response." });
  } catch {
    messages.value.push({ role: "assistant", content: "Could not reach the AI service." });
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <div class="max-w-2xl mx-auto px-4 py-16">
    <p class="text-sm font-medium uppercase tracking-widest text-primary/70 mb-4">Chat</p>
    <h1 class="text-3xl md:text-5xl font-bold tracking-tight mb-10">Ask the barista</h1>

    <div class="card bg-base-200 shadow-sm">
      <div class="card-body gap-4 min-h-[400px] max-h-[60dvh] overflow-y-auto">
        <div v-if="messages.length === 0" class="text-center text-base-content/40 py-12">
          Ask about our menu or get a recommendation.
        </div>
        <div
          v-for="(msg, i) in messages"
          :key="i"
          class="chat"
          :class="msg.role === 'user' ? 'chat-end' : 'chat-start'"
        >
          <div
            class="chat-bubble"
            :class="msg.role === 'user' ? 'chat-bubble-primary' : 'chat-bubble'"
          >
            {{ msg.content }}
          </div>
        </div>
        <div v-if="loading" class="flex items-center gap-2 text-base-content/50 text-sm px-4">
          <span class="loading loading-dots loading-xs"></span>
          Thinking...
        </div>
      </div>
      <div class="card-actions p-4 pt-0">
        <form class="flex gap-2 w-full" @submit.prevent="send">
          <input
            v-model="input"
            type="text"
            placeholder="Ask about the menu..."
            class="input input-bordered flex-1"
            :disabled="loading"
          />
          <button type="submit" class="btn btn-primary" :disabled="loading || !input.trim()">
            Send
          </button>
        </form>
      </div>
    </div>
  </div>
</template>
