import axios from "axios";
import { defineStore } from "pinia";
import { computed, ref, type ComputedRef, type Ref } from "vue";
import { useAuthStore } from "@/stores/auth";
import type { GetUserLinksResponse, UserInfo } from "@/types";

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const useProfileStore = defineStore("profile", () => {
  const authStore = useAuthStore();
  const headers: ComputedRef<{ Authorization: string }> = computed(() => {
    return { Authorization: `Bearer ${authStore.token}` };
  });
  const profile: Ref<UserInfo | null> = ref(null);
  const links: Ref<GetUserLinksResponse | null> = ref(null);

  const getUserInfo = async () => {
    try {
      const response = await axios.get(`${BASE_URL}/user`, {
        headers: headers.value,
      });

      profile.value = response.data;
    } catch (e: any) {
      console.error(e.response?.data?.message || e.message);
    }
  };

  const getUserLinks = async () => {
    try {
      const response = await axios.get(
        `${BASE_URL}/links/${profile.value?.id}`,
        { headers: headers.value },
      );

      links.value = response.data;
    } catch (e: any) {
      console.error(e.response?.data?.message || e.message);
    }
  };

  return { getUserInfo, getUserLinks, profile, links };
});
