# Inventory System

langkah-langkah menjalankan proyek, salin repository ke local komputer dengan cara:

```bash
git@github.com:Junx27/inventory-golang.git
cd inventory-golang
```

Setelah repository berhasil lakukan perintah berikut untuk menjalankan proyek di local komputer:

```bash
docker compose up
```

atau

```bash
docker compose up --build
```

Untuk menghentikan program dapat melakukan dengan menjalankan perintah sebagai berikut:

```bash
docker compose down
```

# Implementasi API Route

## Route Product

```go
r.group.GET("/products", r.handler.GetAllProductsHandler)
r.group.GET("/products/:id", r.handler.GetProductByIDHandler)
r.group.POST("/products", r.handler.StoreProductHandler)
r.group.PUT("/products/:id", r.handler.UpdateProductHandler)
r.group.DELETE("/products/:id", r.handler.DeleteProductHandler)
```

## Route Inventory

```go
r.group.GET("/inventory", r.handler.GetAllInventoriesHandler)
r.group.GET("/inventory/:id", r.handler.GetInventoryByIDHandler)
r.group.POST("/inventory", r.handler.StoreInventoryHandler)
r.group.PUT("/inventory/:id", r.handler.UpdateInventoryHandler)
r.group.DELETE("/inventory/:id", r.handler.DeleteInventoryHandler)
```

## Route Order

```go
r.group.GET("/orders", r.handler.GetAllOrdersHandler)
r.group.GET("/orders/:id", r.handler.GetOrderByIDHandler)
r.group.POST("orders", r.handler.CreateOrderHandler)
r.group.PUT("/orders/:id", r.handler.UpdateOrderHandler)
r.group.DELETE("/orders/:id", r.handler.DeleteOrderHandler)
```

# Dokumentasi API

## 1. Product

### 1.1 Add Product

# Input Data Produk

Untuk menambahkan data produk baru, Anda dapat menggunakan format berikut:

| **Field**     | **Description**                     | **Example**                        |
| ------------- | ----------------------------------- | ---------------------------------- |
| `name`        | Nama produk                         | "Laptop Gaming"                    |
| `description` | Deskripsi produk                    | "Laptop dengan spesifikasi tinggi" |
| `price`       | Harga produk (dalam IDR)            | 15000000                           |
| `category`    | Kategori produk                     | "Elektronik"                       |
| `image_path`  | Path atau lokasi file gambar produk | "/images/laptop.jpg"               |

### **Contoh Input Data Produk**

Untuk menginput data produk menggunakan API atau form, pastikan Anda mengisi field sesuai dengan format di atas.

#### Contoh JSON untuk Input Produk:

```json
{
  "name": "Laptop Gaming",
  "description": "Laptop dengan spesifikasi tinggi",
  "price": 15000000,
  "category": "Elektronik",
  "image_path": "/images/laptop.jpg"
}
```
