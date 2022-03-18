create table products(
    id serial primary key,
    name varchar,
    editorial_name varchar
);

INSERT INTO products(name, editorial_name) VALUES
('iphone x', 'iphone X 128gb'),
('samsung s22', 'samsung s22 ultra 512gb');