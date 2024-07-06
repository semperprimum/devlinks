<script setup lang="ts">
import type { Component } from "vue";

defineProps<{
  secondary?: boolean;
  danger?: boolean;
  leading?: Component;
}>();
</script>

<template>
  <button
    class="font-semibold py-3 px-6 rounded-lg disabled:opacity-25 disabled:cursor-not-allowed"
    :class="{
      'bg-brand-300 text-neutral-100': !secondary && !danger,
      'bg-transparent outline outline-1 outline-brand-300 text-brand-300':
        secondary,
      'bg-red text-neutral-100': danger,
    }"
  >
    <slot></slot>
  </button>
</template>
