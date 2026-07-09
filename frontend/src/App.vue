<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
  <div class="navbar bg-base-200 shadow-xs">
    <div class="navbar-start">
      <RouterLink to="/" class="btn btn-ghost text-xl">
        <img alt="logo" src="@/assets/logo.svg" width="28" height="28" class="mr-2" />
        Café de Microservices
      </RouterLink>
    </div>
    <div class="navbar-end">
      <RouterLink to="/" class="btn btn-ghost">Home</RouterLink>
      <RouterLink to="/about" class="btn btn-ghost">About</RouterLink>
    </div>
  </div>

  <main class="mx-auto max-w-5xl px-4 py-8">
    <RouterView />
  </main>
</template>
