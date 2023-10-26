package handlers

import (
	"go-server/db"
	"go-server/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
)

func GetFiberDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := db.GetAllDeviceEthernetFiberDetail()
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusOK, "fiber_details.html", gin.H{"data": nil})
			return
		}
		c.HTML(http.StatusOK, "fiber_details.html", gin.H{"data": data})

	}
}

func CreateNewFiberDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.DeviceEthernetFiberDetail

		// Retrieve form data
		serialNumber := c.PostForm("serial_number")
		deviceMakeModel := c.PostForm("device_make_model")
		model := c.PostForm("model")
		deviceType := c.PostForm("device_type")
		devicePhysicalPort := c.PostForm("device_physical_port")
		devicePortType := c.PostForm("device_port_type")
		devicePortMACWWN := c.PostForm("device_port_macwwn")
		connectedDevicePort := c.PostForm("connected_device_port")

		data = models.DeviceEthernetFiberDetail{
			SerialNumber:        serialNumber,
			DeviceMakeModel:     deviceMakeModel,
			Model:               model,
			DeviceType:          deviceType,
			DevicePhysicalPort:  devicePhysicalPort,
			DevicePortType:      devicePortType,
			DevicePortMACWWN:    devicePortMACWWN,
			ConnectedDevicePort: connectedDevicePort,
		}

		if err := db.CreateDeviceEthernetFiberDetail(&data); err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "Failed to create entry"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Entry Added Successfully"})
	}
}

func UpdateDeviceEthernetFiberDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var data models.DeviceEthernetFiberDetail
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		if err := db.UpdateDeviceEthernetFiberDetail(id, &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update DeviceEthernetFiberDetail"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceEthernetFiberDetail updated successfully"})
	}
}

func DeleteDeviceEthernetFiberDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := db.DeleteDeviceEthernetFiberDetail(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete DeviceEthernetFiberDetail"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceEthernetFiberDetail deleted successfully"})
	}
}

func DownloadDeviceEthernetFiberDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM device_ethernet_fiber")
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the database"})
			return
		}
		defer rows.Close()

		file := xlsx.NewFile()
		sheet, err := file.AddSheet("DeviceEthernetFiberDetails")
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Excel sheet"})
			return
		}

		headerRow := sheet.AddRow()
		headerRow.AddCell().SetString("ID")
		headerRow.AddCell().SetString("Serial Number")
		headerRow.AddCell().SetString("Device Make Model")
		headerRow.AddCell().SetString("Model")
		headerRow.AddCell().SetString("Device Type")
		headerRow.AddCell().SetString("Device Physical Port")
		headerRow.AddCell().SetString("Device Port Type")
		headerRow.AddCell().SetString("Device Port MAC/WWN")
		headerRow.AddCell().SetString("Connected Device Port")

		for rows.Next() {
			var device models.DeviceEthernetFiberDetail
			if err := rows.Scan(&device.Id, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.DevicePhysicalPort, &device.DevicePortType, &device.DevicePortMACWWN, &device.ConnectedDevicePort); err != nil {
				log.Fatal(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan database row"})
				return
			}
			dataRow := sheet.AddRow()
			dataRow.AddCell().SetInt(device.Id)
			dataRow.AddCell().SetString(device.SerialNumber)
			dataRow.AddCell().SetString(device.DeviceMakeModel)
			dataRow.AddCell().SetString(device.Model)
			dataRow.AddCell().SetString(device.DeviceType)
			dataRow.AddCell().SetString(device.DevicePhysicalPort)
			dataRow.AddCell().SetString(device.DevicePortType)
			dataRow.AddCell().SetString(device.DevicePortMACWWN)
			dataRow.AddCell().SetString(device.ConnectedDevicePort)
		}

		c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Header("Content-Disposition", "attachment; filename=DeviceEthernetFiberDetails.xlsx")
		err = file.Write(c.Writer)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write Excel file to response"})
		}
	}
}

func DownloadDeviceEthernetFiberDetailPDF(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM device_ethernet_fiber")
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the database"})
			return
		}
		defer rows.Close()

		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()

		pdf.SetFont("Arial", "", 12)

		headers := []string{"ID", "Serial Number", "Device Make Model", "Model", "Device Type", "Device Physical Port", "Device Port Type", "Device Port MAC/WWN", "Connected Device Port"}
		colWidths := []float64{10, 30, 50, 20, 30, 30, 30, 50, 50}

		for i, header := range headers {
			pdf.CellFormat(colWidths[i], 10, header, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)

		for rows.Next() {
			var device models.DeviceEthernetFiberDetail
			if err := rows.Scan(&device.Id, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.DevicePhysicalPort, &device.DevicePortType, &device.DevicePortMACWWN, &device.ConnectedDevicePort); err != nil {
				log.Fatal(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan database row"})
				return
			}

			data := []string{
				strconv.Itoa(device.Id),
				device.SerialNumber,
				device.DeviceMakeModel,
				device.Model,
				device.DeviceType,
				device.DevicePhysicalPort,
				device.DevicePortType,
				device.DevicePortMACWWN,
				device.ConnectedDevicePort,
			}

			for i, str := range data {
				pdf.CellFormat(colWidths[i], 10, str, "1", 0, "", false, 0, "")
			}
			pdf.Ln(-1)
		}

		c.Header("Content-Type", "application/pdf")
		c.Header("Content-Disposition", "attachment; filename=DeviceEthernetFiberDetails.pdf")
		pdf.Output(c.Writer)
	}
}
