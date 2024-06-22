<script setup lang="ts">
import Input from "@/components/Input.vue";
import Envelope from "@/components/icons/Envelope.vue";
import Lock from "@/components/icons/Lock.vue";
import LogoFull from "@/components/icons/LogoFull.vue";
import Label from "@/components/Label.vue";
import Button from "@/components/Button.vue";
import { RouterLink } from "vue-router";
</script>

<template>
  <div
    class="p-8 md:min-h-screen md:grid md:place-items-center md:grid-flow-row-dense"
  >
    <div>
      <LogoFull class="md:mx-auto" />

      <div
        class="mt-16 md:bg-neutral-100 md:p-10 md:rounded-xl md:w-[29.75rem]"
      >
        <h2 class="text-2xl font-bold mb-3 md:text-[2rem]">Create account</h2>
        <p class="text-neutral-400 mb-10">
          Let’s get you started sharing your links!
        </p>

        <Label text="Email address" />
        <Input
          class="mb-6 mt-1"
          placeholder="e.g. alex@email.com"
          :leading="Envelope"
        />

        <Label text="Create password" />
        <Input
          class="mt-1 mb-6"
          placeholder="At least .8 characters"
          :leading="Lock"
        />

        <Label text="Confirm password" />
        <Input
          class="mt-1"
          placeholder="At least .8 characters"
          :leading="Lock"
        />

        <p class="text-neutral-400 text-xs mt-6">
          Password must contain at least 8 characters
        </p>

        <Button class="w-full mt-6">Login</Button>

        <p class="text-neutral-400 w-full text-center mt-6">
          Already have an account?
          <RouterLink to="/login" class="text-brand-300 cursor-pointer"
            >Login</RouterLink
          >
        </p>
      </div>
    </div>
  </div>
</template>
