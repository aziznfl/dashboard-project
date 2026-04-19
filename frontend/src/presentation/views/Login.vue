<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/presentation/store/auth';
import { Mail, Lock, LogIn } from 'lucide-vue-next';
import BaseInput from '@/presentation/components/common/BaseInput.vue';
import BaseButton from '@/presentation/components/common/BaseButton.vue';

const router = useRouter();
const authStore = useAuthStore();

const email = ref('');
const password = ref('');
const error = ref('');
const loading = ref(false);

const handleLogin = async () => {
  if (!email.value || !password.value) {
    error.value = 'Please fill in all fields';
    return;
  }

  loading.value = true;
  error.value = '';
  
  try {
    await authStore.login(email.value, password.value);
    router.push('/');
  } catch (err) {
    error.value = err;
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div class="flex min-h-screen items-center justify-center p-6 sm:p-12 relative overflow-hidden bg-surface-950">
    <!-- Animated background blob -->
    <div class="absolute top-0 -left-4 w-72 h-72 bg-primary-600/20 rounded-full blur-[128px] animate-pulse" />
    <div class="absolute bottom-0 -right-4 w-96 h-96 bg-primary-900/10 rounded-full blur-[128px]" />

    <div class="w-full max-w-md space-y-8 z-10">
      <div class="text-center">
        <div class="inline-flex items-center justify-center w-16 h-16 rounded-2xl bg-primary-600/20 text-primary-500 mb-6 border border-primary-500/20 shadow-xl shadow-primary-500/10">
          <LogIn :size="32" />
        </div>
        <h2 class="text-3xl font-bold tracking-tight text-white">Welcome back</h2>
        <p class="mt-2 text-surface-400">Please enter your credentials to access your dashboard</p>
      </div>

      <form class="space-y-6" @submit.prevent="handleLogin">
        <div class="glass-card p-8 space-y-5">
          <BaseInput
            label="Email Address"
            v-model="email"
            type="email"
            placeholder="name@company.com"
            :icon="Mail"
            :error="error && error.includes('email') ? error : ''"
          />

          <BaseInput
            label="Password"
            v-model="password"
            type="password"
            placeholder="••••••••"
            :icon="Lock"
            :error="error && error.includes('password') ? error : ''"
          />

          <div v-if="error && !error.includes('email') && !error.includes('password')" class="text-sm text-rose-500 bg-rose-500/10 border border-rose-500/20 p-3 rounded-lg flex items-center gap-2 animate-in fade-in zoom-in duration-300">
            <span class="w-1.5 h-1.5 rounded-full bg-rose-500" />
            {{ error }}
          </div>

          <BaseButton
            type="submit"
            class="w-full"
            :loading="loading"
          >
            Sign in to Dashboard
          </BaseButton>
        </div>
      </form>
      
      <p class="text-center text-sm text-surface-500">
        Demo Credentials: <span class="text-surface-300">admin@example.com / password</span>
      </p>
    </div>
  </div>
</template>
