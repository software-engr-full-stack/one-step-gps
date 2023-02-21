-- Switch to using the `one_step_gps` database.
USE one_step_gps;

-- Add some dummy records.
INSERT INTO customers (name, business_category, payload, color, img_url, created, updated) VALUES (
  'UPS',
  'Package delivery',
  'Packages',
  'green',
  'https://images.pexels.com/photos/6868797/pexels-photo-6868797.jpeg',
  UTC_TIMESTAMP(),
  UTC_TIMESTAMP()
);

INSERT INTO customers (name, business_category, payload, color, img_url, created, updated) VALUES (
  'Bruce Wayne, Inc.',
  'Private investigations',
  'Batmobile',
  'blue',
  'https://images.pexels.com/photos/12966795/pexels-photo-12966795.jpeg',
  UTC_TIMESTAMP(),
  UTC_TIMESTAMP()
);

INSERT INTO customers (name, business_category, payload, color, img_url, created, updated) VALUES (
  'Fuji Heavy Industries',
  'Construction equipment rental',
  'Backhoe',
  'red',
  'https://images.pexels.com/photos/1078884/pexels-photo-1078884.jpeg',
  UTC_TIMESTAMP(),
  UTC_TIMESTAMP()
);

INSERT INTO customers (name, business_category, payload, color, img_url, created, updated) VALUES (
  'FBI',
  'Federal law enforcement',
  '*Classified*',
  'yellow',
  'https://images.pexels.com/photos/20258/pexels-photo.jpg',
  UTC_TIMESTAMP(),
  UTC_TIMESTAMP()
);
