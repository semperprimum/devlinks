<script setup lang="ts">
import Input from "@/components/Input.vue";
import Envelope from "@/components/icons/Envelope.vue";
import Lock from "@/components/icons/Lock.vue";
import LogoFull from "@/components/icons/LogoFull.vue";
import Label from "@/components/Label.vue";
import Button from "@/components/Button.vue";
import { useRouter, type RouterLink } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import * as yup from "yup";
import { useForm } from "vee-validate";

const authStore = useAuthStore();
const router = useRouter();

interface FormData {
  email: string;
  password: string;
}

const validationSchema = yup.object({
  email: yup.string().email("Please check again").required("Can't be empty"),
  password: yup
    .string()
    .required("Can't be empty")
    .min(8, "Please check again"),
});

const { defineField, handleSubmit, errors } = useForm<FormData>({
  validationSchema,
});

const [email] = defineField("email");
const [password] = defineField("password");

const onSubmit = handleSubmit(async (values) => {
  await authStore.login(values.email, values.password);

  router.replace("/");
});
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
        <h2 class="text-2xl font-bold mb-3 md:text-[2rem]">Login</h2>
        <p class="text-neutral-400 mb-10">
          Add your details below to get back into the app
        </p>

        <form @submit.prevent="onSubmit">
          <Label :error="errors.email" text="Email address" />
          <Input
            v-model="email"
            class="mb-6 mt-1"
            :error="errors.email"
            placeholder="e.g. alex@email.com"
            :leading="Envelope"
          />

          <Label :error="errors.password" text="Password" />
          <Input
            v-model="password"
            type="password"
            class="mt-1"
            :error="errors.password"
            placeholder="Enter your password"
            :leading="Lock"
          />

          <Button class="w-full mt-6">Log in</Button>
        </form>

        <p class="text-neutral-400 w-full text-center mt-6">
          Don’t have an account?
          <RouterLink to="/register" class="text-brand-300 cursor-pointer"
            >Create account</RouterLink
          >
        </p>
      </div>
    </div>
  </div>
</template>
