import api from '../services/api';
import { PaymentRepository } from '@/domain/repositories/PaymentRepository';
import { PaymentDTO } from '../dto/PaymentDTO';

export class PaymentRepositoryImpl extends PaymentRepository {
  async getPayments({ sort, status, id } = {}) {
    const params = {};
    if (sort) params.sort = sort;
    if (status) params.status = status;
    if (id) params.id = id;

    const response = await api.get('/dashboard/v1/payments', { params });
    // The response schema from openapi.yaml: { payments: [Payment] }
    return PaymentDTO.toEntities(response.data.payments);
  }
}
