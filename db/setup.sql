DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS articles;
DROP TABLE IF EXISTS profiles;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
  id serial PRIMARY KEY,
  uuid VARCHAR(64) NOT NULL UNIQUE,
  name VARCHAR(255),
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE profiles(
  id serial PRIMARY KEY,
  uuid VARCHAR(64) NOT NULL UNIQUE,
  bio VARCHAR(255),
  user_id INTEGER REFERENCES users(id),
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE articles(
  id serial PRIMARY KEY,
  uuid VARCHAR(64) NOT NULL UNIQUE,
  tags VARCHAR(255),
  content TEXT,
  user_id INTEGER REFERENCES users(id),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE comments(
  id serial PRIMARY KEY,
  uuid VARCHAR(64) NOT NULL UNIQUE,
  content TEXT,
  user_id INTEGER REFERENCES users(id),
  article_id INTEGER REFERENCES articles(id),
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE sessions(
  id serial PRIMARY KEY,
  uuid VARCHAR(64) NOT NULL UNIQUE,
  email VARCHAR(255),
  user_id INTEGER REFERENCES users(id),
  created_at TIMESTAMP NOT NULL
);