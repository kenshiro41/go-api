CREATE TABLE favorites
(
  id SERIAL NOT NULL PRIMARY KEY,
  tweet_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL
);
ALTER TABLE favorites
ADD FOREIGN KEY (tweet_id) REFERENCES tweets (id);
ALTER TABLE favorites
ADD FOREIGN KEY (user_id) REFERENCES users (id);