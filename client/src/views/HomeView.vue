<script setup lang="ts">
import Header from "@/components/Header.vue";
import Phone from "@/components/Phone.vue";
import { useAuthStore } from "@/stores/auth";
import { useProfileStore } from "@/stores/profile";
import { onMounted } from "vue";

const authStore = useAuthStore();
const profileStore = useProfileStore();

onMounted(async () => {
  if (authStore.token) {
    await profileStore.getUserInfo();
    await profileStore.getUserLinks();
  }
});
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <Header />

    <main class="flex-grow p-8 flex items-stretch gap-6 md:pt-0 min-h-full">
      <div
        class="hidden xl:grid py-4 flex-shrink-0 place-items-center w-[35rem] bg-neutral-100 rounded-3xl"
      >
        <Phone />
      </div>
      <div class="bg-neutral-100 rounded-3xl flex-1"></div>
    </main>
  </div>
</template>
