<script setup>
import { onMounted, ref, computed } from 'vue';
import { useAuthStore } from '@/presentation/store/auth';
import { usePaymentStore } from '@/presentation/store/payments';
import { 
  RefreshCw, 
  CheckCircle2,
  Clock,
  AlertCircle,
} from 'lucide-vue-next';

// Components
import Navbar from '@/presentation/components/layout/Navbar.vue';
import PaymentFilters from '@/presentation/components/dashboard/PaymentFilters.vue';
import BaseTable from '@/presentation/components/common/BaseTable.vue';
import BasePagination from '@/presentation/components/common/BasePagination.vue';
import BaseButton from '@/presentation/components/common/BaseButton.vue';

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

const handlePageChange = (page) => {
  paymentStore.setPage(page);
};

// Summary metrics
const totalAmount = computed(() => {
  return paymentStore.payments.reduce((sum, p) => sum + parseFloat(p.amount), 0);
});

const formattedTotalVolume = computed(() => {
  return totalAmount.value.toCurrency();
});

const currentDate = computed(() => {
  return new Date().toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' });
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
    <Navbar :user-email="authStore.user?.email" @logout="logout" />

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

      <!-- Filters & Table Section -->
      <div class="space-y-6">
        <PaymentFilters 
          v-model="filters" 
          @apply="applyFilters" 
        />

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

          <template #footer>
            <BasePagination
              v-if="paymentStore.meta.totalPages > 1"
              :current-page="paymentStore.meta.page"
              :total-pages="paymentStore.meta.totalPages"
              :total-items="paymentStore.meta.total"
              :page-size="paymentStore.meta.limit"
              @page-change="handlePageChange"
            />
          </template>
        </BaseTable>
      </div>
    </main>

    <!-- Footer -->
    <footer class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12 border-t border-surface-800 text-center text-surface-500 text-sm">
      &copy; 2026 PayDash Inc. Aziz Nurfalah.
    </footer>
  </div>
</template>

