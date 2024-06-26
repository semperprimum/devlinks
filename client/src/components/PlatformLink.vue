<script setup lang="ts">
import { getPlatformProperties } from "@/utils/platform";
import ArrowRight from "@/components/icons/ArrowRight.vue";

const props = defineProps<{
  platform: string;
  link: string;
}>();

const platformProperties = getPlatformProperties(props.platform);
</script>

<template>
  <a
    :href="link"
    :style="{ '--clr': platformProperties.color }"
    class="w-full h-11 rounded-lg text-neutral-100 py-3 px-4 flex items-center gap-2"
    :class="{
      'bg-[--clr]': true,
      ['text-neutral-500']: platform === 'frontendmentor',
      ['text-neutral-100']: platform !== 'frontendmentor',
    }"
  >
    <component
      class="max-w-4 aspect-square"
      :class="{
        ['fill-neutral-500']: platform === 'frontendmentor',
        ['fill-neutral-100']: platform !== 'frontendmentor',
      }"
      :is="platformProperties.icon"
    />
    <span class="flex-1">{{ platformProperties.name }}</span>
    <component
      :class="{
        ['fill-neutral-500']: platform === 'frontendmentor',
        ['fill-neutral-100']: platform !== 'frontendmentor',
      }"
      :is="ArrowRight"
    />
  </a>
</template>
