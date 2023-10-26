package db

import (
	"go-server/models"
	"log"
)

// CreateDevicePowerDetail creates a new record in the DevicePowerDetail table.
func (db *DB) CreateDevicePowerDetail(data *models.DevicePowerDetail) error {
	query := `
		INSERT INTO device_power (serial_number, device_make_model, model, device_type, total_power_watt, total_btu, total_power_cable, power_socket_type)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.TotalPowerWatt, data.TotalBTU, data.TotalPowerCable, data.PowerSocketType)
	if err != nil {
		log.Printf("Error creating DevicePowerDetail: %v", err)
		return err
	}
	return nil
}

// GetAllDevicePowerDetail retrieves all records from the DevicePowerDetail table.
func (db *DB) GetAllDevicePowerDetail() ([]models.DevicePowerDetail, error) {
	query := "SELECT * FROM device_power"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying DevicePowerDetail: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.DevicePowerDetail
	for rows.Next() {
		var data models.DevicePowerDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.TotalPowerWatt, &data.TotalBTU, &data.TotalPowerCable, &data.PowerSocketType)
		if err != nil {
			log.Printf("Error scanning DevicePowerDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
}

// UpdateDevicePowerDetail updates an existing record in the device_power table based on the ID.
func (db *DB) UpdateDevicePowerDetail(id int, data *models.DevicePowerDetail) error {
	query := `
        UPDATE device_power
        SET serial_number = $2, device_make_model = $3, model = $4, device_type = $5, total_power_watt = $6, total_btu = $7, total_power_cable = $8, power_socket_type = $9
        WHERE id = $1
    `
	_, err := db.Exec(query, id, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.TotalPowerWatt, data.TotalBTU, data.TotalPowerCable, data.PowerSocketType)
	if err != nil {
		log.Printf("Error updating DevicePowerDetail: %v", err)
		return err
	}
	return nil
}

// DeleteDevicePowerDetail deletes a record from the device_power table based on the ID.
func (db *DB) DeleteDevicePowerDetail(id int) error {
	query := "DELETE FROM device_power WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting DevicePowerDetail: %v", err)
		return err
	}
	return nil
}

// FetchDataFromDevicePower retrieves data from table 4.
func (db *DB) FetchDataFromDevicePower() ([]*models.DevicePowerDetail, error) {
	query := "SELECT * FROM device_power"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error fetching data from table 4: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []*models.DevicePowerDetail
	for rows.Next() {
		var data models.DevicePowerDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.TotalPowerWatt, &data.TotalBTU, &data.TotalPowerCable, &data.PowerSocketType)
		if err != nil {
			log.Printf("Error scanning data from table 4: %v", err)
			return nil, err
		}
		results = append(results, &data)
	}

	return results, nil
}
