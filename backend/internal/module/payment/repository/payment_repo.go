package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
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
	where, args := r.buildWhere(filter)
	query := "SELECT id, amount, merchant, status, created_at FROM payments WHERE 1=1" + where
	query = r.applySort(query, filter.Sort)
	query, args = r.applyLimit(query, args, filter)

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
	where, args := r.buildWhere(filter)
	query := "SELECT COUNT(*) FROM payments WHERE 1=1" + where

	var count int64
	if err := r.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, entity.WrapError(err, entity.ErrorCodeInternal, "db error")
	}

	return count, nil
}

func (r *Payment) buildWhere(f PaymentFilter) (string, []interface{}) {
	var where strings.Builder
	var args []interface{}

	if f.ID != nil && *f.ID != "" {
		where.WriteString(" AND id = ?")
		args = append(args, *f.ID)
	}
	if f.Status != nil && *f.Status != "" {
		where.WriteString(" AND status = ?")
		args = append(args, *f.Status)
	}
	if f.Merchant != nil && *f.Merchant != "" {
		where.WriteString(" AND merchant = ?")
		args = append(args, *f.Merchant)
	}
	if f.Amount != nil {
		where.WriteString(" AND amount = ?")
		args = append(args, *f.Amount)
	}
	if f.LastID != nil && *f.LastID != "" {
		where.WriteString(" AND id > ?")
		args = append(args, *f.LastID)
	}

	return where.String(), args
}

func (r *Payment) applySort(query string, sort *string) string {
	field, order := "created_at", "DESC"

	if sort != nil && *sort != "" {
		s := *sort
		if strings.HasPrefix(s, "-") {
			field, order = s[1:], "DESC"
		} else {
			field, order = s, "ASC"
		}

		allowed := map[string]bool{"id": true, "amount": true, "merchant": true, "status": true, "created_at": true}
		if !allowed[field] {
			field, order = "created_at", "DESC"
		}
	}

	return fmt.Sprintf("%s ORDER BY %s %s", query, field, order)
}

func (r *Payment) applyLimit(query string, args []interface{}, f PaymentFilter) (string, []interface{}) {
	if f.Limit <= 0 {
		return query, args
	}

	query += " LIMIT ?"
	args = append(args, f.Limit)

	if f.Page > 1 {
		query += " OFFSET ?"
		args = append(args, (f.Page-1)*f.Limit)
	}

	return query, args
}
