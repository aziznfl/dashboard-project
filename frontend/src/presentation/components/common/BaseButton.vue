<script setup>
defineProps({
  variant: {
    type: String,
    default: 'primary' // primary, secondary, danger, ghost
  },
  loading: Boolean,
  disabled: Boolean,
  type: {
    type: String,
    default: 'button'
  }
})
</script>

<template>
  <button
    :type="type"
    :disabled="disabled || loading"
    class="relative inline-flex items-center justify-center px-6 py-2.5 font-medium transition-all duration-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-surface-950 disabled:opacity-50 disabled:cursor-not-allowed group overflow-hidden"
    :class="{
      'bg-primary-600 text-white hover:bg-primary-500 focus:ring-primary-500 shadow-lg shadow-primary-500/20': variant === 'primary',
      'bg-surface-800 text-surface-50 hover:bg-surface-700 focus:ring-surface-700': variant === 'secondary',
      'bg-rose-600 text-white hover:bg-rose-500 focus:ring-rose-500': variant === 'danger',
      'bg-transparent text-surface-300 hover:text-surface-50 hover:bg-surface-800 focus:ring-surface-800 border border-surface-800': variant === 'ghost'
    }"
  >
    <div v-if="loading" class="absolute inset-0 flex items-center justify-center bg-inherit">
      <svg class="w-5 h-5 animate-spin" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none" />
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
      </svg>
    </div>
    
    <span :class="{ 'opacity-0': loading }" class="flex items-center gap-2">
      <slot />
    </span>
    
    <!-- Shine effect -->
    <div class="absolute inset-0 w-1/2 h-full transition-transform duration-500 -translate-x-full bg-gradient-to-r from-transparent via-white/10 to-transparent group-hover:translate-x-[200%] ease-in-out" />
  </button>
</template>
