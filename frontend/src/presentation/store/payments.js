import { defineStore } from 'pinia';
import { PaymentService } from '@/application/services/PaymentService';

export const usePaymentStore = defineStore('payments', {
  state: () => ({
    payments: [],
    loading: false,
    error: null,
    filters: {
      sort: '-created_at',
      status: '',
      id: '',
      page: 1,
      limit: 10
    },
    meta: {
      total: 0,
      limit: 10,
      page: 1,
      totalPages: 0,
      last_id: null
    }
  }),
  
  actions: {
    async fetchPayments() {
      this.loading = true;
      this.error = null;
      try {
        const { data, meta } = await PaymentService.fetchPayments(this.filters);
        this.payments = data;
        // Map camelCase for frontend if needed, but the API might return snake_case in meta
        this.meta = {
          total: meta.total,
          limit: meta.limit,
          page: meta.page,
          totalPages: meta.total_pages,
          lastId: meta.last_id
        };
      } catch (err) {
        this.error = err;
      } finally {
        this.loading = false;
      }
    },
    
    setPage(page) {
      this.filters.page = page;
      return this.fetchPayments();
    },

    setFilters(newFilters) {
      // If setting filters, reset page to 1
      this.filters = { ...this.filters, ...newFilters, page: 1 };
      return this.fetchPayments();
    }
  }
});
