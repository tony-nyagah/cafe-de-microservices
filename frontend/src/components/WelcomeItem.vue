<template>
  <div class="card bg-base-100 card-border">
    <div class="card-body gap-4">
      <i
        class="inline-flex h-12 w-12 items-center justify-center rounded-btn bg-base-300 text-base-content"
      >
        <slot name="icon"></slot>
      </i>
      <div>
        <h3 class="card-title">
          <slot name="heading"></slot>
        </h3>
        <p class="text-base-content/70">
          <slot></slot>
        </p>
      </div>
    </div>
  </div>
</template>
