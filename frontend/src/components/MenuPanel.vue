<script setup lang="ts">
const items = [
  { name: "Espresso", description: "A bold, concentrated shot of pure coffee.", price: 3.5 },
  {
    name: "Cappuccino",
    description: "Espresso with steamed milk and a crown of foam.",
    price: 4.5,
  },
  { name: "Latte", description: "Smooth espresso with velvety steamed milk.", price: 4.75 },
  { name: "Mocha", description: "Rich chocolate meets espresso and steamed milk.", price: 5.25 },
  {
    name: "Cold Brew",
    description: "Slow-steeped for 18 hours. Smooth and naturally sweet.",
    price: 4.25,
  },
  {
    name: "Matcha Latte",
    description: "Ceremonial-grade matcha whisked with steamed milk.",
    price: 5.5,
  },
  {
    name: "Chai Latte",
    description: "Spiced black tea concentrate with steamed milk.",
    price: 4.75,
  },
  { name: "Pour Over", description: "Single-origin beans, brewed to order.", price: 5.0 },
  { name: "Croissant", description: "Buttery, flaky, and baked fresh daily.", price: 3.75 },
];
</script>

<template>
  <div class="max-w-3xl mx-auto px-4 py-16">
    <p class="text-sm font-medium uppercase tracking-widest text-primary/70 mb-4">Menu</p>
    <h1 class="text-3xl md:text-5xl font-bold tracking-tight mb-10">What we're serving</h1>

    <div class="space-y-1">
      <div
        v-for="item in items"
        :key="item.name"
        class="flex items-center justify-between py-4 border-b border-base-200 last:border-b-0"
      >
        <div>
          <h3 class="font-semibold text-base-content">{{ item.name }}</h3>
          <p class="text-sm text-base-content/50">{{ item.description }}</p>
        </div>
        <span class="text-base-content/70 font-medium tabular-nums"
          >${{ item.price.toFixed(2) }}</span
        >
      </div>
    </div>
  </div>
</template>
