package main

import (
	"html/template"
	"log"
	"os"

	"go-server/db"
	"go-server/handlers"
	"go-server/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	loadEnvVariables()

	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Initialize the database connection
	dbConn, err := db.NewDB(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbConn.Close()

	r := gin.Default()

	// Serve static files from the "static" directory
	r.Static("/static", "./static")
	//  Define a custom template function
	r.SetFuncMap(template.FuncMap{
		"add1": func(i int) int {
			return i + 1
		},
	})
	// Load HTML templates
	r.LoadHTMLGlob("templates/*.html")

	// Unprotected routes
	r.GET("/health-check", handlers.HealthCheck)
	r.GET("/", handlers.RenderIndexPage)
	r.GET("/forget-password-page", handlers.RenderForgotPasswordPage)
	r.POST("/signup", handlers.SignUp(dbConn))
	r.POST("/login", handlers.Login(dbConn))
	r.POST("/logout", handlers.Logout())
	r.POST("/forget-password", handlers.ForgotPassword(dbConn))
	r.POST("/reset-password", handlers.ResetPassword(dbConn))

	// Protected routes
	protected := r.Group("/api/v1", middleware.AuthMiddleware("admin", "general"))

	// Homepage
	protected.GET("/homepage", handlers.RenderHomePage(dbConn))

	// User
	protected.GET("/get-current-user", handlers.GetCurrentUser(dbConn))

	// Location Details
	protected.GET("/location-details", handlers.GetLocationDetails(dbConn))
	protected.POST("/location-details", handlers.CreateNewLocationDetails(dbConn))
	protected.PUT("/location-details/:id", handlers.UpdateDeviceLocationDetail(dbConn))
	protected.DELETE("/location-details/:id", handlers.DeleteDeviceLocationDetail(dbConn))
	protected.GET("/location-details/pdf", handlers.DownloadDeviceLocationDetailPDF(dbConn))
	protected.GET("/location-details/excel", handlers.DownloadDeviceLocationDetail(dbConn))

	// Owner Details
	protected.GET("/owner-details", handlers.GetOwnerDetails(dbConn))
	protected.POST("/owner-details", handlers.CreateNewOwnerDetails(dbConn))
	protected.PUT("/owner-details/:id", handlers.UpdateDeviceAMCOwnerDetail(dbConn))
	protected.DELETE("/owner-details/:id", handlers.DeleteDeviceAMCOwnerDetail(dbConn))
	protected.GET("/owner-details/pdf", handlers.DownloadDeviceAMCOwnerDetailPDF(dbConn))
	protected.GET("/owner-details/excel", handlers.DownloadDeviceAMCOwnerDetail(dbConn))

	// Power Details
	protected.GET("/power-details", handlers.GetPowerDetails(dbConn))
	protected.POST("/power-details", handlers.CreateNewPowerDetails(dbConn))
	protected.PUT("/power-details/:id", handlers.UpdateDevicePowerDetail(dbConn))
	protected.DELETE("/power-details/:id", handlers.DeleteDevicePowerDetail(dbConn))
	protected.GET("/power-details/pdf", handlers.DownloadDevicePowerDetailPDF(dbConn))
	protected.GET("/power-details/excel", handlers.DownloadDevicePowerDetail(dbConn))

	// Fiber Details
	protected.GET("/fiber-details", handlers.GetFiberDetails(dbConn))
	protected.POST("/fiber-details", handlers.CreateNewFiberDetails(dbConn))
	protected.PUT("/fiber-details/:id", handlers.UpdateDeviceEthernetFiberDetail(dbConn))
	protected.DELETE("/fiber-details/:id", handlers.DeleteDeviceEthernetFiberDetail(dbConn))
	protected.GET("/fiber-details/pdf", handlers.DownloadDeviceEthernetFiberDetailPDF(dbConn))
	protected.GET("/fiber-details/excel", handlers.DownloadDeviceEthernetFiberDetail(dbConn))

	log.Fatal(r.Run(":" + port))

}
