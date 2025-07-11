-- +goose Up 

CREATE TABLE feed_follows (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  -- The user_id column is a foreign key that references the users table.
  -- It is used to associate a feed follow with a specific user.
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  -- The feed_id column is a foreign key that references the feeds table.
  feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;