import { Payment } from '@/domain/entities/Payment';

export class PaymentDTO {
  static toEntity(data) {
    return new Payment({
      id: data.id,
      merchant: data.merchant,
      status: data.status,
      amount: data.amount,
      created_at: data.created_at,
    });
  }

  static toEntities(dataList) {
    if (!dataList || !Array.isArray(dataList)) return [];
    return dataList.map(item => this.toEntity(item));
  }
}
