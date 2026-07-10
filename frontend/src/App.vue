<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
  <nav class="navbar bg-base-100 sticky top-0 z-10 border-b border-base-200">
    <div class="navbar-start">
      <RouterLink to="/menu" class="btn btn-ghost text-xl font-semibold tracking-tight">
        <img
          alt="The Good Café logo"
          src="@/assets/logo.svg"
          width="28"
          height="28"
          class="text-primary"
        />
        The Good Café
      </RouterLink>
    </div>
    <div class="navbar-end gap-1">
      <RouterLink to="/menu" class="btn btn-ghost">Menu</RouterLink>
      <RouterLink to="/chat" class="btn btn-ghost">Chat</RouterLink>
    </div>
  </nav>

  <main>
    <RouterView />
  </main>
</template>
