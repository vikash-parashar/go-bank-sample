package db

import (
	"database/sql"
	"log"
)

func CreateDatabaseTables(db *sql.DB) error {

	// Create the tables
	createTables := `
		CREATE TABLE IF NOT EXISTS device_location_detail (
			serial_number TEXT PRIMARY KEY,
			device_make_model TEXT,
			model TEXT,
			device_type TEXT,
			data_center TEXT,
			region TEXT,
			dc_location TEXT,
			device_location TEXT,
			device_row_number INT,
			device_rack_number INT,
			device_ru_number TEXT
		);
		
		CREATE TABLE IF NOT EXISTS device_amc_owner_detail (
			serial_number TEXT PRIMARY KEY,
			device_make_model TEXT,
			model TEXT,
			po_number TEXT,
			po_order_date DATE,
			eosl_date DATE,
			amc_start_date DATE,
			amc_end_date DATE,
			device_owner TEXT
		);
		
		CREATE TABLE IF NOT EXISTS device_power_detail (
			serial_number TEXT PRIMARY KEY,
			device_make_model TEXT,
			model TEXT,
			device_type TEXT,
			total_power_watt INT,
			total_btu DOUBLE PRECISION,
			total_power_cable INT,
			power_socket_type TEXT
		);
		
		CREATE TABLE IF NOT EXISTS device_ethernet_fiber_detail (
			serial_number TEXT PRIMARY KEY,
			device_make_model TEXT,
			model TEXT,
			device_type TEXT,
			device_physical_port TEXT,
			device_port_type TEXT,
			device_port_macwwn TEXT,
			connected_device_port TEXT
		);
		
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(255),
			reset_token VARCHAR(255),
			reset_token_expiry TIMESTAMPTZ,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		);
		
	`
	res, err := db.Exec(createTables)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("Rows Affected : ", count)

	return nil
}
