import api from '../services/api';
import { PaymentRepository } from '@/domain/repositories/PaymentRepository';
import { PaymentDTO } from '../dto/PaymentDTO';

export class PaymentRepositoryImpl extends PaymentRepository {
  async getPayments(filters = {}) {
    const params = {};
    Object.keys(filters).forEach(key => {
      const val = filters[key];
      if (val !== undefined && val !== null && val !== '') {
        params[key] = val;
      }
    });

    // Ensure page and limit have defaults if not set
    if (!params.page) params.page = 1;
    if (!params.limit) params.limit = 10;

    const response = await api.get('/dashboard/v1/payments', { params });
    // The response schema from openapi.yaml: { data: [Payment], meta: PaginationMeta }
    return {
      data: PaymentDTO.toEntities(response.data.data),
      meta: response.data.meta
    };
  }
}
