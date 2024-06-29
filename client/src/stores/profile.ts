import axios from "axios";
import { defineStore } from "pinia";
import {
  computed,
  reactive,
  ref,
  type ComputedRef,
  type Reactive,
  type Ref,
} from "vue";
import { useAuthStore } from "@/stores/auth";
import type {
  Buffer,
  GetUserLinksResponse,
  SelectOption,
  UserInfo,
  UserLink,
} from "@/types";

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const useProfileStore = defineStore("profile", () => {
  const authStore = useAuthStore();
  const headers: ComputedRef<{ Authorization: string }> = computed(() => {
    return { Authorization: `Bearer ${authStore.token}` };
  });
  const profile: Ref<UserInfo | null> = ref(null);
  const links: Ref<GetUserLinksResponse | null> = ref(null);
  const buffer: Reactive<Buffer> = reactive({
    added: [],
    deleted: [],
    updated: [],
  });
  const changesMade: Ref<boolean> = ref(false);

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

  const updatePlatform = (id: number | string, platform: string) => {
    if (!links.value?.links) return;
    let linkIndex: number | null = null;

    if (typeof id === "number") {
      linkIndex = links.value.links.findIndex((link) => link.id === id);
      if (!buffer.updated.includes(id)) {
        buffer.updated.push(id);
      }
    } else {
      linkIndex = links.value.links.findIndex((link) => link.temp_id === id);
    }

    if (linkIndex !== -1) {
      links.value.links[linkIndex].platform = platform;
      changesMade.value = true;
    }
  };

  const addLink = () => {
    const newLink: UserLink = {
      temp_id: crypto.randomUUID(),
      platform: "github",
      url: "",
    };

    if (!newLink.temp_id) return;

    links.value?.links.push(newLink);
    buffer.added.push(newLink.temp_id);
    changesMade.value = true;
  };

  const updateUrl = (id: number | string, url: string) => {
    if (!links.value?.links) return;
    let linkIndex: number | null = null;

    if (typeof id === "number") {
      linkIndex = links.value.links.findIndex((link) => link.id === id);
      if (!buffer.updated.includes(id)) {
        buffer.updated.push(id);
      }
    } else {
      linkIndex = links.value.links.findIndex((link) => link.temp_id === id);
    }

    if (linkIndex !== -1) {
      links.value.links[linkIndex].url = url;
      changesMade.value = true;
    }
  };

  const removeLink = (id: number | string) => {
    if (!links.value?.links) return;

    if (typeof id === "number") {
      links.value.links = links.value.links.filter((link) => link.id !== id);
      buffer.updated = buffer.updated.filter((linkId) => linkId !== id);
      buffer.deleted.push(id);
      changesMade.value = true;
    } else {
      links.value.links = links.value.links.filter(
        (link) => link.temp_id !== id,
      );
      buffer.added = buffer.added.filter((temp_id) => temp_id !== id);
      changesMade.value = true;
    }
  };

  const $reset = () => {
    profile.value = null;
    links.value = null;
    buffer.updated = [];
    buffer.deleted = [];
    buffer.added = [];
    changesMade.value = false;
  };

  return {
    getUserInfo,
    getUserLinks,
    updatePlatform,
    updateUrl,
    removeLink,
    addLink,
    buffer,
    changesMade,
    profile,
    links,
    $reset,
  };
});
