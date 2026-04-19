export const toCurrency = (value) => {
  if (value === null || value === undefined) return 'Rp 0';
  
  const num = typeof value === 'string' ? parseFloat(value) : value;
  
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(num);
};

// Extend Number prototype for direct access as requested
if (!Number.prototype.toCurrency) {
  Object.defineProperty(Number.prototype, 'toCurrency', {
    value: function() {
      return toCurrency(this);
    },
    enumerable: false,
    configurable: true,
    writable: true
  });
}

// Optional: Extend String prototype for convenience
if (!String.prototype.toCurrency) {
  Object.defineProperty(String.prototype, 'toCurrency', {
    value: function() {
      return toCurrency(this);
    },
    enumerable: false,
    configurable: true,
    writable: true
  });
}
