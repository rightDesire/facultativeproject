CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE regions (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    boundary geometry(POLYGON),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);
