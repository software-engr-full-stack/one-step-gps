-- Switch to using the `app` database.
USE app;

-- Create a `resources` table.
CREATE TABLE IF NOT EXISTS resources (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created DATETIME NOT NULL,
  expires DATETIME NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_resources_created ON resources(created);

-- Add some dummy records.
INSERT INTO resources (title, content, created, expires) VALUES (
  'An old silent pond',
  'An old silent pond...\nA frog jumps into the pond, splash!',
  UTC_TIMESTAMP(),
  DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO resources (title, content, created, expires) VALUES (
  'Over the wintry forest',
  'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.',
  UTC_TIMESTAMP(),
  DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO resources (title, content, created, expires) VALUES (
  'First autumn morning',
  'First autumn morning\nthe mirror I stare into...',
  UTC_TIMESTAMP(),
  DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
