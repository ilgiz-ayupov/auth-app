CREATE TABLE phone_numbers (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER REFERENCES users(id),
  phone VARCHAR(12) NOT NULL UNIQUE,
  description VARCHAR(255), 
  is_fax BOOLEAN
);
