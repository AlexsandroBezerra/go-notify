CREATE TYPE delivery_status AS ENUM ('pending', 'delivered');

CREATE TABLE emails
(
    id         SERIAL PRIMARY KEY,
    recipient  VARCHAR         NOT NULL,
    subject    VARCHAR         NOT NULL,
    body       VARCHAR         NOT NULL,
    priority   SMALLINT        NOT NULL,
    status     delivery_status NOT NULL,
    created_at TIMESTAMP       NOT NULL DEFAULT now()
);