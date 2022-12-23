CREATE TABLE orders (
    id int NOT NULL AUTO_INCREMENT,
    service_code varchar(50) NOT NULL,
    account_id varchar(100) NOT NULL,
    account_zone varchar(50) NULL,
    status ENUM("success", "pending","failed"),
    created_at datetime DEFAULT NOW(),
    updated_at datetime ON UPDATE NOW(),
    PRIMARY KEY(id)
);