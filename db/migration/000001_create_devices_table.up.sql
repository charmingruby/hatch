CREATE TABLE IF NOT EXISTS devices(
    id varchar PRIMARY KEY NOT NULL,
    hardware_id varchar,
    hardware_type varchar,
    created_at timestamp DEFAULT now() NOT NULL
);

CREATE INDEX idx_devices_hardware_id ON devices (hardware_id);
