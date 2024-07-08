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
    {
      path: "/preview",
      name: "preview",
      component: () => import("@/views/Preview.vue"),
    },
    {
      path: "/:id",
      name: "link",
      component: () => import("@/views/LinkView.vue"),
    },
    {
      path: "/notfound",
      name: "404",
      component: () => import("@/views/404.vue"),
    },
  ],
});

router.beforeEach((to, _) => {
  if (
    !localStorage.getItem("auth-token") &&
    to.name !== "link" &&
    to.name !== "login" &&
    to.name !== "register"
  ) {
    console.log("Redirect to login");
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
