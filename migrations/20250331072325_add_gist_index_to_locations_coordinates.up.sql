CREATE INDEX idx_locations_coordinates ON locations USING GIST (coordinates);
