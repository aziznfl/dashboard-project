<script setup>
import { computed } from 'vue';
import { 
  BarChart3, 
  TrendingUp, 
  CheckCircle2, 
  AlertCircle 
} from 'lucide-vue-next';
import BaseCard from '@/presentation/components/common/BaseCard.vue';

const props = defineProps({
  payments: {
    type: Array,
    required: true
  },
  meta: {
    type: Object,
    required: true
  },
  loading: Boolean
});

/**
 * Total Volume calculation
 * Note: Since backend only provides paginated data, this currently
 * sums the amounts of the payments on the current page.
 */
const totalVolume = computed(() => {
  const amount = props.payments.reduce((sum, p) => sum + (Number(p.amount) || 0), 0);
  return amount.toCurrency();
});

const successRate = computed(() => {
  if (props.payments.length === 0) return '0%';
  const completed = props.payments.filter(p => p.status === 'completed').length;
  return Math.round((completed / props.payments.length) * 100) + '%';
});

const failedCount = computed(() => {
  return props.payments.filter(p => p.status === 'failed').length;
});

const stats = computed(() => [
  {
    label: 'Total Volume',
    value: totalVolume.value,
    icon: TrendingUp,
    color: 'text-primary-400',
    bg: 'bg-primary-400/10',
    description: 'Current page volume'
  },
  {
    label: 'Total Transactions',
    value: props.meta.total.toLocaleString(),
    icon: BarChart3,
    color: 'text-surface-100',
    bg: 'bg-surface-100/10',
    description: 'Across all records'
  },
  {
    label: 'Success Rate',
    value: successRate.value,
    icon: CheckCircle2,
    color: 'text-emerald-400',
    bg: 'bg-emerald-400/10',
    description: 'Current page distribution'
  },
  {
    label: 'Failed Payments',
    value: failedCount.value,
    icon: AlertCircle,
    color: 'text-rose-400',
    bg: 'bg-rose-400/10',
    description: 'Current page failures'
  }
]);
</script>

<template>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
    <template v-if="loading">
      <div v-for="i in 4" :key="i" class="glass-card animate-pulse h-28 border border-white/5 bg-white/5 opacity-50"></div>
    </template>
    
    <template v-else>
      <BaseCard v-for="stat in stats" :key="stat.label" no-padding>
        <div class="p-5 flex items-start justify-between">
          <div class="space-y-1">
            <p class="text-[10px] font-semibold text-surface-400 uppercase tracking-widest">{{ stat.label }}</p>
            <h4 class="text-2xl font-bold text-white tracking-tight">{{ stat.value }}</h4>
            <p class="text-[10px] text-surface-500 font-medium">{{ stat.description }}</p>
          </div>
          <div :class="['p-2.5 rounded-xl shrink-0 border border-white/5 shadow-lg', stat.bg, stat.color]">
            <component :is="stat.icon" :size="20" />
          </div>
        </div>
      </BaseCard>
    </template>
  </div>
</template>
