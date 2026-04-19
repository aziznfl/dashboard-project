<script setup>
defineProps({
  modelValue: [String, Number],
  label: String,
  type: {
    type: String,
    default: 'text'
  },
  placeholder: String,
  error: String,
  icon: Object // Optional icon component
})

const emit = defineEmits(['update:modelValue'])
</script>

<template>
  <div class="flex flex-col gap-1.5 w-full">
    <label v-if="label" class="text-sm font-medium text-surface-400 ml-1">
      {{ label }}
    </label>
    
    <div class="relative group">
      <div v-if="icon" class="absolute left-3 top-1/2 -translate-y-1/2 text-surface-500 transition-colors group-focus-within:text-primary-500">
        <component :is="icon" :size="18" />
      </div>
      
      <input
        :type="type"
        :value="modelValue"
        @input="emit('update:modelValue', $event.target.value)"
        :placeholder="placeholder"
        class="glass-input w-full"
        :class="{
          'pl-10': icon,
          'border-rose-500 focus:border-rose-500 focus:ring-rose-500/50': error,
          'border-surface-800': !error
        }"
      >
    </div>
    
    <p v-if="error" class="text-xs text-rose-500 ml-1 mt-0.5 animate-in fade-in slide-in-from-top-1 duration-200">
      {{ error }}
    </p>
  </div>
</template>
