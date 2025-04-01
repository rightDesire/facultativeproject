CREATE INDEX idx_routes_coordinates ON routes USING GIST (coordinates);
