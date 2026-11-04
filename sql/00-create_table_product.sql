CREATE TABLE product (
	id SERIAL PRIMARY KEY,
	product_name VARCHAR(100) NOT NULL,
	price NUMERIC(10, 2) NOT NULL
);

SELECT * FROM product p;