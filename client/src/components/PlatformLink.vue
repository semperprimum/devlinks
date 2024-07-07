<script setup lang="ts">
import { getPlatformProperties } from "@/utils/platform";
import ArrowRight from "@/components/icons/ArrowRight.vue";
import { computed } from "vue";

const props = defineProps<{
  platform: string;
  link: string;
  large?: boolean;
}>();

const platformProperties = computed(() =>
  getPlatformProperties(props.platform),
);
</script>

<template>
  <a
    :href="link"
    :style="{ '--clr': platformProperties.color }"
    class="w-full rounded-lg h-11 text-neutral-100 py-3 px-4 flex items-center gap-2 bg-[--clr]"
    :class="{
      'text-neutral-500': platform === 'frontendmentor',
      'py-4 h-auto': large,
    }"
  >
    <component
      class="w-4 h-4 fill-neutral-100"
      :class="{
        'fill-neutral-500': platform === 'frontendmentor',
        'w-5 h-5': large,
      }"
      :is="platformProperties.icon"
    />
    <span
      class="flex-1"
      :class="{
        'text-base': large,
        'text-xs': !large,
      }"
      >{{ platformProperties.name }}</span
    >
    <component
      class="fill-neutral-100"
      :class="{
        'fill-neutral-500': platform === 'frontendmentor',
      }"
      :is="ArrowRight"
    />
  </a>
</template>
