CREATE TABLE developer_profiles
(
    id                      UUID      PRIMARY KEY DEFAULT gen_random_uuid(),
    name                    VARCHAR   NOT NULL,
    email                   VARCHAR   NOT NULL UNIQUE,
    phone                   VARCHAR,
    address                 TEXT,
    accept_terms_of_service BOOLEAN   NOT NULL DEFAULT FALSE,
    created_at              TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMP NOT NULL DEFAULT NOW()
);
