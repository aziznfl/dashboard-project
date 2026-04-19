package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
)

type PaymentRepository interface {
	List(ctx context.Context, filter PaymentFilter) ([]*entity.Payment, error)
	Count(ctx context.Context, filter PaymentFilter) (int64, error)
}

type PaymentFilter struct {
	ID       *string
	Status   *string
	Merchant *string
	Amount   *int64
	Sort     *string
	LastID   *string
	Page     int
	Limit    int
}

type Payment struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *Payment {
	return &Payment{db: db}
}

func (r *Payment) List(ctx context.Context, filter PaymentFilter) ([]*entity.Payment, error) {
	query := `SELECT id, amount, merchant, status, created_at FROM payments WHERE 1=1`
	var args []interface{}

	if filter.ID != nil && *filter.ID != "" {
		query += ` AND id = ?`
		args = append(args, *filter.ID)
	}
	if filter.Status != nil && *filter.Status != "" {
		query += ` AND status = ?`
		args = append(args, *filter.Status)
	}
	if filter.Merchant != nil && *filter.Merchant != "" {
		query += ` AND merchant = ?`
		args = append(args, *filter.Merchant)
	}
	if filter.Amount != nil {
		query += ` AND amount = ?`
		args = append(args, *filter.Amount)
	}

	if filter.LastID != nil && *filter.LastID != "" {
		query += ` AND id > ?`
		args = append(args, *filter.LastID)
	}

	if filter.Sort != nil && *filter.Sort != "" {
		switch *filter.Sort {
		case "amount":
			query += ` ORDER BY amount ASC`
		case "-amount":
			query += ` ORDER BY amount DESC`
		case "merchant":
			query += ` ORDER BY merchant ASC`
		case "-merchant":
			query += ` ORDER BY merchant DESC`
		case "created_at":
			query += ` ORDER BY created_at ASC`
		case "-created_at":
			query += ` ORDER BY created_at DESC`
		default:
			query += ` ORDER BY created_at DESC`
		}
	} else {
		query += ` ORDER BY created_at DESC`
	}

	if filter.Limit > 0 {
		query += ` LIMIT ?`
		args = append(args, filter.Limit)
		if filter.Page > 1 {
			offset := (filter.Page - 1) * filter.Limit
			query += ` OFFSET ?`
			args = append(args, offset)
		}
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, entity.WrapError(err, entity.ErrorCodeInternal, "db error")
	}
	defer rows.Close()

	var payments []*entity.Payment
	for rows.Next() {
		var p entity.Payment
		var createdAt time.Time
		if err := rows.Scan(&p.ID, &p.Amount, &p.Merchant, &p.Status, &createdAt); err != nil {
			return nil, entity.WrapError(err, entity.ErrorCodeInternal, "scan error")
		}
		p.CreatedAt = createdAt
		payments = append(payments, &p)
	}

	return payments, nil
}
func (r *Payment) Count(ctx context.Context, filter PaymentFilter) (int64, error) {
	query := `SELECT COUNT(*) FROM payments WHERE 1=1`
	var args []interface{}

	if filter.ID != nil && *filter.ID != "" {
		query += ` AND id = ?`
		args = append(args, *filter.ID)
	}
	if filter.Status != nil && *filter.Status != "" {
		query += ` AND status = ?`
		args = append(args, *filter.Status)
	}
	if filter.Merchant != nil && *filter.Merchant != "" {
		query += ` AND merchant = ?`
		args = append(args, *filter.Merchant)
	}
	if filter.Amount != nil {
		query += ` AND amount = ?`
		args = append(args, *filter.Amount)
	}

	var count int64
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, entity.WrapError(err, entity.ErrorCodeInternal, "db error")
	}

	return count, nil
}
