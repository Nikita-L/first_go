#!/bin/bash

set -e
psql -v ON_ERROR_STOP=1 --username "postgres" <<-EOSQL
CREATE DATABASE clients;
CREATE USER admin WITH ENCRYPTED PASSWORD 'admin#';
GRANT ALL PRIVILEGES ON DATABASE clients TO admin;
EOSQL

psql -d clients -v ON_ERROR_STOP=1 --username "admin" <<-EOSQL
CREATE TABLE people (
    id SERIAL,
    name text,
    email text,
    mobile char(11)
);
EOSQL