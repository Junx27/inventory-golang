
INSERT INTO products (name, description, price, category, image_path) VALUES
('Laptop', 'High performance laptop', 1500, 'Electronics', '/images/laptop.jpg'),
('Chair', 'Ergonomic office chair', 2000, 'Furniture', '/images/chair.jpg'),
('Book', 'Programming book', 5000, 'Books', '/images/book.jpg');


INSERT INTO inventory (product_id, quantity, location) VALUES
(1, 10, 'Warehouse A'),
(2, 20, 'Warehouse B'),
(3, 50, 'Warehouse C');

INSERT INTO orders (product_id, quantity, order_date) VALUES
(1, 2, '2024-12-25 10:00:00'),
(2, 1, '2024-12-26 14:30:00'),
(3, 5, '2024-12-27 09:00:00');
