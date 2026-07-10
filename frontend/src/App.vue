<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
  <!-- DaisyUI navbar for switching between pages -->
  <div class="navbar bg-base-100 shadow-sm">
    <div class="flex-1">
      <a class="btn btn-ghost text-xl">Cafe Manager</a>
    </div>
    <div class="flex-none">
      <ul class="menu menu-horizontal px-1">
        <li><RouterLink to="/menu">Menu</RouterLink></li>
        <li><RouterLink to="/chat">Chat</RouterLink></li>
      </ul>
    </div>
  </div>

  <!-- Route content renders here -->
  <div class="container mx-auto" style="height: calc(100vh - 64px)">
    <RouterView />
  </div>
</template>
