package models

type DeviceLocationDetail struct {
	Id               int    `json:"id"`
	SerialNumber     string `json:"serial_number"`
	DeviceMakeModel  string `json:"device_make_model"`
	Model            string `json:"model"`
	DeviceType       string `json:"device_type"`
	DataCenter       string `json:"data_center"`
	Region           string `json:"region"`
	DCLocation       string `json:"dc_location"`
	DeviceLocation   string `json:"device_location"`
	DeviceRowNumber  int    `json:"device_row_number"`
	DeviceRackNumber int    `json:"device_rack_number"`
	DeviceRUNumber   string `json:"device_ru_number"`
}
