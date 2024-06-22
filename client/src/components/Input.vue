<script setup lang="ts">
import type { Component } from "vue";

defineProps<{
  placeholder?: string;
  leading?: Component;
}>();
</script>

<template>
  <div
    class="flex items-center gap-3 px-4 bg-neutral-100 outline outline-1 outline-neutral-300 rounded-lg focus-within:outline focus-within:outline-1 focus-within:outline-brand-300 focus-within:shadow-inputShadow"
  >
    <component v-if="leading" :is="leading" />
    <input
      class="w-full h-full py-3 outline-none"
      type="text"
      v-bind="{ placeholder }"
    />
  </div>
</template>
