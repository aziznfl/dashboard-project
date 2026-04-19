import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/presentation/store/auth'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('@/presentation/views/Login.vue'),
      meta: { guest: true }
    },
    {
      path: '/',
      name: 'dashboard',
      component: () => import('@/presentation/views/Dashboard.vue'),
      meta: { auth: true }
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: '/'
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  if (to.meta.auth && !authStore.isAuthenticated) {
    next('/login')
  } else if (to.meta.guest && authStore.isAuthenticated) {
    next('/')
  } else {
    next()
  }
})

export default router
