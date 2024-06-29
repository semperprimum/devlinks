<script setup lang="ts">
import type { Component } from "vue";

const emit = defineEmits(["update:modelValue"]);

defineProps<{
  placeholder?: string;
  leading?: Component;
  modelValue?: string;
  type?: "text" | "password";
}>();
</script>

<template>
  <div
    class="flex items-center gap-3 px-4 bg-neutral-100 outline outline-1 outline-neutral-300 rounded-lg focus-within:outline focus-within:outline-1 focus-within:outline-brand-300 focus-within:shadow-inputShadow transition-shadow"
  >
    <component class="fill-neutral-400" v-if="leading" :is="leading" />
    <input
      class="w-full h-full py-3 outline-none caret-brand-300"
      :type="type || 'text'"
      v-bind="{ placeholder }"
      :value="modelValue"
      @input="
        (event: Event) => {
          emit('update:modelValue', (event.target as HTMLInputElement).value);
        }
      "
    />
  </div>
</template>
