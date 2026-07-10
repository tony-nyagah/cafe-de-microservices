import { createRouter, createWebHistory } from "vue-router";
import MenuPanel from "../components/MenuPanel.vue";
import ChatPanel from "../components/ChatPanel.vue";

const routes = [
  { path: "/", redirect: "/menu" },
  { path: "/menu", component: MenuPanel },
  { path: "/chat", component: ChatPanel },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
