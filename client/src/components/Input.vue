<script setup lang="ts">
import type { Component } from "vue";

const emit = defineEmits(["update:modelValue"]);

defineProps<{
  placeholder?: string;
  leading?: Component;
  modelValue?: string;
  type?: "text" | "password";
  error?: string;
}>();
</script>

<template>
  <div
    class="flex items-center gap-3 px-4 bg-neutral-100 outline outline-1 rounded-lg focus-within:shadow-inputShadow transition-shadow"
    :class="{
      ['outline-red']: error,
      ['outline-neutral-300 focus-within:outline-brand-300 ']: !error,
    }"
  >
    <component class="fill-neutral-400 shrink-0" v-if="leading" :is="leading" />
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
    <span class="text-xs text-red whitespace-nowrap" v-if="error">{{
      error
    }}</span>
  </div>
</template>
