CREATE TABLE route_regions (
    route_uuid UUID NOT NULL,
    region_uuid UUID NOT NULL,
    PRIMARY KEY (route_uuid, region_uuid)
);
