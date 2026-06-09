CREATE TABLE sensor_data (
    id          BIGSERIAL PRIMARY KEY,
    device_id   UUID REFERENCES devices(id) ON DELETE CASCADE,
    
    rain_level      REAL,
    light           REAL,
    soil_moisture   REAL, 
    ph              REAL,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP 
);

CREATE INDEX idx_sensor_data_device_time ON sensor_data (device_id, created_at DESC);
