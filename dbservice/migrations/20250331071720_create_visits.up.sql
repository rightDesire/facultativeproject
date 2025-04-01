CREATE TABLE visits (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    user_uuid UUID NOT NULL,
    route_uuid UUID NOT NULL,
    location_uuid UUID NOT NULL,
    visit_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);
