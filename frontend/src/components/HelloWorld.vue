<script setup lang="ts">
defineProps<{
  msg: string;
}>();
</script>

<template>
  <div class="text-center lg:text-left">
    <h1 class="text-3xl font-bold text-base-content">{{ msg }}</h1>
    <h3 class="mt-2 text-lg">
      You've successfully created a project with
      <a href="https://vite.dev/" target="_blank" rel="noopener" class="link link-primary">Vite</a>
      +
      <a href="https://vuejs.org/" target="_blank" rel="noopener" class="link link-primary">Vue 3</a
      >. What's next?
    </h3>
  </div>
</template>
