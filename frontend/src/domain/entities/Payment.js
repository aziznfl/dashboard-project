export class Payment {
  constructor({ id, merchant, status, amount, created_at }) {
    this.id = id;
    this.merchant = merchant;
    this.status = status; // completed, processing, failed
    this.amount = amount;
    this.createdAt = new Date(created_at);
  }

  get formattedAmount() {
    return this.amount.toCurrency();
  }

  get statusColor() {
    switch (this.status) {
      case 'completed': return 'text-emerald-400 bg-emerald-400/10 border-emerald-400/20';
      case 'processing': return 'text-amber-400 bg-amber-400/10 border-amber-400/20';
      case 'failed': return 'text-rose-400 bg-rose-400/10 border-rose-400/20';
      default: return 'text-surface-400 bg-surface-400/10 border-surface-400/20';
    }
  }

  get formattedDate() {
    return this.createdAt.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  }
}
