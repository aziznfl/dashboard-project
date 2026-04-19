<script setup>
import { CreditCard, LogOut } from 'lucide-vue-next';
import BaseButton from '@/presentation/components/common/BaseButton.vue';

defineProps({
  userEmail: String
});

defineEmits(['logout']);
</script>

<template>
  <nav class="sticky top-0 z-30 border-b border-surface-800 bg-surface-950/80 backdrop-blur-md">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">
        <div class="flex items-center gap-2">
          <div class="w-8 h-8 rounded-lg bg-primary-600 flex items-center justify-center text-white shadow-lg shadow-primary-500/20">
            <CreditCard :size="20" />
          </div>
          <span class="text-xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-white to-surface-500">
            PayDash
          </span>
        </div>
        
        <div class="flex items-center gap-4">
          <div v-if="userEmail" class="hidden sm:flex items-center gap-2 px-3 py-1.5 rounded-full bg-surface-900 border border-surface-800">
            <div class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse" />
            <span class="text-xs font-medium text-surface-400">{{ userEmail }}</span>
          </div>
          
          <BaseButton variant="ghost" @click="$emit('logout')" class="!px-3 flex items-center gap-2">
            <LogOut :size="18" />
            <span class="hidden sm:inline">Logout</span>
          </BaseButton>
        </div>
      </div>
    </div>
  </nav>
</template>
