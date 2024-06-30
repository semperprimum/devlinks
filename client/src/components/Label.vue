<script setup lang="ts">
defineProps<{
  text: string;
  error?: string;
}>();
</script>

<template>
  <span class="text-xs" :class="{ ['text-red']: error }">
    {{ text }}
  </span>
</template>
