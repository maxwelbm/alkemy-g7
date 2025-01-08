CREATE DATABASE meli_fresh;

USE meli_fresh;


CREATE TABLE seller(

	id int NOT NULL AUTO_INCREMENT,
	cid int,
	company_name VARCHAR(150),
	address VARCHAR(150),
	telephone VARCHAR(20),
	PRIMARY KEY(id),
	UNIQUE(cid)

);

CREATE TABLE warehouse(

	id int not null AUTO_INCREMENT,
	address VARCHAR(150),
	telephone VARCHAR(20),
	minimun_capacity int,
	minimun_temperature int,
	PRIMARY KEY(id)

);

CREATE TABLE product_type(

id int NOT NULL AUTO_INCREMENT,
type_name varchar(100),
PRIMARY KEY(id)

);


CREATE TABLE product(

	id int NOT NULL AUTO_INCREMENT,
	product_code VARCHAR(100),
	description varchar(255),
	width DECIMAL(19,2),
	height DECIMAL(19,2),
	length decimal (3,2),
	net_weight decimal (3,2),
	expiration_rate decimal (3,2),
	recommended_freezing_temperature decimal (3,2),
	freezing_rate DECIMAL(19,2),
	product_type_id int,
	seller_id int,
	PRIMARY KEY(id),
	UNIQUE(product_code),
	FOREIGN KEY (product_type_id) REFERENCES product_type(id),
	FOREIGN KEY (seller_id) REFERENCES seller(id)

);


CREATE TABLE employee(

id int NOT NULL AUTO_INCREMENT,
card_number_id varchar(100) NOT NULL,
first_name varchar(100),
last_name varchar(100),
warehouse_id int,
PRIMARY KEY(id),
UNIQUE(card_number_id),
FOREIGN KEY (warehouse_id) REFERENCES warehouse(id)
);


CREATE TABLE section(

	id int NOT NULL AUTO_INCREMENT,
	section_number VARCHAR(255),
	current_temperature DECIMAL(19,2),
	minimum_temperature DECIMAL(19,2),
	current_capacity int,
	minimun_capacity int,
	maximum_capacity int,
	warehouse_id int,
	product_type_id int,
	PRIMARY KEY(id),
	UNIQUE(section_number),
	FOREIGN KEY (warehouse_id) REFERENCES warehouse(id),
	FOREIGN KEY (product_type_id) REFERENCES product_type(id)
	
);


CREATE TABLE buyer(
	
	id int NOT NULL AUTO_INCREMENT,
	card_number_id varchar(100),
	first_name varchar(100),
	last_name varchar(100),
	PRIMARY KEY(id)

);

CREATE TABLE countries(
	
id int NOT NULL AUTO_INCREMENT,
country_name varchar(255),
PRIMARY KEY(id)

);

CREATE TABLE provinces(
	
id int NOT NULL AUTO_INCREMENT,
province_name varchar(255),
id_country_fk int,
PRIMARY KEY(id),
FOREIGN KEY (id_country_fk) REFERENCES countries(id)

);

CREATE TABLE locality(

id int NOT NULL AUTO_INCREMENT,
locality_name varchar(255),
province_id int,
PRIMARY KEY (id),
FOREIGN KEY (province_id) REFERENCES provinces(id)

);

CREATE TABLE carriers(
	
	id int NOT NULL AUTO_INCREMENT,
	cid VARCHAR(100),
	company_name VARCHAR(100),
	address varchar(100),
	telephone varchar(20),
	locality_id int,
	PRIMARY KEY(id),
	UNIQUE(cid),
	FOREIGN KEY (locality_id) REFERENCES locality(id)
);

CREATE TABLE product_batches(

	id int NOT NULL AUTO_INCREMENT,
	batch_number varchar(100),
	current_quantity int,
	current_temperature int,
	due_date DATETIME(6),
	initial_quantity int,
	manufacturing_date DATETIME(6),
	manufacturing_hour DATETIME(6),
	minumum_temperature DECIMAL(19,2),
	product_id int,
	section_id int,
	PRIMARY KEY(id),
	FOREIGN KEY(product_id) REFERENCES product(id),
	FOREIGN KEY (section_id) REFERENCES section(id)
);

create table product_records(

	id int NOT NULL AUTO_INCREMENT,
	last_update_date DATETIME(6),
	purchase_price DECIMAL(19,2),
	sale_price DECIMAL(19,2),
	product_id int,
	PRIMARY KEY (id),
	FOREIGN KEY(product_id) REFERENCES product(id)
);


create table inbound_orders(
id int NOT NULL AUTO_INCREMENT,
order_date DATETIME(6),
order_number varchar(255),
employee_id int,
product_batch_id int,
warehouse_id int,
PRIMARY KEY(id),
UNIQUE(order_number),
FOREIGN KEY (employee_id) REFERENCES employee(id),
FOREIGN KEY (product_batch_id) REFERENCES product_batches(id),
FOREIGN KEY (warehouse_id) REFERENCES warehouse(id)

);

create table purchase_orders(

id int NOT NULL AUTO_INCREMENT,
order_number varchar(255),
order_date DATETIME(6),
tracking_code varchar(255),
buyer_id int,
product_record_id int,
PRIMARY KEY(id),
UNIQUE(order_number),
FOREIGN KEY (buyer_id) REFERENCES buyer(id),
FOREIGN KEY (product_record_id) REFERENCES product_records(id)

);

















