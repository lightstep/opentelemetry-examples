CREATE DATABASE sample;
USE sample;
SET GLOBAL ignore_foreign_keys = true;

/* The memsql-exporter process (or simply “the exporter”) collects data about a running cluster. The user that starts
   the exporter (other than the SingleStore DB root user) must have the following permissions at a minimum:
*/

#GRANT CLUSTER on *.* to <user>
#GRANT SHOW METADATA on *.* to <user>
#GRANT SELECT on *.* to <user>

/* HTTP */
SET GLOBAL exporter_user = root;
SET GLOBAL exporter_password = 'password_here';
SET GLOBAL exporter_port= 9104;

/* HTTPS */
#SET GLOBAL exporter_user = root;
#SET GLOBAL exporter_password = '<secure-password>';
#SET GLOBAL exporter_use_https= true;
#SET GLOBAL exporter_ssl_cert = '/path/to/server-cert.pem';
#SET GLOBAL exporter_ssl_key = '/path/to/server-key.pem';
#SET GLOBAL exporter_ssl_key_passphrase= '<passphrase>';

/* Use an engine variable to stop the exporter process by setting the port to 0. */
# SET GLOBAL exporter_port = 0;

/* Sample Data */
CREATE TABLE customers (
                           customer_id INT PRIMARY KEY,
                           first_name VARCHAR(50),
                           last_name VARCHAR(50),
                           email VARCHAR(100),
                           phone_number VARCHAR(15),
                           shipping_address VARCHAR(200)
);

INSERT INTO customers (customer_id, first_name, last_name, email, phone_number, shipping_address)
VALUES
    (1, 'John', 'Doe', 'john.doe@example.com', '555-555-1212', '123 Main St, Anytown USA'),
    (2, 'Jane', 'Smith', 'jane.smith@example.com', '555-555-1213', '456 Oak Ave, Anytown USA'),
    (3, 'Bob', 'Johnson', 'bob.johnson@example.com', '555-555-1214', '789 Maple Rd, Anytown USA');

CREATE TABLE orders (
                        order_id INT PRIMARY KEY,
                        customer_id INT,
                        order_date DATETIME,
                        total DECIMAL(10,2),
                        FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

INSERT INTO orders (order_id, customer_id, order_date, total)
VALUES
    (1, 1, '2023-05-01 12:34:56', 100.00),
    (2, 2, '2023-05-02 01:23:45', 50.00),
    (3, 1, '2023-05-03 11:22:33', 75.00),
    (4, 3, '2023-05-04 09:08:07', 125.00);

CREATE TABLE order_items (
                             order_id INT,
                             product_id INT,
                             quantity INT,
                             price DECIMAL(10,2),
                             PRIMARY KEY (order_id, product_id),
                             FOREIGN KEY (order_id) REFERENCES orders(order_id)
);

INSERT INTO order_items (order_id, product_id, quantity, price)
VALUES
    (1, 1, 2, 25.00),
    (1, 2, 1, 50.00),
    (2, 3, 1, 50.00),
    (3, 1, 3, 25.00),
    (4, 2, 2, 50.00),
    (4, 3, 1, 25.00);

CREATE TABLE products (
                          product_id INT PRIMARY KEY,
                          name VARCHAR(100),
                          description VARCHAR(500),
                          price DECIMAL(10,2)
);

INSERT INTO products (product_id, name, description, price)
VALUES
    (1, 'Product A', 'A great product', 25.00),
    (2, 'Product B', 'Another great product', 50.00),
    (3, 'Product C', 'Yet another great product', 25.00);
