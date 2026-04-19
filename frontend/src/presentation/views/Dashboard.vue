<script setup>
import { onMounted, ref, computed } from 'vue';
import { useAuthStore } from '@/presentation/store/auth';
import { usePaymentStore } from '@/presentation/store/payments';
import { 
  LogOut, 
  Search, 
  Filter, 
  RefreshCw, 
  CreditCard,
  User as UserIcon,
  ShoppingBag,
  Calendar,
  CheckCircle2,
  Clock,
  AlertCircle
} from 'lucide-vue-next';
import BaseTable from '@/presentation/components/common/BaseTable.vue';
import BaseButton from '@/presentation/components/common/BaseButton.vue';
import BaseInput from '@/presentation/components/common/BaseInput.vue';
import BaseCard from '@/presentation/components/common/BaseCard.vue';

const authStore = useAuthStore();
const paymentStore = usePaymentStore();

const columns = [
  { key: 'id', label: 'Payment ID', sortable: true },
  { key: 'merchant', label: 'Merchant', sortable: true },
  { key: 'amount', label: 'Amount', sortable: true, align: 'right' },
  { key: 'status', label: 'Status', sortable: true },
  { key: 'createdAt', label: 'Date', sortable: true, align: 'right' }
];

const filters = ref({
  status: '',
  id: ''
});

onMounted(() => {
  paymentStore.fetchPayments();
});

const handleSort = (sortData) => {
  const sortPrefix = sortData.order === 'desc' ? '-' : '';
  const sortValue = `${sortPrefix}${sortData.key === 'createdAt' ? 'created_at' : sortData.key}`;
  paymentStore.setFilters({ sort: sortValue });
};

const handleRefresh = () => {
  paymentStore.fetchPayments();
};

const applyFilters = () => {
  paymentStore.setFilters({
    status: filters.value.status,
    id: filters.value.id
  });
};

const logout = () => {
  authStore.logout();
};

// Summary metrics (mock/derived)
const totalAmount = computed(() => {
  return paymentStore.payments.reduce((sum, p) => sum + parseFloat(p.amount), 0);
});

const getStatusIcon = (status) => {
  switch (status) {
    case 'completed': return CheckCircle2;
    case 'processing': return Clock;
    case 'failed': return AlertCircle;
    default: return RefreshCw;
  }
};
</script>

