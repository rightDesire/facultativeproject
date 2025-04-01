CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE route_points (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    route_uuid UUID NOT NULL,
    name VARCHAR(255),
    description TEXT,
    sequence INTEGER NOT NULL,
    coordinates geometry(POINT, 4326)
);