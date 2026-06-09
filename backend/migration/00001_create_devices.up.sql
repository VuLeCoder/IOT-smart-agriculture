CREATE TABLE devices (
    id          UUID PRIMARY KEY,
    device_name VARCHAR(100) NOT NULL,
    api_key     VARCHAR(255) UNIQUE NOT NULL,
    location    VARCHAR(255),
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
