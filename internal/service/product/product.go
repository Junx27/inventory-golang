package product

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Junx27/inventory-golang/pkg/database"
	"github.com/jackc/pgx/v5"
)

type Product struct {
	ID          int        `json:"id" db:"id"`
	Name        string     `form:"name" json:"name"`
	Description string     `form:"description" json:"description"`
	Price       float64    `form:"price" json:"price"`
	Category    string     `form:"category" json:"category"`
	ImagePath   string     `json:"image_path" db:"image_path"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	Inventory   *Inventory `json:"inventory,omitempty"`
	Orders      []Order    `json:"orders,omitempty"`
}

type Inventory struct {
	ID        int       `json:"id"`
	Quantity  int       `json:"quantity"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	ID        int       `json:"id"`
	Quantity  int       `json:"quantity"`
	OrderDate time.Time `json:"order_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductNotFound struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func GetAllProducts(ctx context.Context) ([]Product, error) {
	query := `
		SELECT 
			p.id AS product_id, p.name, p.description, p.price, p.category, p.image_path, 
			p.created_at AS product_created_at, p.updated_at AS product_updated_at,
			COALESCE(i.id, 0) AS inventory_id, 
			COALESCE(i.quantity, 0) AS inventory_quantity, 
			COALESCE(i.location, '') AS inventory_location, 
			COALESCE(i.created_at, '1970-01-01 00:00:00') AS inventory_created_at, 
			COALESCE(i.updated_at, '1970-01-01 00:00:00') AS inventory_updated_at,
			COALESCE(o.id, 0) AS order_id, 
			COALESCE(o.quantity, 0) AS order_quantity, 
			COALESCE(o.order_date, NOW()) AS order_date, 
			COALESCE(o.created_at, '1970-01-01 00:00:00') AS order_created_at, 
			COALESCE(o.updated_at, '1970-01-01 00:00:00') AS order_updated_at
		FROM products p
		LEFT JOIN inventory i ON p.id = i.product_id
		LEFT JOIN orders o ON p.id = o.product_id
	`

	rows, err := database.DB.Query(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		var inv Inventory
		var ord Order

		err := rows.Scan(
			&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.ImagePath, &product.CreatedAt, &product.UpdatedAt,
			&inv.ID, &inv.Quantity, &inv.Location, &inv.CreatedAt, &inv.UpdatedAt,
			&ord.ID, &ord.Quantity, &ord.OrderDate, &ord.CreatedAt, &ord.UpdatedAt,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		if inv.ID != 0 {
			product.Inventory = &inv
		}

		if ord.ID != 0 {
			product.Orders = append(product.Orders, ord)
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return products, nil
}

func GetProductByID(ctx context.Context, id string) (Product, *ProductNotFound, error) {
	query := `
		SELECT 
			p.id, p.name, p.description, p.price, p.category, p.image_path, p.created_at, p.updated_at,
			COALESCE(i.id, 0) AS inventory_id, 
			COALESCE(i.quantity, 0) AS inventory_quantity, 
			COALESCE(i.location, '') AS inventory_location, 
			COALESCE(i.created_at, '1970-01-01 00:00:00') AS inventory_created_at, 
			COALESCE(i.updated_at, '1970-01-01 00:00:00') AS inventory_updated_at,
			COALESCE(o.id, 0) AS order_id, 
			COALESCE(o.quantity, 0) AS order_quantity, 
			COALESCE(o.order_date, NOW()) AS order_date, 
			COALESCE(o.created_at, '1970-01-01 00:00:00') AS order_created_at, 
			COALESCE(o.updated_at, '1970-01-01 00:00:00') AS order_updated_at
		FROM products p
		LEFT JOIN inventory i ON p.id = i.product_id
		LEFT JOIN orders o ON p.id = o.product_id
		WHERE p.id = $1
	`

	rows, err := database.DB.Query(ctx, query, id)
	if err != nil {
		log.Println(err.Error())
		return Product{}, nil, fmt.Errorf("failed to query product with ID %s: %v", id, err)
	}
	defer rows.Close()

	var product Product
	if !rows.Next() {
		notFound := &ProductNotFound{
			ID:      id,
			Message: fmt.Sprintf("Product with ID %s not found", id),
		}
		return product, notFound, nil
	}

	var inv Inventory
	var ord Order

	err = rows.Scan(
		&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.ImagePath, &product.CreatedAt, &product.UpdatedAt,
		&inv.ID, &inv.Quantity, &inv.Location, &inv.CreatedAt, &inv.UpdatedAt,
		&ord.ID, &ord.Quantity, &ord.OrderDate, &ord.CreatedAt, &ord.UpdatedAt,
	)
	if err != nil {
		log.Println(err.Error())
		return Product{}, nil, err
	}

	if inv.ID != 0 {
		product.Inventory = &inv
	}

	if ord.ID != 0 {
		product.Orders = append(product.Orders, ord)
	}

	if err := rows.Err(); err != nil {
		log.Println(err.Error())
		return Product{}, nil, err
	}

	return product, nil, nil
}

func StoreProduct(ctx context.Context, req *Product) error {
	query := `INSERT INTO products (name, description, price, category, image_path)
              VALUES (@name, @description, @price, @category, @image_path)
			  RETURNING id`

	args := pgx.NamedArgs{
		"name":        req.Name,
		"description": req.Description,
		"price":       req.Price,
		"category":    req.Category,
		"image_path":  req.ImagePath,
	}

	row := database.DB.QueryRow(ctx, query, args)
	if err := row.Scan(&req.ID); err != nil {
		log.Println("Error inserting product:", err)
		return err
	}
	return nil
}

func UpdateProduct(ctx context.Context, product Product) error {
	query := `UPDATE products 
			  SET name = @name, description = @description, price = @price, category = @category, 
				  image_path = @image_path, updated_at = @updated_at 
			  WHERE id = @id`
	args := pgx.NamedArgs{
		"id":          product.ID,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"category":    product.Category,
		"image_path":  product.ImagePath,
		"updated_at":  product.UpdatedAt,
	}

	_, err := database.DB.Exec(ctx, query, args)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func DeleteProduct(ctx context.Context, id string) error {

	checkQuery := "SELECT COUNT(*) FROM products WHERE id = $1"
	var count int
	err := database.DB.QueryRow(ctx, checkQuery, id).Scan(&count)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if count == 0 {
		return fmt.Errorf("product with ID %s not found", id)
	}

	deleteQuery := "DELETE FROM products WHERE id = $1"
	_, err = database.DB.Exec(ctx, deleteQuery, id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
