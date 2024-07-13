<script setup lang="ts">
import Input from "@/components/Input.vue";
import Envelope from "@/components/icons/Envelope.vue";
import Lock from "@/components/icons/Lock.vue";
import LogoFull from "@/components/icons/LogoFull.vue";
import Label from "@/components/Label.vue";
import Button from "@/components/Button.vue";
import { RouterLink } from "vue-router";
import * as yup from "yup";
import { useForm } from "vee-validate";
import { useAuthStore } from "@/stores/auth";

const authStore = useAuthStore();

interface FormData {
  email: string;
  password: string;
  confirmPassword: string;
}

const validationSchema = yup.object({
  email: yup.string().email("Must be a valid Email").required("Can't be empty"),
  password: yup
    .string()
    .required("Can't be empty")
    .min(8, "Please check again"),
  confirmPassword: yup
    .string()
    .required("Can't be empty")
    .oneOf([yup.ref("password")], "Please check again"),
});

const { defineField, handleSubmit, errors } = useForm<FormData>({
  validationSchema,
});

const [email] = defineField("email");
const [password] = defineField("password");
const [confirmPassword] = defineField("confirmPassword");

const onSubmit = handleSubmit(async (values) => {
  await authStore.register(values.email, values.password);
});
</script>

<template>
  <main
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

        <form @submit.prevent="onSubmit">
          <Label :error="errors.email" text="Email address" />
          <Input
            v-model="email"
            :error="errors.email"
            class="mb-6 mt-1"
            placeholder="e.g. alex@email.com"
            :leading="Envelope"
          />

          <Label :error="errors.password" text="Create password" />
          <Input
            v-model="password"
            :error="errors.password"
            type="password"
            class="mt-1 mb-6"
            placeholder="At least .8 characters"
            :leading="Lock"
          />

          <Label :error="errors.confirmPassword" text="Confirm password" />
          <Input
            v-model="confirmPassword"
            :error="errors.confirmPassword"
            type="password"
            class="mt-1"
            placeholder="At least .8 characters"
            :leading="Lock"
          />

          <p class="text-neutral-400 text-xs mt-6">
            Password must contain at least 8 characters
          </p>

          <Button class="w-full mt-6">Register</Button>
        </form>

        <p class="text-neutral-400 w-full text-center mt-6">
          Already have an account?
          <RouterLink to="/login" class="text-brand-300 cursor-pointer"
            >Login</RouterLink
          >
        </p>
      </div>
    </div>
  </main>
</template>
