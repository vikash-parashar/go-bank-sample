
CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        first_name VARCHAR(255) NOT NULL,
        last_name VARCHAR(255) NOT NULL,
        phone VARCHAR(255),
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        role VARCHAR(255),
        reset_token VARCHAR(255),
        reset_token_expiry TIMESTAMPTZ,
        created_at TIMESTAMPTZ DEFAULT NOW(),
        updated_at TIMESTAMPTZ DEFAULT NOW()
    );

-- Create the device_power table

CREATE TABLE
    IF NOT EXISTS device_power (
        id SERIAL PRIMARY KEY,
        serial_number VARCHAR(255) NOT NULL,
        device_make_model VARCHAR(255),
        model VARCHAR(255),
        device_type VARCHAR(255),
        total_power_watt INT,
        total_btu DECIMAL(10, 2),
        total_power_cable INT,
        power_socket_type VARCHAR(255)
    );

-- Create the device_ethernet_fiber table

CREATE TABLE
    IF NOT EXISTS device_ethernet_fiber (
        id SERIAL PRIMARY KEY,
        serial_number VARCHAR(255) NOT NULL,
        device_make_model VARCHAR(255),
        model VARCHAR(255),
        device_type VARCHAR(255),
        device_physical_port VARCHAR(255),
        device_port_type VARCHAR(255),
        device_port_mac_address_wwn VARCHAR(255),
        connected_device_port VARCHAR(255)
    );

-- Create the device_amc_owner table

CREATE TABLE
    IF NOT EXISTS device_amc_owner (
        id SERIAL PRIMARY KEY,
        serial_number VARCHAR(255) NOT NULL,
        device_make_model VARCHAR(255),
        model VARCHAR(255),
        po_number VARCHAR(255),
        po_order_date DATE,
        eosl_date DATE,
        amc_start_date DATE,
        amc_end_date DATE,
        device_owner VARCHAR(255)
    );

-- Create the device_location table

CREATE TABLE
    IF NOT EXISTS device_location (
        id SERIAL PRIMARY KEY,
        serial_number VARCHAR(255) NOT NULL,
        device_make_model VARCHAR(255),
        model VARCHAR(255),
        device_type VARCHAR(255),
        data_center VARCHAR(255),
        region VARCHAR(255),
        dc_location VARCHAR(255),
        device_location VARCHAR(255),
        device_row_number INT,
        device_rack_number INT,
        device_ru_number VARCHAR(255)
    );