<script setup lang="ts">
import type { SelectOption } from "@/types";
import { onClickOutside } from "@vueuse/core";
import { computed, ref, type Component, type Ref } from "vue";
import ChevronDown from "@/components/icons/ChevronDown.vue";
import { flip, offset, useFloating } from "@floating-ui/vue";
import { getPlatformProperties } from "@/utils/platform";

const props = defineProps<{
  modelValue: string | null;
  options: SelectOption[];
  title: string;
  icon?: Component;
}>();

const isOpen: Ref<boolean> = ref(false);
const emit = defineEmits(["update:modelValue"]);
const select: Ref<HTMLElement | null> = ref(null);
const optionsRef: Ref<HTMLElement | null> = ref(null);
const platformProperties = computed(() =>
  getPlatformProperties(props.modelValue || ""),
);
const { floatingStyles } = useFloating(select, optionsRef, {
  placement: "bottom-start",
  middleware: [offset(12), flip()],
});

const handleSelect = (newValue: string) => {
  emit("update:modelValue", newValue);
  isOpen.value = false;
};

const toggleIsOpen = () => {
  isOpen.value = !isOpen.value;
};

onClickOutside(select, () => {
  isOpen.value = false;
});
</script>

<template>
  <div class="relative" ref="select">
    <button
      class="flex items-center gap-3 w-full px-4 py-3 bg-neutral-100 outline outline-1 rounded-lg transition-shadow"
      :class="{
        ['shadow-inputShadow outline-brand-300']: isOpen,
        ['outline-neutral-300']: !isOpen,
      }"
      @click="toggleIsOpen"
    >
      <!-- Use selected item's icon if exists -->
      <component
        class="fill-neutral-400"
        v-if="platformProperties.icon"
        :is="platformProperties.icon"
      />

      <!-- Use props' icon if exists and nothing is selected -->
      <component
        class="fill-neutral-400"
        v-else-if="!modelValue && icon"
        :is="icon"
      />

      <span class="flex-1 text-start text-neutral-500">{{
        platformProperties.name || title
      }}</span>
      <component :class="{ ['rotate-180']: isOpen }" :is="ChevronDown" />
    </button>

    <ul
      class="options | bg-neutral-100 px-4 rounded-xl border border-1 border-neutral-300 min-w-full max-h-64 overflow-auto shadow-dropdownShadow z-10"
      v-if="isOpen"
      ref="optionsRef"
      :style="floatingStyles"
    >
      <li
        class="border-b border-1 border-neutral-300 last:border-0"
        v-for="option in options"
        :key="option.value"
      >
        <button
          class="flex items-center gap-3 py-3 w-full"
          :class="{ ['text-brand-300']: modelValue === option.value }"
          @click="handleSelect(option.value)"
        >
          <component
            :class="{
              ['fill-brand-300']: modelValue === option.value,
              ['fill-neutral-400']: modelValue !== option.value,
            }"
            v-if="option.icon"
            :is="option.icon"
          />
          <span>{{ option.label }}</span>
        </button>
      </li>
    </ul>
  </div>
</template>

<style scoped>
.options {
  -ms-overflow-style: none; /* Internet Explorer 10+ */
  scrollbar-width: none; /* Firefox */
}
.options::-webkit-scrollbar {
  display: none; /* Safari and Chrome */
}
</style>
