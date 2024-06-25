<script setup lang="ts">
import Header from "@/components/Header.vue";
import { useAuthStore } from "@/stores/auth";
import { useRouter } from "vue-router";
import Button from "@/components/Button.vue";

const authStore = useAuthStore();
const router = useRouter();

const handleLogout = () => {
  authStore.logout();

  router.replace("/login");
};
</script>

<template>
  <Header />

  <main class="p-8 md:pt-0">
    <h1 class="text-4xl text-brand-300">Profile</h1>

    <Button @click="handleLogout" secondary>Logout</Button>
  </main>
</template>
