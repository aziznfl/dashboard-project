<script setup>
defineProps({
  title: String,
  subtitle: String,
  noPadding: Boolean
})
</script>

<template>
  <div class="glass-card">
    <div v-if="title || $slots.header" class="px-6 py-4 border-b border-surface-800 flex justify-between items-center">
      <div>
        <h3 v-if="title" class="text-lg font-semibold text-surface-50">{{ title }}</h3>
        <p v-if="subtitle" class="text-sm text-surface-400 mt-0.5">{{ subtitle }}</p>
      </div>
      <slot name="header-action" />
    </div>
    
    <div :class="{ 'p-6': !noPadding }">
      <slot />
    </div>
    
    <div v-if="$slots.footer" class="px-6 py-4 border-t border-surface-800 bg-surface-900/30">
      <slot name="footer" />
    </div>
  </div>
</template>
