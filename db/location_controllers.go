package db

import (
	"go-server/models"
	"log"
)

// CreateDeviceLocationDetail creates a new record in the DeviceLocationDetail table.
func (db *DB) CreateDeviceLocationDetail(data *models.DeviceLocationDetail) error {
	query := `
		INSERT INTO device_location (serial_number, device_make_model, model, device_type, data_center, region, dc_location, device_location, device_row_number, device_rack_number, device_ru_number)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.DataCenter, data.Region, data.DCLocation, data.DeviceLocation, data.DeviceRowNumber, data.DeviceRackNumber, data.DeviceRUNumber)
	if err != nil {
		log.Printf("Error creating DeviceLocationDetail: %v", err)
		return err
	}
	return nil
}

// GetAllDeviceLocationDetail retrieves all records from the DeviceLocationDetail table.
func (db *DB) GetAllDeviceLocationDetail() ([]models.DeviceLocationDetail, error) {
	query := "SELECT * FROM device_location"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying DeviceLocationDetail: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.DeviceLocationDetail
	for rows.Next() {
		var data models.DeviceLocationDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DataCenter, &data.Region, &data.DCLocation, &data.DeviceLocation, &data.DeviceRowNumber, &data.DeviceRackNumber, &data.DeviceRUNumber)
		if err != nil {
			log.Printf("Error scanning DeviceLocationDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
}

// UpdateDeviceLocationDetail updates an existing record in the device_location table based on the ID.
func (db *DB) UpdateDeviceLocationDetail(id int, data *models.DeviceLocationDetail) error {
	query := `
        UPDATE device_location
        SET serial_number = $2, device_make_model = $3, model = $4, device_type = $5, data_center = $6, region = $7, dc_location = $8, device_location = $9, device_row_number = $10, device_rack_number = $11, device_ru_number = $12
        WHERE id = $1
    `
	_, err := db.Exec(query, id, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.DataCenter, data.Region, data.DCLocation, data.DeviceLocation, data.DeviceRowNumber, data.DeviceRackNumber, data.DeviceRUNumber)
	if err != nil {
		log.Printf("Error updating DeviceLocationDetail: %v", err)
		return err
	}
	return nil
}

// DeleteDeviceLocationDetail deletes a record from the device_location table based on the ID.
func (db *DB) DeleteDeviceLocationDetail(id int) error {
	query := "DELETE FROM device_location WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting DeviceLocationDetail: %v", err)
		return err
	}
	return nil
}

// FetchDataFromTable3 retrieves data from table 3.
func (db *DB) FetchDataFromDeviceLocation() ([]*models.DeviceLocationDetail, error) {
	query := "SELECT * FROM device_location"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error fetching data from table 3: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []*models.DeviceLocationDetail
	for rows.Next() {
		var data models.DeviceLocationDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DataCenter, &data.Region, &data.DCLocation, &data.DeviceLocation, &data.DeviceRowNumber, &data.DeviceRackNumber, &data.DeviceRUNumber)
		if err != nil {
			log.Printf("Error scanning data from table 3: %v", err)
			return nil, err
		}
		results = append(results, &data)
	}

	return results, nil
}
