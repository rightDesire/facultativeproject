CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE routes (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    difficulty_level_uuid UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    length NUMERIC,
    coordinates geometry(LINESTRING),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);
