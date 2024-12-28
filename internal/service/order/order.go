package order

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/Junx27/inventory-golang/pkg/database"
)

type Order struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	OrderDate time.Time `json:"order_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllOrders(ctx context.Context) ([]Order, error) {
	query := `SELECT id, product_id, quantity, order_date, created_at, updated_at FROM orders`
	rows, err := database.DB.Query(ctx, query)
	if err != nil {
		log.Printf("Database query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.ProductID, &order.Quantity, &order.OrderDate, &order.CreatedAt, &order.UpdatedAt); err != nil {
			log.Printf("Row scan failed: %v", err)
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrderByID(ctx context.Context, id string) (*Order, error) {
	query := `SELECT id, product_id, quantity, order_date, created_at, updated_at FROM orders WHERE id = $1`
	row := database.DB.QueryRow(ctx, query, id)

	var order Order
	if err := row.Scan(&order.ID, &order.ProductID, &order.Quantity, &order.OrderDate, &order.CreatedAt, &order.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}

func CreateOrder(ctx context.Context, order *Order) error {
	query := `INSERT INTO orders (product_id, quantity, order_date) VALUES ($1, $2, NOW()) RETURNING id, created_at, updated_at`
	err := database.DB.QueryRow(ctx, query, order.ProductID, order.Quantity).Scan(&order.ID, &order.CreatedAt, &order.UpdatedAt)
	return err
}

func UpdateOrder(ctx context.Context, id int, order *Order) error {
	query := `UPDATE orders SET product_id = $1, quantity = $2, updated_at = NOW() WHERE id = $3`
	_, err := database.DB.Exec(ctx, query, order.ProductID, order.Quantity, id)
	return err
}

func DeleteOrder(ctx context.Context, id string) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := database.DB.Exec(ctx, query, id)
	return err
}
