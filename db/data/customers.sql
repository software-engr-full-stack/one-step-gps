-- Switch to using the `one_step_gps` database.
USE one_step_gps;

-- Add some dummy records.
INSERT INTO customers (name, business_category, payload, img_url, created, updated) VALUES (
  'UHL',
  'Package delivery',
  'Packages',
  'https://images.pexels.com/photos/6868797/pexels-photo-6868797.jpeg',
  UTC_TIMESTAMP(),
  UTC_TIMESTAMP()
);

INSERT INTO customers (name, business_category, payload, img_url, created, updated) VALUES (
  'Bruce Wayne, Inc.',
  'Private investigations',
  'Batmobile',
  'https://images.pexels.com/photos/12966795/pexels-photo-12966795.jpeg',
  UTC_TIMESTAMP(),
  UTC_TIMESTAMP()
);

INSERT INTO customers (name, business_category, payload, img_url, created, updated) VALUES (
  'Fuji Heavy Industries',
  'Construction equipment rental',
  'Backhoe',
  'https://images.pexels.com/photos/1078884/pexels-photo-1078884.jpeg',
  UTC_TIMESTAMP(),
  UTC_TIMESTAMP()
);

INSERT INTO customers (name, business_category, payload, img_url, created, updated) VALUES (
  'FBI',
  'Federal law enforcement',
  '*Classified*',
  'https://images.pexels.com/photos/20258/pexels-photo.jpg',
  UTC_TIMESTAMP(),
  UTC_TIMESTAMP()
);
