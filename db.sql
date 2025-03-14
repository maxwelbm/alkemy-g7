DROP DATABASE IF EXISTS `meli_fresh`;

CREATE DATABASE `meli_fresh`;

USE `meli_fresh`;

-- table `warehouses`
CREATE TABLE `warehouses` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `warehouse_code` varchar(25) NOT NULL,
    `address` varchar(255) NOT NULL,
    `telephone` varchar(15) NOT NULL,
    `minimum_capacity` int NOT NULL,
    `minimum_temperature` float NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `product_type`
CREATE TABLE `product_type`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `type_name` varchar(100) NOT NULL,
    PRIMARY KEY(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `employees`
CREATE TABLE `employees` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `card_number_id` varchar(25) NOT NULL,
    `first_name` varchar(50) NOT NULL,
    `last_name` varchar(50) NOT NULL,
    `warehouse_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE(`card_number_id`),
    FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses`(`id`)  -- Corrigido para 'warehouses'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `sections`
CREATE TABLE `sections` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `section_number` varchar(255) NOT NULL,
    `current_temperature` float NOT NULL,
    `minimum_temperature` float NOT NULL,
    `current_capacity` int NOT NULL,
    `minimum_capacity` int NOT NULL,
    `maximum_capacity` int NOT NULL,
    `warehouse_id` int(11) NOT NULL,
    `product_type_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE(`section_number`),
    FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses`(`id`),  -- Corrigido para 'warehouses'
    FOREIGN KEY (`product_type_id`) REFERENCES `product_type`(`id`)  -- Corrigido para 'product_type'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `buyers`
CREATE TABLE `buyers` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `card_number_id` varchar(25) NOT NULL,
    `first_name` varchar(50) NOT NULL,
    `last_name` varchar(50) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE(`card_number_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `countries`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `country_name` varchar(255),
    PRIMARY KEY(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `provinces`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `province_name` varchar(255),
    `id_country_fk` int(11),
    PRIMARY KEY(`id`),
    FOREIGN KEY (`id_country_fk`) REFERENCES `countries`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `locality`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `locality_name` varchar(255),
    `province_name` varchar(255),
    `country_name` varchar(255),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `sellers`
CREATE TABLE `sellers` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `cid` int(11) NOT NULL,
    `company_name` varchar(255) NOT NULL,
    `address` varchar(255) NOT NULL,
    `telephone` varchar(15) NOT NULL,
    `locality_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE(`cid`),
    FOREIGN KEY (`locality_id`) REFERENCES `locality`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `products`
CREATE TABLE `products` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `product_code` varchar(25) NOT NULL,
    `description` text NOT NULL,
    `height` float NOT NULL,
    `length` float NOT NULL,
    `width` float NOT NULL,
    `net_weight` float NOT NULL,
    `expiration_rate` float NOT NULL,
    `freezing_rate` float NOT NULL,
    `recommended_freezing_temperature` float NOT NULL,
    `seller_id` int(11) NOT NULL,
    `product_type_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE(`product_code`),
    FOREIGN KEY (`product_type_id`) REFERENCES `product_type`(`id`),
    FOREIGN KEY (`seller_id`) REFERENCES `sellers`(`id`)  -- Corrigido para 'sellers'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `carriers`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `cid` VARCHAR(100),
    `company_name` VARCHAR(100),
    `address` varchar(100),
    `telephone` varchar(20),
    `locality_id` int(11),
    PRIMARY KEY(`id`),
    UNIQUE(`cid`),
    FOREIGN KEY (`locality_id`) REFERENCES `locality`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `product_batches`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `batch_number` varchar(100),
    `current_quantity` int,
    `current_temperature` DECIMAL(19,2),
    `due_date` DATETIME(6),
    `initial_quantity` int,
    `manufacturing_date` DATETIME(6),
    `manufacturing_hour` int(11),
    `minimum_temperature` DECIMAL(19,2),
    `product_id` int(11),
    `section_id` int(11),
    PRIMARY KEY(`id`),
    UNIQUE(`batch_number`),
    FOREIGN KEY (`product_id`) REFERENCES `products`(`id`),  -- Corrigido para 'products'
    FOREIGN KEY (`section_id`) REFERENCES `sections`(`id`)  -- Corrigido para 'sections'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `product_records`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `last_update_date` DATETIME(6),
    `purchase_price` DECIMAL(19,2),
    `sale_price` DECIMAL(19,2),
    `product_id` int(11),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `products`(`id`)  -- Corrigido para 'products'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `inbound_orders`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `order_date` DATETIME(6),
    `order_number` varchar(255),
    `employee_id` int(11),
    `product_batch_id` int(11),
    `warehouse_id` int(11),
    PRIMARY KEY(`id`),
    UNIQUE(`order_number`),
    FOREIGN KEY (`employee_id`) REFERENCES `employees`(`id`),  -- Corrigido para 'employees'
    FOREIGN KEY (`product_batch_id`) REFERENCES `product_batches`(`id`),  -- Corrigido para 'product_batches'
    FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses`(`id`)  -- Corrigido para 'warehouses'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `purchase_orders`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `order_number` varchar(255),
    `order_date` DATETIME(6),
    `tracking_code` varchar(255),
    `buyer_id` int(11),
    `product_record_id` int(11),
    PRIMARY KEY(`id`),
    UNIQUE(`order_number`),
    FOREIGN KEY (`buyer_id`) REFERENCES `buyers`(`id`),  -- Corrigido para 'buyers'
    FOREIGN KEY (`product_record_id`) REFERENCES `product_records`(`id`)  -- Corrigido para 'product_records'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;


CREATE TABLE logs (
                      DROP DATABASE IF EXISTS `meli_fresh`;

CREATE DATABASE `meli_fresh`;

USE `meli_fresh`;

-- table `warehouses`
CREATE TABLE `warehouses` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `warehouse_code` varchar(25) NOT NULL,
                              `address` varchar(255) NOT NULL,
                              `telephone` varchar(15) NOT NULL,
                              `minimum_capacity` int NOT NULL,
                              `minimum_temperature` float NOT NULL,
                              PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `product_type`
CREATE TABLE `product_type`(
                               `id` int(11) NOT NULL AUTO_INCREMENT,
                               `type_name` varchar(100) NOT NULL,
                               PRIMARY KEY(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `employees`
CREATE TABLE `employees` (
                             `id` int(11) NOT NULL AUTO_INCREMENT,
                             `card_number_id` varchar(25) NOT NULL,
                             `first_name` varchar(50) NOT NULL,
                             `last_name` varchar(50) NOT NULL,
                             `warehouse_id` int(11) NOT NULL,
                             PRIMARY KEY (`id`),
                             UNIQUE(`card_number_id`),
                             FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses`(`id`)  -- Corrigido para 'warehouses'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `sections`
CREATE TABLE `sections` (
                            `id` int(11) NOT NULL AUTO_INCREMENT,
                            `section_number` varchar(255) NOT NULL,
                            `current_temperature` float NOT NULL,
                            `minimum_temperature` float NOT NULL,
                            `current_capacity` int NOT NULL,
                            `minimum_capacity` int NOT NULL,
                            `maximum_capacity` int NOT NULL,
                            `warehouse_id` int(11) NOT NULL,
                            `product_type_id` int(11) NOT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE(`section_number`),
                            FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses`(`id`),  -- Corrigido para 'warehouses'
                            FOREIGN KEY (`product_type_id`) REFERENCES `product_type`(`id`)  -- Corrigido para 'product_type'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `buyers`
CREATE TABLE `buyers` (
                          `id` int(11) NOT NULL AUTO_INCREMENT,
                          `card_number_id` varchar(25) NOT NULL,
                          `first_name` varchar(50) NOT NULL,
                          `last_name` varchar(50) NOT NULL,
                          PRIMARY KEY (`id`),
                          UNIQUE(`card_number_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `countries`(
                            `id` int(11) NOT NULL AUTO_INCREMENT,
                            `country_name` varchar(255),
                            PRIMARY KEY(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `provinces`(
                            `id` int(11) NOT NULL AUTO_INCREMENT,
                            `province_name` varchar(255),
                            `id_country_fk` int(11),
                            PRIMARY KEY(`id`),
                            FOREIGN KEY (`id_country_fk`) REFERENCES `countries`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `locality`(
                           `id` int(11) NOT NULL AUTO_INCREMENT,
                           `locality_name` varchar(255),
                           `province_name` varchar(255),
                           `country_name` varchar(255),
                           PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `sellers`
CREATE TABLE `sellers` (
                           `id` int(11) NOT NULL AUTO_INCREMENT,
                           `cid` int(11) NOT NULL,
                           `company_name` varchar(255) NOT NULL,
                           `address` varchar(255) NOT NULL,
                           `telephone` varchar(15) NOT NULL,
                           `locality_id` int(11) NOT NULL,
                           PRIMARY KEY (`id`),
                           UNIQUE(`cid`),
                           FOREIGN KEY (`locality_id`) REFERENCES `locality`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `products`
CREATE TABLE `products` (
                            `id` int(11) NOT NULL AUTO_INCREMENT,
                            `product_code` varchar(25) NOT NULL,
                            `description` text NOT NULL,
                            `height` float NOT NULL,
                            `length` float NOT NULL,
                            `width` float NOT NULL,
                            `net_weight` float NOT NULL,
                            `expiration_rate` float NOT NULL,
                            `freezing_rate` float NOT NULL,
                            `recommended_freezing_temperature` float NOT NULL,
                            `seller_id` int(11) NOT NULL,
                            `product_type_id` int(11) NOT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE(`product_code`),
                            FOREIGN KEY (`product_type_id`) REFERENCES `product_type`(`id`),
                            FOREIGN KEY (`seller_id`) REFERENCES `sellers`(`id`)  -- Corrigido para 'sellers'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `carriers`(
                           `id` int(11) NOT NULL AUTO_INCREMENT,
                           `cid` VARCHAR(100),
                           `company_name` VARCHAR(100),
                           `address` varchar(100),
                           `telephone` varchar(20),
                           `locality_id` int(11),
                           PRIMARY KEY(`id`),
                           UNIQUE(`cid`),
                           FOREIGN KEY (`locality_id`) REFERENCES `locality`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `product_batches`(
                                  `id` int(11) NOT NULL AUTO_INCREMENT,
                                  `batch_number` varchar(100),
                                  `current_quantity` int,
                                  `current_temperature` DECIMAL(19,2),
                                  `due_date` DATETIME(6),
                                  `initial_quantity` int,
                                  `manufacturing_date` DATETIME(6),
                                  `manufacturing_hour` int(11),
                                  `minimum_temperature` DECIMAL(19,2),
                                  `product_id` int(11),
                                  `section_id` int(11),
                                  PRIMARY KEY(`id`),
                                  UNIQUE(`batch_number`),
                                  FOREIGN KEY (`product_id`) REFERENCES `products`(`id`),  -- Corrigido para 'products'
                                  FOREIGN KEY (`section_id`) REFERENCES `sections`(`id`)  -- Corrigido para 'sections'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `product_records`(
                                  `id` int(11) NOT NULL AUTO_INCREMENT,
                                  `last_update_date` DATETIME(6),
                                  `purchase_price` DECIMAL(19,2),
                                  `sale_price` DECIMAL(19,2),
                                  `product_id` int(11),
                                  PRIMARY KEY (`id`),
                                  FOREIGN KEY (`product_id`) REFERENCES `products`(`id`)  -- Corrigido para 'products'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `inbound_orders`(
                                 `id` int(11) NOT NULL AUTO_INCREMENT,
                                 `order_date` DATETIME(6),
                                 `order_number` varchar(255),
                                 `employee_id` int(11),
                                 `product_batch_id` int(11),
                                 `warehouse_id` int(11),
                                 PRIMARY KEY(`id`),
                                 UNIQUE(`order_number`),
                                 FOREIGN KEY (`employee_id`) REFERENCES `employees`(`id`),  -- Corrigido para 'employees'
                                 FOREIGN KEY (`product_batch_id`) REFERENCES `product_batches`(`id`),  -- Corrigido para 'product_batches'
                                 FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses`(`id`)  -- Corrigido para 'warehouses'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `purchase_orders`(
                                  `id` int(11) NOT NULL AUTO_INCREMENT,
                                  `order_number` varchar(255),
                                  `order_date` DATETIME(6),
                                  `tracking_code` varchar(255),
                                  `buyer_id` int(11),
                                  `product_record_id` int(11),
                                  PRIMARY KEY(`id`),
                                  UNIQUE(`order_number`),
                                  FOREIGN KEY (`buyer_id`) REFERENCES `buyers`(`id`),  -- Corrigido para 'buyers'
                                  FOREIGN KEY (`product_record_id`) REFERENCES `product_records`(`id`)  -- Corrigido para 'product_records'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;


CREATE TABLE logs (
                      id INT AUTO_INCREMENT PRIMARY KEY,   -- ID único para cada log
                      level VARCHAR(50),                    -- Para armazenar o nível do log (INFO, ERROR, etc.)
                      message TEXT,                         -- Mensagem do log
                      time VARCHAR(250) -- Data e hora do log (usando DATETIME)
);

-- POPULATE

USE `meli_fresh`;

-- DML
INSERT INTO meli_fresh.warehouses (`warehouse_code`, `address`, `telephone`, `minimum_capacity`, `minimum_temperature`) VALUES
                                                                                                                            ('WH01', '200 Warehouse Rd', '234-567-8901', 100, 0),
                                                                                                                            ('WH02', '201 Warehouse Ln', '234-567-8902', 150, -5),
                                                                                                                            ('WH03', '202 Storage Blvd', '234-567-8903', 120, 2),
                                                                                                                            ('WH04', '203 Distribution Ave', '234-567-8904', 200, -2),
                                                                                                                            ('WH05', '204 Inventory St', '234-567-8905', 180, 0),
                                                                                                                            ('WH06', '205 Logistics Way', '234-567-8906', 160, -3),
                                                                                                                            ('WH07', '206 Depot Dr', '234-567-8907', 140, 1),
                                                                                                                            ('WH08', '207 Supply Ct', '234-567-8908', 170, -4),
                                                                                                                            ('WH09', '208 Goods Rd', '234-567-8909', 130, 3),
                                                                                                                            ('WH10', '209 Freight St', '234-567-8910', 190, -1);

INSERT INTO meli_fresh.product_type (`type_name`) VALUES
                                                      ('Dairy'),
                                                      ('Fruits'),
                                                      ('Vegetables'),
                                                      ('Meat'),
                                                      ('Frozen Foods'),
                                                      ('Beverages'),
                                                      ('Snacks'),
                                                      ('Confectionery'),
                                                      ('Grains'),
                                                      ('Spices');

INSERT INTO meli_fresh.sections (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES
                                                                                                                                                                                                    ("S01", 10, 10, 1, 1, 1, 1, 1),
                                                                                                                                                                                                    ("S02", 12, 12, 2, 2, 2, 2, 2),
                                                                                                                                                                                                    ("S03", 13, 13, 3, 3, 3, 3, 3),
                                                                                                                                                                                                    ("S04", 14, 14, 4, 4, 4, 4, 4),
                                                                                                                                                                                                    ("S05", 15, 15, 5, 5, 5, 5, 5);

INSERT INTO meli_fresh.employees (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES
                                                                                                   ('E1001', 'John', 'Doe', 1),
                                                                                                   ('E1002', 'Jane', 'Smith', 2),
                                                                                                   ('E1003', 'Michael', 'Johnson', 3),
                                                                                                   ('E1004', 'Emily', 'Davis', 4),
                                                                                                   ('E1005', 'David', 'Miller', 5),
                                                                                                   ('E1006', 'Sarah', 'Wilson', 6),
                                                                                                   ('E1007', 'Robert', 'Moore', 7),
                                                                                                   ('E1008', 'Jennifer', 'Taylor', 8),
                                                                                                   ('E1009', 'William', 'Anderson', 9),
                                                                                                   ('E1010', 'Jessica', 'Thomas', 10);

INSERT INTO meli_fresh.buyers (`card_number_id`, `first_name`, `last_name`) VALUES
                                                                                ('B1001', 'Alice', 'Brown'),
                                                                                ('B1002', 'Mark', 'Jones'),
                                                                                ('B1003', 'Linda', 'Garcia'),
                                                                                ('B1004', 'Brian', 'Williams'),
                                                                                ('B1005', 'Susan', 'Martinez'),
                                                                                ('B1006', 'Richard', 'Lee'),
                                                                                ('B1007', 'Karen', 'Harris'),
                                                                                ('B1008', 'Steven', 'Clark'),
                                                                                ('B1009', 'Betty', 'Lopez'),
                                                                                ('B1010', 'Edward', 'Gonzalez');

-- Insert records into the 'countries' table
INSERT INTO meli_fresh.countries (`country_name`) VALUES
                                                      ('Country 1'),
                                                      ('Country 2'),
                                                      ('Country 3'),
                                                      ('Country 4'),
                                                      ('Country 5');

-- Insert records into the 'provinces' table
INSERT INTO meli_fresh.provinces (`province_name`, `id_country_fk`) VALUES
                                                                        ('Province A', 1),
                                                                        ('Province B', 1),
                                                                        ('Province C', 2),
                                                                        ('Province D', 3),
                                                                        ('Province E', 4);

-- Insert records into the 'locality' table
INSERT INTO meli_fresh.locality (locality_name, province_name, country_name) VALUES
                                                                                 ('Locality X', 'Province 1', 'Country A'),
                                                                                 ('Locality Y', 'Province 2', 'Country B'),
                                                                                 ('Locality Z', 'Province 3', 'Country C'),
                                                                                 ('Locality W', 'Province 4', 'Country D'),
                                                                                 ('Locality V', 'Province 5', 'Country E');

INSERT INTO meli_fresh.sellers (cid, company_name, address, telephone, locality_id) VALUES
                                                                                        (1, 'Company A', '123 Main St', '123-456-7890', 1),
                                                                                        (2, 'Company B', '456 Elm St', '123-456-7891', 2),
                                                                                        (3, 'Company C', '789 Oak St', '123-456-7892', 3),
                                                                                        (4, 'Company D', '101 Pine St', '123-456-7893', 4),
                                                                                        (5, 'Company E', '102 Maple St', '123-456-7894', 5),
                                                                                        (6, 'Company F', '103 Cedar St', '123-456-7895', 5),
                                                                                        (7, 'Company G', '104 Birch St', '123-456-7896', 4),
                                                                                        (8, 'Company H', '105 Willow St', '123-456-7897', 3),
                                                                                        (9, 'Company I', '106 Cherry St', '123-456-7898', 3),
                                                                                        (10, 'Company J', '107 Walnut St', '123-456-7899', 1);

INSERT INTO meli_fresh.products (`product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `seller_id`, `product_type_id`) VALUES
                                                                                                                                                                                                                       ('P1001', 'Product 1', 10, 5, 8, 2, 0.1, 0.2, -5, 1, 1),
                                                                                                                                                                                                                       ('P1002', 'Product 2', 12, 6, 9, 2.5, 0.15, 0.25, -6, 2, 2),
                                                                                                                                                                                                                       ('P1003', 'Product 3', 14, 7, 10, 3, 0.2, 0.3, -7, 3, 2),
                                                                                                                                                                                                                       ('P1004', 'Product 4', 16, 8, 11, 3.5, 0.25, 0.35, -8, 4, 3),
                                                                                                                                                                                                                       ('P1005', 'Product 5', 18, 9, 12, 4, 0.3, 0.4, -9, 5, 3),
                                                                                                                                                                                                                       ('P1006', 'Product 6', 20, 10, 13, 4.5, 0.35, 0.45, -10, 6, 3),
                                                                                                                                                                                                                       ('P1007', 'Product 7', 22, 11, 14, 5, 0.4, 0.5, -11, 7, 4),
                                                                                                                                                                                                                       ('P1008', 'Product 8', 24, 12, 15, 5.5, 0.45, 0.55, -12, 8, 4),
                                                                                                                                                                                                                       ('P1009', 'Product 9', 26, 13, 16, 6, 0.5, 0.6, -13, 9, 4),
                                                                                                                                                                                                                       ('P1010', 'Product 10', 28, 14, 17, 6.5, 0.55, 0.65, -14, 10, 5);

-- Insert records into the 'carriers' table
INSERT INTO meli_fresh.carriers (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES
                                                                                                   ('C001', 'Carrier A', '500 Carrier Rd', '345-678-9011', 1),
                                                                                                   ('C002', 'Carrier B', '501 Carrier Ln', '345-678-9012', 2),
                                                                                                   ('C003', 'Carrier C', '502 Transport Blvd', '345-678-9013', 3),
                                                                                                   ('C004', 'Carrier D', '503 Logistics Ave', '345-678-9014', 4),
                                                                                                   ('C005', 'Carrier E', '504 Freight St', '345-678-9015', 5);

-- Insert records into the 'product_batches' table
INSERT INTO meli_fresh.product_batches (`batch_number`, `current_quantity`, `current_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `minimum_temperature`, `product_id`, `section_id`) VALUES
                                                                                                                                                                                                                                      ('B0001', 500, 10, '2024-12-01 00:00:00', 1000, '2023-01-01 00:00:00', 10, -5, 1, 1),
                                                                                                                                                                                                                                      ('B0002', 550, 15, '2024-11-01 00:00:00', 1100, '2023-02-01 00:00:00', 11, -6, 2, 2),
                                                                                                                                                                                                                                      ('B0003', 600, 20, '2024-10-01 00:00:00', 1200, '2023-03-01 00:00:00', 12, -7, 2, 2),
                                                                                                                                                                                                                                      ('B0004', 650, 25, '2024-09-01 00:00:00', 1300, '2023-04-01 00:00:00', 13, -8, 3, 3),
                                                                                                                                                                                                                                      ('B0005', 700, 30, '2024-08-01 00:00:00', 1400, '2023-05-01 00:00:00', 14, -9, 3, 3),
                                                                                                                                                                                                                                      ('B0006', 750, 10, '2024-12-01 00:00:00', 1000, '2023-01-01 00:00:00', 10, -5, 3, 3),
                                                                                                                                                                                                                                      ('B0007', 800, 15, '2024-11-01 00:00:00', 1100, '2023-02-01 00:00:00', 11, -6, 4, 4),
                                                                                                                                                                                                                                      ('B0008', 850, 20, '2024-10-01 00:00:00', 1200, '2023-03-01 00:00:00', 12, -7, 4, 4),
                                                                                                                                                                                                                                      ('B0009', 900, 25, '2024-09-01 00:00:00', 1300, '2023-04-01 00:00:00', 13, -8, 4, 4),
                                                                                                                                                                                                                                      ('B0010', 1000, 30, '2024-08-01 00:00:00', 1400, '2023-05-01 00:00:00', 14, -9, 5, 5);

-- Insert records into the 'product_records' table
INSERT INTO meli_fresh.product_records (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES
                                                                                                              ('2023-09-01 00:00:00', 10.50, 15.75, 1),
                                                                                                              ('2023-09-02 00:00:00', 11.00, 16.25, 2),
                                                                                                              ('2023-09-03 00:00:00', 11.50, 16.75, 3),
                                                                                                              ('2023-09-04 00:00:00', 12.00, 17.25, 4),
                                                                                                              ('2023-09-05 00:00:00', 12.50, 17.75, 5);

-- Insert records into the 'inbound_orders' table
INSERT INTO meli_fresh.inbound_orders (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`) VALUES
                                                                                                                            ('2023-07-10 00:00:00', 'IO001', 1, 1, 1),
                                                                                                                            ('2023-07-11 00:00:00', 'IO002', 2, 2, 2),
                                                                                                                            ('2023-07-12 00:00:00', 'IO003', 3, 3, 3),
                                                                                                                            ('2023-07-13 00:00:00', 'IO004', 4, 4, 4),
                                                                                                                            ('2023-07-14 00:00:00', 'IO005', 5, 5, 5);

-- Insert records into the 'purchase_orders' table
INSERT INTO meli_fresh.purchase_orders (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `product_record_id`) VALUES
                                                                                                                            ('PO001', '2023-08-10 00:00:00', 'TC001', 1, 1),
                                                                                                                            ('PO002', '2023-08-11 00:00:00', 'TC002', 2, 2),
                                                                                                                            ('PO003', '2023-08-12 00:00:00', 'TC003', 3, 3),
                                                                                                                            ('PO004', '2023-08-13 00:00:00', 'TC004', 4, 4),
                                                                                                                            ('PO005', '2023-08-14 00:00:00', 'TC005', 5, 5);    id INT AUTO_INCREMENT PRIMARY KEY,   -- ID único para cada log
                      level VARCHAR(50),                    -- Para armazenar o nível do log (INFO, ERROR, etc.)
                      message TEXT,                         -- Mensagem do log
                      time VARCHAR(250) -- Data e hora do log (usando DATETIME)
);

-- POPULATE

USE `meli_fresh`;

-- DML
INSERT INTO meli_fresh.warehouses (`warehouse_code`, `address`, `telephone`, `minimum_capacity`, `minimum_temperature`) VALUES
('WH01', '200 Warehouse Rd', '234-567-8901', 100, 0),
('WH02', '201 Warehouse Ln', '234-567-8902', 150, -5),
('WH03', '202 Storage Blvd', '234-567-8903', 120, 2),
('WH04', '203 Distribution Ave', '234-567-8904', 200, -2),
('WH05', '204 Inventory St', '234-567-8905', 180, 0),
('WH06', '205 Logistics Way', '234-567-8906', 160, -3),
('WH07', '206 Depot Dr', '234-567-8907', 140, 1),
('WH08', '207 Supply Ct', '234-567-8908', 170, -4),
('WH09', '208 Goods Rd', '234-567-8909', 130, 3),
('WH10', '209 Freight St', '234-567-8910', 190, -1);

INSERT INTO meli_fresh.product_type (`type_name`) VALUES
('Dairy'),
('Fruits'),
('Vegetables'),
('Meat'),
('Frozen Foods'),
('Beverages'),
('Snacks'),
('Confectionery'),
('Grains'),
('Spices');

INSERT INTO meli_fresh.sections (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES
("S01", 10, 10, 1, 1, 1, 1, 1),
("S02", 12, 12, 2, 2, 2, 2, 2),
("S03", 13, 13, 3, 3, 3, 3, 3),
("S04", 14, 14, 4, 4, 4, 4, 4),
("S05", 15, 15, 5, 5, 5, 5, 5);

INSERT INTO meli_fresh.employees (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES
('E1001', 'John', 'Doe', 1),
('E1002', 'Jane', 'Smith', 2),
('E1003', 'Michael', 'Johnson', 3),
('E1004', 'Emily', 'Davis', 4),
('E1005', 'David', 'Miller', 5),
('E1006', 'Sarah', 'Wilson', 6),
('E1007', 'Robert', 'Moore', 7),
('E1008', 'Jennifer', 'Taylor', 8),
('E1009', 'William', 'Anderson', 9),
('E1010', 'Jessica', 'Thomas', 10);

INSERT INTO meli_fresh.buyers (`card_number_id`, `first_name`, `last_name`) VALUES
('B1001', 'Alice', 'Brown'),
('B1002', 'Mark', 'Jones'),
('B1003', 'Linda', 'Garcia'),
('B1004', 'Brian', 'Williams'),
('B1005', 'Susan', 'Martinez'),
('B1006', 'Richard', 'Lee'),
('B1007', 'Karen', 'Harris'),
('B1008', 'Steven', 'Clark'),
('B1009', 'Betty', 'Lopez'),
('B1010', 'Edward', 'Gonzalez');

-- Insert records into the 'countries' table
INSERT INTO meli_fresh.countries (`country_name`) VALUES
('Country 1'),
('Country 2'),
('Country 3'),
('Country 4'),
('Country 5');

-- Insert records into the 'provinces' table
INSERT INTO meli_fresh.provinces (`province_name`, `id_country_fk`) VALUES
('Province A', 1),
('Province B', 1),
('Province C', 2),
('Province D', 3),
('Province E', 4);

-- Insert records into the 'locality' table
INSERT INTO meli_fresh.locality (locality_name, province_name, country_name) VALUES
('Locality X', 'Province 1', 'Country A'),
('Locality Y', 'Province 2', 'Country B'),
('Locality Z', 'Province 3', 'Country C'),
('Locality W', 'Province 4', 'Country D'),
('Locality V', 'Province 5', 'Country E');

INSERT INTO meli_fresh.sellers (cid, company_name, address, telephone, locality_id) VALUES
(1, 'Company A', '123 Main St', '123-456-7890', 1),
(2, 'Company B', '456 Elm St', '123-456-7891', 2),
(3, 'Company C', '789 Oak St', '123-456-7892', 3),
(4, 'Company D', '101 Pine St', '123-456-7893', 4),
(5, 'Company E', '102 Maple St', '123-456-7894', 5),
(6, 'Company F', '103 Cedar St', '123-456-7895', 5),
(7, 'Company G', '104 Birch St', '123-456-7896', 4),
(8, 'Company H', '105 Willow St', '123-456-7897', 3),
(9, 'Company I', '106 Cherry St', '123-456-7898', 3),
(10, 'Company J', '107 Walnut St', '123-456-7899', 1);

INSERT INTO meli_fresh.products (`product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `seller_id`, `product_type_id`) VALUES
('P1001', 'Product 1', 10, 5, 8, 2, 0.1, 0.2, -5, 1, 1),
('P1002', 'Product 2', 12, 6, 9, 2.5, 0.15, 0.25, -6, 2, 2),
('P1003', 'Product 3', 14, 7, 10, 3, 0.2, 0.3, -7, 3, 2),
('P1004', 'Product 4', 16, 8, 11, 3.5, 0.25, 0.35, -8, 4, 3),
('P1005', 'Product 5', 18, 9, 12, 4, 0.3, 0.4, -9, 5, 3),
('P1006', 'Product 6', 20, 10, 13, 4.5, 0.35, 0.45, -10, 6, 3),
('P1007', 'Product 7', 22, 11, 14, 5, 0.4, 0.5, -11, 7, 4),
('P1008', 'Product 8', 24, 12, 15, 5.5, 0.45, 0.55, -12, 8, 4),
('P1009', 'Product 9', 26, 13, 16, 6, 0.5, 0.6, -13, 9, 4),
('P1010', 'Product 10', 28, 14, 17, 6.5, 0.55, 0.65, -14, 10, 5);

-- Insert records into the 'carriers' table
INSERT INTO meli_fresh.carriers (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES
('C001', 'Carrier A', '500 Carrier Rd', '345-678-9011', 1),
('C002', 'Carrier B', '501 Carrier Ln', '345-678-9012', 2),
('C003', 'Carrier C', '502 Transport Blvd', '345-678-9013', 3),
('C004', 'Carrier D', '503 Logistics Ave', '345-678-9014', 4),
('C005', 'Carrier E', '504 Freight St', '345-678-9015', 5);

-- Insert records into the 'product_batches' table
INSERT INTO meli_fresh.product_batches (`batch_number`, `current_quantity`, `current_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `minimum_temperature`, `product_id`, `section_id`) VALUES
('B0001', 500, 10, '2024-12-01 00:00:00', 1000, '2023-01-01 00:00:00', 10, -5, 1, 1),
('B0002', 550, 15, '2024-11-01 00:00:00', 1100, '2023-02-01 00:00:00', 11, -6, 2, 2),
('B0003', 600, 20, '2024-10-01 00:00:00', 1200, '2023-03-01 00:00:00', 12, -7, 2, 2),
('B0004', 650, 25, '2024-09-01 00:00:00', 1300, '2023-04-01 00:00:00', 13, -8, 3, 3),
('B0005', 700, 30, '2024-08-01 00:00:00', 1400, '2023-05-01 00:00:00', 14, -9, 3, 3),
('B0006', 750, 10, '2024-12-01 00:00:00', 1000, '2023-01-01 00:00:00', 10, -5, 3, 3),
('B0007', 800, 15, '2024-11-01 00:00:00', 1100, '2023-02-01 00:00:00', 11, -6, 4, 4),
('B0008', 850, 20, '2024-10-01 00:00:00', 1200, '2023-03-01 00:00:00', 12, -7, 4, 4),
('B0009', 900, 25, '2024-09-01 00:00:00', 1300, '2023-04-01 00:00:00', 13, -8, 4, 4),
('B0010', 1000, 30, '2024-08-01 00:00:00', 1400, '2023-05-01 00:00:00', 14, -9, 5, 5);

-- Insert records into the 'product_records' table
INSERT INTO meli_fresh.product_records (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES
('2023-09-01 00:00:00', 10.50, 15.75, 1),
('2023-09-02 00:00:00', 11.00, 16.25, 2),
('2023-09-03 00:00:00', 11.50, 16.75, 3),
('2023-09-04 00:00:00', 12.00, 17.25, 4),
('2023-09-05 00:00:00', 12.50, 17.75, 5);

-- Insert records into the 'inbound_orders' table
INSERT INTO meli_fresh.inbound_orders (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`) VALUES
('2023-07-10 00:00:00', 'IO001', 1, 1, 1),
('2023-07-11 00:00:00', 'IO002', 2, 2, 2),
('2023-07-12 00:00:00', 'IO003', 3, 3, 3),
('2023-07-13 00:00:00', 'IO004', 4, 4, 4),
('2023-07-14 00:00:00', 'IO005', 5, 5, 5);

-- Insert records into the 'purchase_orders' table
INSERT INTO meli_fresh.purchase_orders (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `product_record_id`) VALUES
('PO001', '2023-08-10 00:00:00', 'TC001', 1, 1),
('PO002', '2023-08-11 00:00:00', 'TC002', 2, 2),
('PO003', '2023-08-12 00:00:00', 'TC003', 3, 3),
('PO004', '2023-08-13 00:00:00', 'TC004', 4, 4),
('PO005', '2023-08-14 00:00:00', 'TC005', 5, 5);