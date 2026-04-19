import { defineStore } from 'pinia';
import { AuthService } from '@/application/services/AuthService';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: JSON.parse(localStorage.getItem('auth')) || null,
    loading: false,
    error: null,
  }),
  
  getters: {
    isAuthenticated: (state) => !!state.user?.token,
  },
  
  actions: {
    async login(email, password) {
      this.loading = true;
      this.error = null;
      try {
        const user = await AuthService.login(email, password);
        this.user = user;
        localStorage.setItem('auth', JSON.stringify(user));
        return user;
      } catch (err) {
        this.error = err;
        throw err;
      } finally {
        this.loading = false;
      }
    },
    
    logout() {
      this.user = null;
      localStorage.removeItem('auth');
      window.location.href = '/login';
    }
  }
});
