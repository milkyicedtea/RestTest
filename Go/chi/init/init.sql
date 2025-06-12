CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL
);

-- Populate 10,000 users
DO
$$
BEGIN
  FOR i IN 1..10000 LOOP
    INSERT INTO users (username, email)
    VALUES (FORMAT('user%d', i), FORMAT('user%d@example.com', i));
  END LOOP;
END
$$;