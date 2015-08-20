CREATE TABLE messages (
  id         SERIAL PRIMARY KEY,
  message    TEXT,
  user_name  TEXT,
  type       CHARACTER(256),
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE,
);
