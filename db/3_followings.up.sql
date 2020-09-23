CREATE TABLE followings
(
  id SERIAL NOT NULL PRIMARY KEY,
  following_id INTEGER NOT NULL,
  followed_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP DEFAULT NULL
);
ALTER TABLE followings
ADD FOREIGN KEY (following_id) REFERENCES users (id);
ALTER TABLE followings
ADD FOREIGN KEY (followed_id) REFERENCES users (id);