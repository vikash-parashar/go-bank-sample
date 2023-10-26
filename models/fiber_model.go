package models

type DeviceEthernetFiberDetail struct {
	Id                  int    `json:"id"`
	SerialNumber        string `json:"serial_number"`
	DeviceMakeModel     string `json:"device_make_model"`
	Model               string `json:"model"`
	DeviceType          string `json:"device_type"`
	DevicePhysicalPort  string `json:"device_physical_port"`
	DevicePortType      string `json:"device_port_type"`
	DevicePortMACWWN    string `json:"device_port_macwwn"`
	ConnectedDevicePort string `json:"connected_device_port"`
}
