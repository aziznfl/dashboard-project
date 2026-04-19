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
      id: ''
    }
  }),
  
  actions: {
    async fetchPayments() {
      this.loading = true;
      this.error = null;
      try {
        const payments = await PaymentService.fetchPayments(this.filters);
        this.payments = payments;
      } catch (err) {
        this.error = err;
      } finally {
        this.loading = false;
      }
    },
    
    setFilters(newFilters) {
      this.filters = { ...this.filters, ...newFilters };
      return this.fetchPayments();
    }
  }
});
