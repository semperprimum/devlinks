<script setup lang="ts">
import Header from "@/components/Header.vue";
import Phone from "@/components/Phone.vue";
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <Header />

    <main class="flex-grow p-8 flex items-stretch gap-6 md:pt-0 min-h-full">
      <div
        class="flex-shrink-0 grid place-items-center w-[35rem] bg-neutral-100 rounded-3xl"
      >
        <Phone />
      </div>
      <div class="bg-neutral-100 rounded-3xl flex-1"></div>
    </main>
  </div>
</template>
