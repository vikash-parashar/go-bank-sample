package db

import (
	"go-server/models"
	"log"
)

// CreateDeviceAMCOwnerDetail creates a new record in the device_amc_owner table.
func (db *DB) CreateDeviceAMCOwnerDetail(data *models.DeviceAMCOwnerDetail) error {
	query := `
        INSERT INTO device_amc_owner (serial_number, device_make_model, model, po_number, po_order_date, eosl_date, amc_start_date, amc_end_date, device_owner)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    `
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.PONumber, data.POOrderDate, data.EOSLDate, data.AMCStartDate, data.AMCEndDate, data.DeviceOwner)
	if err != nil {
		log.Printf("Error creating DeviceAMCOwnerDetail: %v", err)
		return err
	}
	return nil
}

// GetAllDeviceAMCOwnerDetail retrieves all records from the device_amc_owner table.
func (db *DB) GetAllDeviceAMCOwnerDetail() ([]models.DeviceAMCOwnerDetail, error) {
	query := "SELECT * FROM device_amc_owner"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying DeviceAMCOwnerDetail: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.DeviceAMCOwnerDetail
	for rows.Next() {
		var data models.DeviceAMCOwnerDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.PONumber, &data.POOrderDate, &data.EOSLDate, &data.AMCStartDate, &data.AMCEndDate, &data.DeviceOwner)
		if err != nil {
			log.Printf("Error scanning DeviceAMCOwnerDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
}

// UpdateDeviceAMCOwnerDetail updates an existing record in the device_amc_owner table based on the ID.
func (db *DB) UpdateDeviceAMCOwnerDetail(id int, data *models.DeviceAMCOwnerDetail) error {
	query := `
        UPDATE device_amc_owner
        SET serial_number = $2, device_make_model = $3, model = $4, po_number = $5, po_order_date = $6, eosl_date = $7, amc_start_date = $8, amc_end_date = $9, device_owner = $10
        WHERE id = $1
    `
	_, err := db.Exec(query, id, data.SerialNumber, data.DeviceMakeModel, data.Model, data.PONumber, data.POOrderDate, data.EOSLDate, data.AMCStartDate, data.AMCEndDate, data.DeviceOwner)
	if err != nil {
		log.Printf("Error updating DeviceAMCOwnerDetail: %v", err)
		return err
	}
	return nil
}

// DeleteDeviceAMCOwnerDetail deletes a record from the device_amc_owner table based on the ID.
func (db *DB) DeleteDeviceAMCOwnerDetail(id int) error {
	query := "DELETE FROM device_amc_owner WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting DeviceAMCOwnerDetail: %v", err)
		return err
	}
	return nil
}

// FetchDataFromTable1 retrieves data from table 1.
func (db *DB) FetchDataFromDeviceOwner() ([]*models.DeviceAMCOwnerDetail, error) {
	query := "SELECT * FROM device_amc_owner"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error fetching data from table 1: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []*models.DeviceAMCOwnerDetail
	for rows.Next() {
		var data models.DeviceAMCOwnerDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.PONumber, &data.POOrderDate, &data.EOSLDate, &data.AMCStartDate, &data.AMCEndDate, &data.DeviceOwner)
		if err != nil {
			log.Printf("Error scanning data from table 1: %v", err)
			return nil, err
		}
		results = append(results, &data)
	}

	return results, nil
}
