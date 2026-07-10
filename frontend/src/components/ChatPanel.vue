<script setup lang="ts">
import { ref } from "vue";
import { useChatStore } from "../stores/chatStore";

const store = useChatStore();
const input = ref("");

async function handleSend() {
  if (!input.value.trim()) return;
  const message = input.value;
  input.value = "";
  await store.sendMessage(message);
}
</script>

<template>
  <div class="flex flex-col h-full p-6">
    <h2 class="text-2xl font-bold mb-4">Cafe Assistant</h2>

    <!-- Scrollable message history -->
    <div class="flex-1 overflow-y-auto space-y-2 mb-4">
      <div v-for="(msg, i) in store.messages" :key="i">
        <!-- User messages appear on the right -->
        <div v-if="msg.role === 'user'" class="chat chat-end">
          <div class="chat-bubble chat-bubble-primary">{{ msg.content }}</div>
        </div>
        <!-- AI messages appear on the left -->
        <div v-else class="chat chat-start">
          <div class="chat-bubble">{{ msg.content }}</div>
        </div>
      </div>

      <!-- Loading indicator while waiting for AI response -->
      <div v-if="store.loading" class="chat chat-start">
        <div class="chat-bubble">
          <span class="loading loading-dots loading-sm"></span>
        </div>
      </div>
    </div>

    <!-- Message input area -->
    <div class="flex gap-2">
      <input
        v-model="input"
        type="text"
        placeholder="Ask about our menu..."
        class="input input-bordered flex-1"
        @keyup.enter="handleSend"
      />
      <button class="btn btn-primary" :disabled="store.loading" @click="handleSend">Send</button>
    </div>
  </div>
</template>
