import { createRouter, createWebHashHistory } from 'vue-router'
import ListView from '../views/ListView.vue'
import EditerView from '../views/EditerView.vue'
import LoginView from '../views/LoginView.vue'
import EmailDetailView from '../views/EmailDetailView.vue'
const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: ListView
    },
    {
      path: '/list',
      name: 'list',
      component: ListView
    },
    {
      path: '/editer',
      name: "editer",
      component: EditerView
    },
    {
      path: '/login',
      name: "login",
      component: LoginView
    },
    {
      path: '/detail/:id',
      name: "detail",
      component: EmailDetailView
    }
  ]
})

export default router
