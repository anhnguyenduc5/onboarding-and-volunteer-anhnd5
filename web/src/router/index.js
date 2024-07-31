import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../components/HomePage.vue'
import IndexPage from '../components/IndexPage.vue'

const routes = [
  { path: '/', name: 'Index', component: IndexPage },
  { path: '/home', name: 'Home', component: HomePage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