<template>
  <div class="min-h-screen bg-surface-950 text-surface-50">
    <!-- Navigation -->
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
            <div class="hidden sm:flex items-center gap-2 px-3 py-1.5 rounded-full bg-surface-900 border border-surface-800">
              <div class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse" />
              <span class="text-xs font-medium text-surface-400">{{ authStore.user?.email }}</span>
            </div>
            
            <BaseButton variant="ghost" @click="logout" class="!px-3 flex items-center gap-2">
              <LogOut :size="18" />
              <span class="hidden sm:inline">Logout</span>
            </BaseButton>
          </div>
        </div>
      </div>
    </nav>

    <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 space-y-8">
      <!-- Header Section -->
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
        <div>
          <h1 class="text-3xl font-bold text-white tracking-tight">Payment Transactions</h1>
          <p class="text-surface-400 mt-1">Monitor and manage all incoming payments across platforms.</p>
        </div>
        
        <div class="flex items-center gap-3">
          <BaseButton variant="secondary" @click="handleRefresh" :loading="paymentStore.loading">
            <RefreshCw :size="18" class="mr-2" />
            Refresh
          </BaseButton>
        </div>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        <BaseCard class="!p-0 overflow-hidden relative group">
          <div class="p-6">
            <div class="flex justify-between items-start mb-4">
              <div class="p-2 rounded-lg bg-primary-500/10 text-primary-500">
                <ShoppingBag :size="20" />
              </div>
            </div>
            <p class="text-sm font-medium text-surface-400 uppercase tracking-wider">Total Payments</p>
            <h4 class="text-2xl font-bold text-white mt-1">{{ paymentStore.payments.length }}</h4>
          </div>
          <div class="absolute bottom-0 left-0 h-1 bg-primary-600 transition-all duration-300 w-full opacity-30 group-hover:opacity-100" />
        </BaseCard>

        <BaseCard class="!p-0 overflow-hidden relative group">
          <div class="p-6">
            <div class="flex justify-between items-start mb-4">
              <div class="p-2 rounded-lg bg-emerald-500/10 text-emerald-500">
                <CheckCircle2 :size="20" />
              </div>
            </div>
            <p class="text-sm font-medium text-surface-400 uppercase tracking-wider">Total Volume</p>
            <h4 class="text-2xl font-bold text-white mt-1">
              {{ new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(totalAmount) }}
            </h4>
          </div>
          <div class="absolute bottom-0 left-0 h-1 bg-emerald-600 transition-all duration-300 w-full opacity-30 group-hover:opacity-100" />
        </BaseCard>

        <BaseCard class="!p-0 overflow-hidden relative group">
          <div class="p-6">
            <div class="flex justify-between items-start mb-4">
              <div class="p-2 rounded-lg bg-surface-500/10 text-surface-500">
                <UserIcon :size="20" />
              </div>
            </div>
            <p class="text-sm font-medium text-surface-400 uppercase tracking-wider">Active Role</p>
            <h4 class="text-2xl font-bold text-white mt-1 capitalize">{{ authStore.user?.role || 'User' }}</h4>
          </div>
          <div class="absolute bottom-0 left-0 h-1 bg-surface-600 transition-all duration-300 w-full opacity-30 group-hover:opacity-100" />
        </BaseCard>

        <BaseCard class="!p-0 overflow-hidden relative group">
          <div class="p-6">
            <div class="flex justify-between items-start mb-4">
              <div class="p-2 rounded-lg bg-amber-500/10 text-amber-500">
                <Calendar :size="20" />
              </div>
            </div>
            <p class="text-sm font-medium text-surface-400 uppercase tracking-wider">Current Date</p>
            <h4 class="text-2xl font-bold text-white mt-1">
              {{ new Date().toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' }) }}
            </h4>
          </div>
          <div class="absolute bottom-0 left-0 h-1 bg-amber-600 transition-all duration-300 w-full opacity-30 group-hover:opacity-100" />
        </BaseCard>
      </div>

      <!-- Filters & Table Section -->
      <div class="space-y-6">
        <!-- Filters -->
        <div class="glass-card p-6 flex flex-col md:flex-row gap-4 items-end">
          <div class="flex-1 w-full space-y-1.5">
            <BaseInput 
              label="Search by ID" 
              v-model="filters.id" 
              placeholder="Ex: 550e8400..."
              :icon="Search"
              @keyup.enter="applyFilters"
            />
          </div>
          
          <div class="w-full md:w-64 space-y-1.5">
            <label class="text-sm font-medium text-surface-400 ml-1">Status</label>
            <div class="relative">
              <div class="absolute left-3 top-1/2 -translate-y-1/2 text-surface-500 pointer-events-none">
                <Filter :size="18" />
              </div>
              <select 
                v-model="filters.status"
                class="glass-input w-full pl-10 appearance-none bg-surface-900"
              >
                <option value="">All Statuses</option>
                <option value="completed">Completed</option>
                <option value="processing">Processing</option>
                <option value="failed">Failed</option>
              </select>
            </div>
          </div>
          
          <BaseButton @click="applyFilters" class="w-full md:w-auto">
            <Search :size="18" class="mr-2" />
            Apply Filters
          </BaseButton>
        </div>

        <!-- Table -->
        <BaseTable
          :columns="columns"
          :items="paymentStore.payments"
          :loading="paymentStore.loading"
          :sort-key="paymentStore.filters.sort.replace('-', '')"
          :sort-order="paymentStore.filters.sort.startsWith('-') ? 'desc' : 'asc'"
          @sort="handleSort"
        >
          <template #cell(id)="{ item }">
            <code class="text-xs font-mono text-primary-400 bg-primary-400/10 px-2 py-1 rounded">
              {{ item.id }}
            </code>
          </template>

          <template #cell(merchant)="{ item }">
            <span class="font-medium text-surface-100">{{ item.merchant }}</span>
          </template>

          <template #cell(amount)="{ item }">
            <span class="font-medium text-white">{{ item.formattedAmount }}</span>
          </template>

          <template #cell(status)="{ item }">
            <div class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium border" :class="item.statusColor">
              <component :is="getStatusIcon(item.status)" :size="14" />
              <span class="capitalize">{{ item.status }}</span>
            </div>
          </template>

          <template #cell(createdAt)="{ item }">
            <span class="text-surface-400">{{ item.formattedDate }}</span>
          </template>
        </BaseTable>
      </div>
    </main>

    <!-- Footer -->
    <footer class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12 border-t border-surface-800 text-center">
      <p class="text-surface-500 text-sm">
        &copy; 2026 PayDash Inc. Built with Vue 3 & Tailwind CSS. Clean Architecture principles applied.
      </p>
    </footer>
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
