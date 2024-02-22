package database

import (
	"database/sql"

	"github.com/carloseduribeiro/order-service-clean-arch/internal/domain/entity"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) FindAll() ([]entity.Order, error) {
	rows, err := r.db.Query("SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := make([]entity.Order, 0)
	for rows.Next() {
		var id string
		var price, tax, finalPrice float64
		if err := rows.Scan(&id, &price, &tax, &finalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, entity.Order{ID: id, Price: price, Tax: tax, FinalPrice: finalPrice})
	}
	return orders, nil
}
