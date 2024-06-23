<script setup lang="ts">
import type { Component, DefineComponent } from "vue";
import { RouterLink } from "vue-router";

defineProps<{
  to: string;
  icon: Component;
  label: string;
  active?: boolean;
}>();
</script>

<template>
  <RouterLink
    class="flex items-center justify-center gap-2 py-3 px-7 rounded-lg"
    :class="{
      ['bg-brand-100 fill-brand-300']: active,
      ['fill-neutral-400']: !active,
    }"
    :to="to"
  >
    <component class="block" :is="icon" />
    <span
      class="hidden md:block font-semibold"
      :class="{
        ['text-brand-300']: active,
        ['text-neutral-400']: !active,
      }"
      >{{ label }}</span
    >
  </RouterLink>
</template>
