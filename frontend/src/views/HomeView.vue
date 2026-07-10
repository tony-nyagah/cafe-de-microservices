<script setup lang="ts">
import TheWelcome from "../components/TheWelcome.vue";
</script>

<template>
  <div class="hero bg-base-200 rounded-box mb-8">
    <div class="hero-content text-center py-12">
      <div>
        <h1 class="text-4xl font-bold text-base-content">Welcome to The Good Café</h1>
        <p class="mt-4 text-lg text-base-content/70">
          Freshly brewed coffee, served one cup at a time.
        </p>
      </div>
    </div>
  </div>

  <TheWelcome />
</template>
