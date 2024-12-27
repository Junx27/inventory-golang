-- Membuat tabel products
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price INT NOT NULL,
    category VARCHAR(100),
    image_path TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Membuat tabel inventory dengan foreign key yang memiliki ON DELETE CASCADE
CREATE TABLE IF NOT EXISTS inventory (
    id SERIAL PRIMARY KEY,
    product_id INT,
    quantity INT NOT NULL,
    location VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT fk_product FOREIGN KEY (product_id) 
        REFERENCES products(id) 
        ON DELETE CASCADE -- Menambahkan CASCADE delete di sini
);

-- Membuat tabel orders dengan foreign key yang memiliki ON DELETE CASCADE
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    product_id INT,
    quantity INT NOT NULL,
    order_date TIMESTAMPTZ DEFAULT NOW(),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT fk_product_order FOREIGN KEY (product_id) 
        REFERENCES products(id) 
        ON DELETE CASCADE -- Menambahkan CASCADE delete di sini
);
