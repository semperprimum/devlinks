<script setup lang="ts">
import { RouterLink } from "vue-router";
import ArrowRight from "@/components/icons/ArrowRight.vue";
</script>

<template>
  <section class="grid min-h-screen place-items-center">
    <div>
      <h1 class="text-6xl text-center">
        404
        <span class="block text-3xl mt-1">not found</span>
      </h1>

      <RouterLink
        to="/"
        class="text-lg text-brand-300 flex gap-2 items-center mt-6"
        >To Homepage <ArrowRight class="fill-brand-300"
      /></RouterLink>
    </div>
  </section>
</template>
