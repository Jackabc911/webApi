CREATE TABLE users (
    id bigserial not null primary key,
    login varchar not null unique,
    hashedpassword varchar not null,
    secretnumber varchar not null
);