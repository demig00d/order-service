CREATE TABLE IF NOT EXISTS orders(
      order_uid  VARCHAR(255) NOT NULL PRIMARY KEY
     ,"order"    jsonb
);
