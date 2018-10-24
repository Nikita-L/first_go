ALTER SYSTEM SET max_connections = 120;
ALTER SYSTEM RESET shared_buffers;

CREATE DATABASE clients;
CREATE USER admin WITH ENCRYPTED PASSWORD 'admin#';
GRANT ALL PRIVILEGES ON DATABASE clients TO admin;

CREATE TABLE people (
    id SERIAL,
    name varchar(10),
    email varchar(40),
    mobile char(10)
);