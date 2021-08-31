CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS ticket
(
    id           uuid             DEFAULT uuid_generate_v4(),
    order_number VARCHAR NOT NULL,
    ticket_name  VARCHAR NOT NULL,
    description  VARCHAR,
    is_deleted   BOOLEAN NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS photo_links
(
    id          SERIAL  NOT NULL,
    hello_world VARCHAR,
    is_main     BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted  BOOLEAN NOT NULL DEFAULT FALSE

);

