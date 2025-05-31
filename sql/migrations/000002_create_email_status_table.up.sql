CREATE TYPE delivery_status AS ENUM ('pending', 'failed', 'delivered');

CREATE TABLE email_status
(
    id         SERIAL PRIMARY KEY,
    email_id   UUID REFERENCES emails (id) NOT NULL,
    status     delivery_status             NOT NULL,
    created_at TIMESTAMP DEFAULT now()     NOT NULL
);

CREATE INDEX email_status_email_id ON email_status (email_id);