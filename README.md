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

**Endpoint**: `POST /products`
**Deskripsi**: Menambahkan produk baru ke dalam sistem.
**Request Body**:

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

### 1.2 Get All Product

**Endpoint**: `GET /products`
**Deskripsi**: Melihat semua daftar produk pada sistem
**Output JSON**:

```json
{
  "data": [
    {
      "id": 2,
      "name": "Chair",
      "description": "Ergonomic office chair",
      "price": 2000,
      "category": "Furniture",
      "image_path": "/images/chair.jpg",
      "created_at": "2024-12-28T06:33:30.121135Z",
      "updated_at": "2024-12-28T06:33:30.121135Z",
      "orders": [
        {
          "id": 2,
          "quantity": 1,
          "order_date": "2024-12-26T14:30:00Z",
          "created_at": "2024-12-28T06:33:30.121135Z",
          "updated_at": "2024-12-28T06:33:30.121135Z"
        }
      ]
    },
    {
      "id": 3,
      "name": "Book",
      "description": "Programming book",
      "price": 5000,
      "category": "Books",
      "image_path": "/images/book.jpg",
      "created_at": "2024-12-28T06:33:30.121135Z",
      "updated_at": "2024-12-28T06:33:30.121135Z",
      "inventory": {
        "id": 3,
        "quantity": 50,
        "location": "Warehouse C",
        "created_at": "2024-12-28T06:33:30.121135Z",
        "updated_at": "2024-12-28T06:33:30.121135Z"
      },
      "orders": [
        {
          "id": 3,
          "quantity": 5,
          "order_date": "2024-12-27T09:00:00Z",
          "created_at": "2024-12-28T06:33:30.121135Z",
          "updated_at": "2024-12-28T06:33:30.121135Z"
        }
      ]
    },
    {
      "id": 1,
      "name": "Laptop",
      "description": "High performance laptop",
      "price": 1500,
      "category": "Electronics",
      "image_path": "/images/laptop.jpg",
      "created_at": "2024-12-28T06:33:30.121135Z",
      "updated_at": "2024-12-28T06:33:30.121135Z",
      "inventory": {
        "id": 1,
        "quantity": 15,
        "location": "Warehouse B",
        "created_at": "2024-12-28T06:33:30.121135Z",
        "updated_at": "2024-12-28T06:35:17.612894Z"
      },
      "orders": [
        {
          "id": 4,
          "quantity": 5,
          "order_date": "2024-12-28T07:10:21.448878Z",
          "created_at": "2024-12-28T07:10:21.448878Z",
          "updated_at": "2024-12-28T07:10:21.448878Z"
        }
      ]
    }
  ]
}
```

### 1.3 Get Product By ID

**Endpoint**: `GET /products/:id`
**Deskripsi**: Melihat daftar produk berdasarkan ID pada sistem
**Output JSON**:

```json
{
  "data": {
    "id": 1,
    "name": "Laptop",
    "description": "High performance laptop",
    "price": 1500,
    "category": "Electronics",
    "image_path": "/images/laptop.jpg",
    "created_at": "2024-12-28T06:33:30.121135Z",
    "updated_at": "2024-12-28T06:33:30.121135Z",
    "inventory": {
      "id": 1,
      "quantity": 15,
      "location": "Warehouse B",
      "created_at": "2024-12-28T06:33:30.121135Z",
      "updated_at": "2024-12-28T06:35:17.612894Z"
    },
    "orders": [
      {
        "id": 4,
        "quantity": 5,
        "order_date": "2024-12-28T07:10:21.448878Z",
        "created_at": "2024-12-28T07:10:21.448878Z",
        "updated_at": "2024-12-28T07:10:21.448878Z"
      }
    ]
  }
}
```

### 1.3 Update Product By ID

**Endpoint**: `PUT /products/:id`
**Deskripsi**: Memperbarui informasi produk berdasarkan ID produk
**Request BODY JSON**:

```json
{
  "name": "Laptop Gaming Updated",
  "description": "Laptop dengan spesifikasi lebih tinggi",
  "price": 17000000,
  "category": "Elektronik",
  "image_path": "/images/laptop_updated.jpg"
}
```

**Response**

```json
{
  "message": "product updated successfully"
}
```

### 1.3 Delete Product By ID

**Endpoint**: `DELETE /products/:id`
**Deskripsi**: Menghapus daftar produk berdasarkan ID pada sistem
**Response**:

```json
{
  "message": "product deleted successfully"
}
```
