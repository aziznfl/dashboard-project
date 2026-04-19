<script setup>
import { Search, Filter } from 'lucide-vue-next';
import BaseInput from '@/presentation/components/common/BaseInput.vue';
import BaseButton from '@/presentation/components/common/BaseButton.vue';

const props = defineProps({
  modelValue: {
    type: Object,
    required: true
  }
});

const emit = defineEmits(['update:modelValue', 'apply']);

const handleInput = (key, value) => {
  emit('update:modelValue', { ...props.modelValue, [key]: value });
};
</script>

<template>
  <div class="glass-card p-6 flex flex-col md:flex-row gap-4 items-end">
    <div class="flex-1 w-full space-y-1.5">
      <BaseInput 
        label="Search by ID" 
        :model-value="modelValue.id"
        @update:model-value="handleInput('id', $event)"
        placeholder="Ex: 550e8400..."
        :icon="Search"
        @keyup.enter="$emit('apply')"
      />
    </div>
    
    <div class="w-full md:w-64 space-y-1.5">
      <label class="text-sm font-medium text-surface-400 ml-1">Status</label>
      <div class="relative">
        <div class="absolute left-3 top-1/2 -translate-y-1/2 text-surface-500 pointer-events-none">
          <Filter :size="18" />
        </div>
        <select 
          :value="modelValue.status"
          @change="handleInput('status', $event.target.value)"
          class="glass-input w-full pl-10 appearance-none bg-surface-900"
        >
          <option value="">All Statuses</option>
          <option value="completed">Completed</option>
          <option value="processing">Processing</option>
          <option value="failed">Failed</option>
        </select>
      </div>
    </div>
    
    <BaseButton @click="$emit('apply')" class="w-full md:w-auto">
      <Search :size="18" class="mr-2" />
      Apply Filters
    </BaseButton>
  </div>
</template>

<style scoped>
select {
  background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3e%3cpath stroke='%2371717a' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='M6 8l4 4 4-4'/%3e%3c/svg%3e");
  background-position: right 0.5rem center;
  background-repeat: no-repeat;
  background-size: 1.5em 1.5em;
  padding-right: 2.5rem;
}
</style>
