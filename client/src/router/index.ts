import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: () => import("@/views/HomeView.vue"),
    },
    {
      path: "/profile",
      name: "profile",
      component: () => import("@/views/Profile.vue"),
    },
    {
      path: "/login",
      name: "login",
      component: () => import("@/views/LoginView.vue"),
    },
    {
      path: "/register",
      name: "register",
      component: () => import("@/views/RegisterView.vue"),
    },
  ],
});

router.beforeEach(async (to, _) => {
  if (
    !localStorage.getItem("auth-token") &&
    to.name !== "login" &&
    to.name !== "register"
  ) {
    return { name: "login" };
  }

  if (
    localStorage.getItem("auth-token") &&
    (to.name === "login" || to.name === "register")
  ) {
    return { name: "home" };
  }
});

export default router;
