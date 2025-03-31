-- Включаем расширение PostGIS
CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE locations (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    coordinates geometry(POINT, 4326),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);
