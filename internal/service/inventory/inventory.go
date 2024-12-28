package inventory

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Junx27/inventory-golang/pkg/database"
)

type Inventory struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllInventories(ctx context.Context) ([]Inventory, error) {
	query := `SELECT id, product_id, quantity, location, created_at, updated_at FROM inventory`
	rows, err := database.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventories []Inventory
	for rows.Next() {
		var inventory Inventory
		if err := rows.Scan(&inventory.ID, &inventory.ProductID, &inventory.Quantity, &inventory.Location, &inventory.CreatedAt, &inventory.UpdatedAt); err != nil {
			return nil, err
		}
		inventories = append(inventories, inventory)
	}

	return inventories, nil
}

func GetInventoryByID(ctx context.Context, id string) (*Inventory, error) {
	query := `SELECT id, product_id, quantity, location, created_at, updated_at FROM inventory WHERE id = $1`
	row := database.DB.QueryRow(ctx, query, id)

	var inventory Inventory
	if err := row.Scan(&inventory.ID, &inventory.ProductID, &inventory.Quantity, &inventory.Location, &inventory.CreatedAt, &inventory.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &inventory, nil
}

func CreateInventory(ctx context.Context, inventory *Inventory) error {
	query := `INSERT INTO inventory (product_id, quantity, location) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	err := database.DB.QueryRow(ctx, query, inventory.ProductID, inventory.Quantity, inventory.Location).Scan(&inventory.ID, &inventory.CreatedAt, &inventory.UpdatedAt)
	return err
}

func UpdateInventory(ctx context.Context, id string, inventory *Inventory) error {
	query := `UPDATE inventory SET product_id = $1, quantity = $2, location = $3, updated_at = NOW() WHERE id = $4`
	_, err := database.DB.Exec(ctx, query, inventory.ProductID, inventory.Quantity, inventory.Location, id)
	return err
}

func DeleteInventory(ctx context.Context, id string) error {
	checkQuery := "SELECT COUNT(*) FROM inventory WHERE id = $1"
	var count int
	err := database.DB.QueryRow(ctx, checkQuery, id).Scan(&count)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if count == 0 {
		return fmt.Errorf("inventory with ID %s not found", id)
	}

	deleteQuery := "DELETE FROM inventory WHERE id = $1"
	_, err = database.DB.Exec(ctx, deleteQuery, id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil

}
