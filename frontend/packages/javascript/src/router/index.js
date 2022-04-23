import { createRouter, createWebHashHistory, createWebHistory } from "vue-router";
import Login from "@/views/Login.vue";
import Register from "@/views/Register.vue";
import Main from "@/views/Main.vue";

const routes = [
  {
    path: "/",
    redirect: "/login",
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
  {
    path: "/register",
    name: "Register",
    component: Register,
    meta: {
      transition: 'slide-down',
    }
  },
  {
    path: "/main",
    name: "Main",
    component: Main,
  }
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
