-- Switch to using the `one_step_gps` database.
USE one_step_gps;

-- Create a `customers` table.
CREATE TABLE IF NOT EXISTS customers (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  business_category TEXT NOT NULL,
  payload TEXT NOT NULL,
  img_url VARCHAR(100) NOT NULL,
  created DATETIME NOT NULL,
  updated DATETIME NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_customers_created ON customers(created);
