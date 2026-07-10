<script setup lang="ts">
import { ref, nextTick } from "vue";

const messages = ref<{ role: string; content: string }[]>([]);
const input = ref("");
const loading = ref(false);
const chatBody = ref<HTMLElement | null>(null);

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
    await nextTick();
    chatBody.value?.scrollTo({ top: chatBody.value.scrollHeight, behavior: "smooth" });
  }
}
</script>

<template>
  <div class="max-w-2xl mx-auto px-4 py-16">
    <!-- Header -->
    <div class="mb-10">
      <p class="text-xs font-medium uppercase tracking-[0.18em] text-primary/70 mb-3">Chat</p>
      <h1 class="text-3xl md:text-5xl font-bold tracking-tight">Ask the barista</h1>
      <p class="mt-2 text-base-content/50">
        Get recommendations, ask about the menu, or just say hi.
      </p>
    </div>

    <!-- Chat card -->
    <div class="card bg-base-100 border border-base-200 shadow-sm">
      <div ref="chatBody" class="card-body gap-4 min-h-[440px] max-h-[60dvh] overflow-y-auto">
        <!-- Empty state -->
        <div
          v-if="messages.length === 0 && !loading"
          class="flex flex-col items-center justify-center text-center flex-1 py-12"
        >
          <div
            class="inline-flex h-16 w-16 items-center justify-center rounded-full bg-primary/10 text-primary mb-4"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="28"
              height="28"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="1.5"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z" />
            </svg>
          </div>
          <p class="text-base-content/40 text-sm">Ask about our menu or get a recommendation.</p>
        </div>

        <!-- Messages -->
        <div
          v-for="(msg, i) in messages"
          :key="i"
          class="chat"
          :class="msg.role === 'user' ? 'chat-end' : 'chat-start'"
        >
          <div
            class="chat-bubble text-sm leading-relaxed"
            :class="
              msg.role === 'user'
                ? 'chat-bubble-primary'
                : 'chat-bubble bg-base-200 text-base-content'
            "
          >
            {{ msg.content }}
          </div>
        </div>

        <!-- Typing indicator -->
        <div v-if="loading" class="chat chat-start">
          <div class="chat-bubble bg-base-200 flex items-center gap-1 px-4">
            <span class="loading loading-dots loading-xs"></span>
          </div>
        </div>
      </div>

      <!-- Input -->
      <div class="p-4 pt-0">
        <form class="flex gap-2" @submit.prevent="send">
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
