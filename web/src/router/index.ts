import { createRouter, createWebHistory } from 'vue-router'
import { BASE } from '../base'

const router = createRouter({
  history: createWebHistory(BASE || '/'),
  routes: [
    {
      path: '/login',
      component: () => import('../pages/LoginPage.vue'),
    },
    {
      path: '/',
      component: () => import('../layout/MainLayout.vue'),
      children: [
        { path: '', component: () => import('../pages/DashboardPage.vue') },
        { path: 'users', component: () => import('../pages/UsersPage.vue') },
        { path: 'network', component: () => import('../pages/NetworkPage.vue') },
        { path: 'outbounds', component: () => import('../pages/OutboundsPage.vue') },
        { path: 'chain', component: () => import('../pages/ChainPage.vue') },
        { path: 'settings', component: () => import('../pages/SettingsPage.vue') },
      ],
    },
  ],
})

export default router
