CREATE TABLE IF NOT EXISTS phone_numbers (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER REFERENCES users(id),
  phone VARCHAR(12) NOT NULL,
  description VARCHAR(255), 
  is_fax BOOLEAN,
  
  UNIQUE(user_id, phone)
);
