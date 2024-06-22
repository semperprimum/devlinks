<script setup lang="ts">
defineProps<{
  text: string;
}>();
</script>

<template>
  <span class="text-xs">
    {{ text }}
  </span>
</template>
