<script setup>
import { ChevronUp, ChevronDown, MonitorOff } from 'lucide-vue-next';

const props = defineProps({
  columns: {
    type: Array, // { key, label, sortable, align }
    required: true
  },
  items: {
    type: Array,
    required: true
  },
  loading: Boolean,
  sortKey: String,
  sortOrder: {
    type: String,
    default: 'asc'
  }
})

const emit = defineEmits(['sort'])

const handleSort = (column) => {
  if (!column.sortable) return;
  
  let newOrder = 'asc';
  if (props.sortKey === column.key) {
    newOrder = props.sortOrder === 'asc' ? 'desc' : 'asc';
  }
  
  emit('sort', { key: column.key, order: newOrder });
}
</script>

<template>
  <div class="glass-card flex flex-col h-full">
    <div class="overflow-x-auto">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="border-b border-surface-800 bg-surface-900/30">
            <th 
              v-for="col in columns" 
              :key="col.key"
              class="px-6 py-4 text-xs font-semibold text-surface-400 uppercase tracking-wider"
              :class="[
                col.align === 'right' ? 'text-right' : 'text-left',
                col.sortable ? 'cursor-pointer hover:text-surface-200 transition-colors' : ''
              ]"
              @click="handleSort(col)"
            >
              <div class="flex items-center gap-1" :class="col.align === 'right' ? 'justify-end' : ''">
                {{ col.label }}
                <div v-if="col.sortable" class="flex flex-col text-surface-600 scale-75">
                  <ChevronUp 
                    :size="14" 
                    :class="{ 'text-primary-500': sortKey === col.key && sortOrder === 'asc' }" 
                    class="-mb-1"
                  />
                  <ChevronDown 
                    :size="14" 
                    :class="{ 'text-primary-500': sortKey === col.key && sortOrder === 'desc' }" 
                    class="-mt-1"
                  />
                </div>
              </div>
            </th>
          </tr>
        </thead>
        
        <tbody class="relative">
          <tr v-if="loading" class="animate-pulse">
            <td :colspan="columns.length" class="px-6 py-12">
              <div class="flex flex-col items-center gap-4 text-surface-500">
                <div class="w-10 h-10 border-2 border-primary-500/30 border-t-primary-500 rounded-full animate-spin" />
                <span>Loading records...</span>
              </div>
            </td>
          </tr>
          
          <template v-else-if="items.length > 0">
            <tr 
              v-for="(item, idx) in items" 
              :key="item.id || idx"
              class="border-b border-surface-800/50 hover:bg-white/[0.02] transition-colors group"
            >
              <td 
                v-for="col in columns" 
                :key="col.key"
                class="px-6 py-4 text-sm"
                :class="col.align === 'right' ? 'text-right' : 'text-left'"
              >
                <slot :name="`cell(${col.key})`" :item="item">
                  {{ item[col.key] }}
                </slot>
              </td>
            </tr>
          </template>
          
          <tr v-else>
            <td :colspan="columns.length" class="px-6 py-20 text-center text-surface-500">
              <div class="flex flex-col items-center gap-2">
                <MonitorOff :size="40" class="opacity-20" />
                <p>No data found matching your criteria</p>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
