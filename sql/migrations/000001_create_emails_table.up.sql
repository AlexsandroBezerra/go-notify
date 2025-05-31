CREATE TABLE emails
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    recipient  VARCHAR                        NOT NULL,
    subject    VARCHAR                        NOT NULL,
    body       VARCHAR                        NOT NULL,
    priority   SMALLINT                       NOT NULL,
    created_at TIMESTAMP        DEFAULT now() NOT NULL
);