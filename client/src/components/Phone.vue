<script setup lang="ts">
import { useProfileStore } from "@/stores/profile";
const BASE_URL = import.meta.env.VITE_API_BASE_URL;

const profileStore = useProfileStore();
// TODO: Fille the phone with links
</script>

<template>
  <div
    class="flex flex-col bg-phone bg-no-repeat bg-contain w-[307px] h-[631px] px-9 pt-16 pb-14"
  >
    <div class="flex-shrink-0">
      <img
        v-if="profileStore.profile?.pic_path.Valid"
        class="w-24 aspect-square mx-auto rounded-full object-cover border-4 border-brand-300"
        :src="BASE_URL + profileStore.profile?.pic_path.String"
        alt=""
      />
      <div
        v-else
        class="w-24 aspect-square mx-auto bg-neutral-250 rounded-full"
      ></div>

      <span
        class="block text-center mt-6 text-lg font-semibold"
        v-if="
          profileStore.profile?.first_name.Valid &&
          profileStore.profile.last_name.Valid
        "
        >{{
          profileStore.profile?.first_name.String +
          " " +
          profileStore.profile?.last_name.String
        }}
      </span>
      <div
        v-else
        class="w-40 h-4 rounded-full bg-neutral-250 mx-auto mt-6"
      ></div>

      <span
        class="block text-center text-neutral-400 mt-2 text-sm"
        v-if="profileStore.profile?.email"
        >{{ profileStore.profile?.email }}</span
      >
      <div
        v-else
        class="w-16 h-2 rounded-full bg-neutral-250 mx-auto mt-3"
      ></div>
    </div>

    <div class="grid gap-5 mt-auto">
      <div class="w-full h-11 bg-neutral-250 rounded-lg"></div>
      <div class="w-full h-11 bg-neutral-250 rounded-lg"></div>
      <div class="w-full h-11 bg-neutral-250 rounded-lg"></div>
      <div class="w-full h-11 bg-neutral-250 rounded-lg"></div>
      <div class="w-full h-11 bg-neutral-250 rounded-lg"></div>
    </div>
  </div>
</template>
