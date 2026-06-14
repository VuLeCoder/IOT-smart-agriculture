CREATE TABLE users (
    id              UUID PRIMARY KEY,
    email           TEXT UNIQUE NOT NULL,
    password_hash   varchar(255) NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

