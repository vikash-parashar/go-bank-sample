package models

import "time"

type DeviceAMCOwnerDetail struct {
	Id              int       `json:"id"`
	SerialNumber    string    `json:"serial_number"`
	DeviceMakeModel string    `json:"device_make_model"`
	Model           string    `json:"model"`
	PONumber        string    `json:"po_number"`
	POOrderDate     time.Time `json:"po_order_date"`
	EOSLDate        time.Time `json:"eosl_date"`
	AMCStartDate    time.Time `json:"amc_start_date"`
	AMCEndDate      time.Time `json:"amc_end_date"`
	DeviceOwner     string    `json:"device_owner"`
}
