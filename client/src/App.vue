<script setup lang="ts">
import { RouterView, useRoute, useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { onMounted } from "vue";
import { useProfileStore } from "@/stores/profile";

const authStore = useAuthStore();
const profileStore = useProfileStore();
const router = useRouter();
const route = useRoute();

onMounted(async () => {
  const isAuthValid = await authStore.checkAuth();

  await router.isReady();

  if (!isAuthValid && route.name !== "link") {
    router.replace("/login");
  }

  if (authStore.token) {
    await profileStore.getUserInfo();
    await profileStore.getUserLinks();
  }
});
</script>

<template>
  <RouterView />
</template>
