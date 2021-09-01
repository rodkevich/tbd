CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS tickets
(
    id              uuid                     DEFAULT uuid_generate_v4(),
    order_number    NUMERIC(2)      NOT NULL,
    ticket_name     VARCHAR(200)    NOT NULL,
    photo_main_link VARCHAR(400),
    currency        NUMERIC(5),
    current_price   NUMERIC(100, 2) NOT NULL DEFAULT 0.00,
    discount        NUMERIC(3),
    min_price       NUMERIC(100, 2),
    max_price       NUMERIC(100, 2),
    description     VARCHAR(1000),
    phone_number    VARCHAR(50),
    is_active       BOOLEAN         NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS photo_links
(
    link_id      BIGSERIAL    NOT NULL,
    ticket_id    uuid,
    link_address VARCHAR(255) NOT NULL,
--     CONSTRAINT fk_ticket
    FOREIGN KEY (ticket_id)
        REFERENCES tickets (id)
        ON DELETE CASCADE
);



