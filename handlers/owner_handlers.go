package handlers

import (
	"fmt"
	"go-server/db"
	"go-server/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
)

// GetOwnerDetails handles the GET request to retrieve owner details.
func GetOwnerDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := db.GetAllDeviceAMCOwnerDetail()
		if err != nil {
			log.Println(err)
			return
		}
		c.HTML(http.StatusOK, "owner_details.html", data)
	}
}

func CreateNewOwnerDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.DeviceAMCOwnerDetail

		// Retrieve form data
		idStr := c.PostForm("id")
		serialNumber := c.PostForm("serial_number")
		deviceMakeModel := c.PostForm("device_make_model")
		model := c.PostForm("model")
		poNumber := c.PostForm("po_number")
		poOrderDateStr := c.PostForm("po_order_date")
		eoslDateStr := c.PostForm("eosl_date")
		amcStartDateStr := c.PostForm("amc_start_date")
		amcEndDateStr := c.PostForm("amc_end_date")
		deviceOwner := c.PostForm("device_owner")

		// Parse the date strings into time.Time values
		poOrderDate, _ := time.Parse("2006-01-02", poOrderDateStr)
		eoslDate, _ := time.Parse("2006-01-02", eoslDateStr)
		amcStartDate, _ := time.Parse("2006-01-02", amcStartDateStr)
		amcEndDate, _ := time.Parse("2006-01-02", amcEndDateStr)

		// Convert ID to int
		id, _ := strconv.Atoi(idStr)

		// Assign the values to the DeviceAMCOwnerDetail struct
		data = models.DeviceAMCOwnerDetail{
			Id:              id,
			SerialNumber:    serialNumber,
			DeviceMakeModel: deviceMakeModel,
			Model:           model,
			PONumber:        poNumber,
			POOrderDate:     poOrderDate,
			EOSLDate:        eoslDate,
			AMCStartDate:    amcStartDate,
			AMCEndDate:      amcEndDate,
			DeviceOwner:     deviceOwner,
		}

		if err := db.CreateDeviceAMCOwnerDetail(&data); err != nil {
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Entry Added Successfully"},
		)
	}
}

// UpdateDeviceAMCOwnerDetailHandler updates a DeviceAMCOwnerDetail record based on its ID.
func UpdateDeviceAMCOwnerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var data models.DeviceAMCOwnerDetail
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		if err := db.UpdateDeviceAMCOwnerDetail(id, &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update DeviceAMCOwnerDetail"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceAMCOwnerDetail updated successfully"})
	}
}

// DeleteDeviceAMCOwnerDetailHandler deletes a DeviceAMCOwnerDetail record based on its ID.
func DeleteDeviceAMCOwnerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := db.DeleteDeviceAMCOwnerDetail(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete DeviceAMCOwnerDetail"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceAMCOwnerDetail deleted successfully"})
	}
}

func DownloadDeviceAMCOwnerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Query the database for DeviceAMCOwnerDetail data
		rows, err := db.Query("SELECT * FROM device_amc_owner")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to query the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a new Excel file
		file := xlsx.NewFile()
		sheet, err := file.AddSheet("DeviceAMCOwnerDetails")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to create Excel sheet", http.StatusInternalServerError)
			return
		}

		// Add header row
		headerRow := sheet.AddRow()
		headerRow.AddCell().SetString("ID")
		headerRow.AddCell().SetString("Serial Number")
		headerRow.AddCell().SetString("Device Make Model")
		headerRow.AddCell().SetString("Model")
		headerRow.AddCell().SetString("PO Number")
		headerRow.AddCell().SetString("PO Order Date")
		headerRow.AddCell().SetString("EOSL Date")
		headerRow.AddCell().SetString("AMC Start Date")
		headerRow.AddCell().SetString("AMC End Date")
		headerRow.AddCell().SetString("Device Owner")

		// Add data rows from the database
		for rows.Next() {
			var device models.DeviceAMCOwnerDetail
			if err := rows.Scan(&device.Id, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.PONumber, &device.POOrderDate, &device.EOSLDate, &device.AMCStartDate, &device.AMCEndDate, &device.DeviceOwner); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}
			dataRow := sheet.AddRow()
			dataRow.AddCell().SetInt(device.Id)
			dataRow.AddCell().SetString(device.SerialNumber)
			dataRow.AddCell().SetString(device.DeviceMakeModel)
			dataRow.AddCell().SetString(device.Model)
			dataRow.AddCell().SetString(device.PONumber)
			dataRow.AddCell().SetDate(device.POOrderDate)
			dataRow.AddCell().SetDate(device.EOSLDate)
			dataRow.AddCell().SetDate(device.AMCStartDate)
			dataRow.AddCell().SetDate(device.AMCEndDate)
			dataRow.AddCell().SetString(device.DeviceOwner)
		}

		// Save the Excel file to the response
		c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=DeviceAMCOwnerDetails.xlsx")
		err = file.Write(c.Writer)
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to write Excel file to response", http.StatusInternalServerError)
		}
	}
}

func DownloadDeviceAMCOwnerDetailPDF(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Query the database for DeviceAMCOwnerDetail data
		rows, err := db.Query("SELECT * FROM device_amc_owner")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to query the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a new PDF document
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()

		// Set font and text size
		pdf.SetFont("Arial", "", 12)

		// Add table headers
		headers := []string{"ID", "Serial Number", "Device Make Model", "Model", "PO Number", "PO Order Date", "EOSL Date", "AMC Start Date", "AMC End Date", "Device Owner"}
		for _, header := range headers {
			pdf.CellFormat(40, 10, header, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)

		// Add data rows from the database
		for rows.Next() {
			var device models.DeviceAMCOwnerDetail
			if err := rows.Scan(&device.Id, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.PONumber, &device.POOrderDate, &device.EOSLDate, &device.AMCStartDate, &device.AMCEndDate, &device.DeviceOwner); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}

			data := []string{
				fmt.Sprint(device.Id),
				device.SerialNumber,
				device.DeviceMakeModel,
				device.Model,
				device.PONumber,
				device.POOrderDate.Format("2006-01-02"),
				device.EOSLDate.Format("2006-01-02"),
				device.AMCStartDate.Format("2006-01-02"),
				device.AMCEndDate.Format("2006-01-02"),
				device.DeviceOwner,
			}

			for _, str := range data {
				pdf.CellFormat(40, 10, str, "1", 0, "C", false, 0, "")
			}
			pdf.Ln(-1)
		}

		// Create the PDF file
		pdf.Output(c.Writer)

		// Set response headers
		c.Header("Content-Type", "application/pdf")
		c.Header("Content-Disposition", "attachment; filename=DeviceAMCOwnerDetails.pdf")
	}
}
