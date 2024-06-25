<script setup lang="ts">
import { RouterView, useRouter } from "vue-router";
import { useAuthStore } from "./stores/auth";
import { onMounted } from "vue";

const authStore = useAuthStore();
const router = useRouter();

onMounted(async () => {
  const isAuthValid = await authStore.checkAuth();
  if (!isAuthValid) {
    router.replace("/login");
  }
});
</script>

<template>
  <RouterView />
</template>
