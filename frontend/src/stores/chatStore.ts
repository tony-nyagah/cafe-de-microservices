import { defineStore } from "pinia";
import { ref } from "vue";

// Each message has a role (user or assistant) and text content
export interface ChatMessage {
  role: "user" | "assistant";
  content: string;
}

export const useChatStore = defineStore("chat", () => {
  const messages = ref<ChatMessage[]>([]);
  const loading = ref(false);

  // Send a message to the AI Service and get a response
  async function sendMessage(userMessage: string) {
    // Add the user's message to the chat history
    messages.value.push({ role: "user", content: userMessage });
    loading.value = true;

    try {
      const response = await fetch("/api/chat", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ message: userMessage }),
      });
      if (!response.ok) throw new Error("AI Service unavailable");
      const data = await response.json();
      // Add the AI's response to the chat history
      messages.value.push({ role: "assistant", content: data.response });
    } catch {
      messages.value.push({
        role: "assistant",
        content: "Sorry, I could not reach the AI service. Please try again.",
      });
    } finally {
      loading.value = false;
    }
  }

  return { messages, loading, sendMessage };
});
