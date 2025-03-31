CREATE TABLE reviews (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    user_uuid UUID NOT NULL,
    route_uuid UUID NOT NULL,
    rating INTEGER NOT NULL,
    comment TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);
