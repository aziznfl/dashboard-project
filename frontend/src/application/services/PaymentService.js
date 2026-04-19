import { PaymentRepositoryImpl } from '@/data/repositories/PaymentRepositoryImpl';

const paymentRepository = new PaymentRepositoryImpl();

export const PaymentService = {
  async fetchPayments(filters = {}) {
    try {
      const payments = await paymentRepository.getPayments(filters);
      return payments;
    } catch (error) {
      throw error.response?.data?.message || 'Failed to fetch payments';
    }
  }
};
