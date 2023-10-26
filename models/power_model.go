package models

type DevicePowerDetail struct {
	Id              int     `json:"id"`
	SerialNumber    string  `json:"serial_number"`
	DeviceMakeModel string  `json:"device_make_model"`
	Model           string  `json:"model"`
	DeviceType      string  `json:"device_type"`
	TotalPowerWatt  int     `json:"total_power_watt"`
	TotalBTU        float64 `json:"total_btu"`
	TotalPowerCable int     `json:"total_power_cable"`
	PowerSocketType string  `json:"power_socket_type"`
}
