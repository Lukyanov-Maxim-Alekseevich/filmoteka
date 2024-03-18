-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
   id          SERIAL PRIMARY KEY,
   email VARCHAR(255) NOT NULL,
   password_hash VARCHAR(255) NOT NULL,
   role_name VARCHAR(10) NOT NULL,
   role_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS roles
(
    id SERIAL PRIMARY KEY ,
    name VARCHAR(10) NOT NULL, 
    description VARCHAR(100) 
);

CREATE TABLE IF NOT EXISTS movies
(
    id          SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    description TEXT NOT NULL,
    release_date VARCHAR(20) NOT NULL,
    rating DECIMAL(3,1) CHECK (rating >= 0 AND rating <= 10) 
);

CREATE TABLE IF NOT EXISTS actors
(
    id          SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    gender VARCHAR(6) NOT NULL,
    date_birth VARCHAR(20) NOT NULL,
    movie_id INT NOT NULL,
    FOREIGN KEY (movie_id) REFERENCES movies(id)
);  


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE movies_actors;
DROP TABLE movies;
DROP TABLE users;
DROP TABLE roles;
-- +goose StatementEnd
