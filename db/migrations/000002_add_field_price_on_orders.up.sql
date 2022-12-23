ALTER TABLE orders
ADD COLUMN total_price int(11) NOT NULL AFTER account_zone,
ADD COLUMN price int(11) NOT NULL AFTER total_price,
ADD COLUMN fee int(11) NOT NULL AFTER price; 