CREATE TABLE comments
(
  id SERIAL NOT NULL PRIMARY KEY,
  comment VARCHAR(1000) NOT NULL,
  tweet_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT NULL,
  deleted_at TIMESTAMP DEFAULT NULL
);
ALTER TABLE comments
ADD FOREIGN KEY (tweet_id) REFERENCES tweets (id);
ALTER TABLE comments
ADD FOREIGN KEY (user_id) REFERENCES users (id);