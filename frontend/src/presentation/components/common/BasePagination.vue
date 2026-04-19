<script setup>
import { ChevronLeft, ChevronRight, ChevronsLeft, ChevronsRight } from 'lucide-vue-next';
import { computed } from 'vue';

const props = defineProps({
  currentPage: {
    type: Number,
    required: true
  },
  totalPages: {
    type: Number,
    required: true
  },
  totalItems: {
    type: Number,
    default: 0
  },
  pageSize: {
    type: Number,
    default: 10
  }
});

const emit = defineEmits(['page-change', 'limit-change']);

const pages = computed(() => {
  const range = 2; // Show 2 pages before and after current
  const history = [];
  
  let start = Math.max(1, props.currentPage - range);
  let end = Math.min(props.totalPages, props.currentPage + range);
  
  if (start > 1) {
    history.push(1);
    if (start > 2) history.push('...');
  }
  
  for (let i = start; i <= end; i++) {
    history.push(i);
  }
  
  if (end < props.totalPages) {
    if (end < props.totalPages - 1) history.push('...');
    history.push(props.totalPages);
  }
  
  return history;
});

const startItem = computed(() => (props.currentPage - 1) * props.pageSize + 1);
const endItem = computed(() => Math.min(props.currentPage * props.pageSize, props.totalItems));

const handlePageChange = (page) => {
  if (page === '...' || page === props.currentPage || page < 1 || page > props.totalPages) return;
  emit('page-change', page);
};
</script>

<template>
  <div class="flex flex-col sm:flex-row items-center justify-between gap-4 px-6 py-4 border-t border-surface-800 bg-surface-900/10">
    <div class="text-sm text-surface-400">
      Showing <span class="font-medium text-surface-200">{{ totalItems > 0 ? startItem : 0 }}</span> 
      to <span class="font-medium text-surface-200">{{ endItem }}</span> 
      of <span class="font-medium text-surface-200">{{ totalItems }}</span> results
    </div>

    <div class="flex items-center gap-2">
      <div class="flex items-center bg-surface-900 border border-surface-800 rounded-lg p-1">
        <button 
          @click="handlePageChange(1)"
          :disabled="currentPage === 1"
          class="p-1.5 rounded-md transition-all hover:bg-surface-800 disabled:opacity-30 disabled:cursor-not-allowed text-surface-400"
        >
          <ChevronsLeft :size="16" />
        </button>
        <button 
          @click="handlePageChange(currentPage - 1)"
          :disabled="currentPage === 1"
          class="p-1.5 rounded-md transition-all hover:bg-surface-800 disabled:opacity-30 disabled:cursor-not-allowed text-surface-400"
        >
          <ChevronLeft :size="16" />
        </button>
        
        <div class="h-4 w-px bg-surface-800 mx-1" />

        <div class="flex items-center">
          <button 
            v-for="page in pages" 
            :key="page"
            @click="handlePageChange(page)"
            class="min-w-[32px] h-8 flex items-center justify-center rounded-md text-sm font-medium transition-all"
            :class="[
              page === currentPage 
                ? 'bg-primary-600 text-white shadow-lg shadow-primary-500/20' 
                : page === '...' 
                  ? 'text-surface-600 cursor-default' 
                  : 'text-surface-400 hover:bg-surface-800 hover:text-surface-200'
            ]"
          >
            {{ page }}
          </button>
        </div>

        <div class="h-4 w-px bg-surface-800 mx-1" />

        <button 
          @click="handlePageChange(currentPage + 1)"
          :disabled="currentPage === totalPages"
          class="p-1.5 rounded-md transition-all hover:bg-surface-800 disabled:opacity-30 disabled:cursor-not-allowed text-surface-400"
        >
          <ChevronRight :size="16" />
        </button>
        <button 
          @click="handlePageChange(totalPages)"
          :disabled="currentPage === totalPages"
          class="p-1.5 rounded-md transition-all hover:bg-surface-800 disabled:opacity-30 disabled:cursor-not-allowed text-surface-400"
        >
          <ChevronsRight :size="16" />
        </button>
      </div>
    </div>
  </div>
</template>
