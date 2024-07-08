<script setup lang="ts">
import { RouterView, useRoute, useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { onMounted } from "vue";
import { useProfileStore } from "@/stores/profile";
import Toast from "@/components/Toast.vue";
import { useToast } from "./services/ToastService";

const authStore = useAuthStore();
const profileStore = useProfileStore();
const router = useRouter();
const route = useRoute();
const { toastMessage, toastType } = useToast();

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
  <Toast :message="toastMessage" :type="toastType" />
</template>
