<script setup lang="ts">
import { RouterLink, useRoute } from "vue-router";
import LogoSmall from "./icons/LogoSmall.vue";
import LogoFull from "./icons/LogoFull.vue";
import HeaderButton from "@/components/HeaderButton.vue";
import Link from "@/components/icons/Link.vue";
import Eye from "@/components/icons/Eye.vue";
import User from "@/components/icons/User.vue";

const route = useRoute();
</script>

<template>
  <div class="hidden md:pt-6 md:block"></div>
  <header
    class="flex justify-between items-center bg-neutral-100 p-4 pl-6 md:m-6 md:mt-0 md:rounded-2xl"
  >
    <div>
      <LogoSmall class="block md:hidden" />
      <LogoFull class="hidden md:block max-h-8 max-w-36" />
    </div>
    <div class="flex">
      <HeaderButton
        :active="route.name === 'home'"
        to="/"
        label="Links"
        :icon="Link"
      />
      <HeaderButton
        :active="route.name === 'profile'"
        to="/profile"
        label="Profile Details"
        :icon="User"
      />
    </div>
    <div class="justify-self-end">
      <RouterLink
        to="/"
        class="block px-4 py-3 outline outline-1 outline-brand-300 rounded-lg md:px-7"
      >
        <span class="hidden md:block text-brand-300 font-semibold"
          >Preview</span
        >
        <Eye class="fill-brand-300 block md:hidden" />
      </RouterLink>
    </div>
  </header>
</template>
