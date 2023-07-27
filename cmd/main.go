package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"schools/internal/schools"
)

func main() {

	// Parse command-line flags
	var portFlag string
	flag.StringVar(&portFlag, "port", "", "Port number for the server")
	flag.Parse()

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default port.")
	}

	// Get the port from the environment variable or use the default (8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// If portFlag is provided, use it instead of the default or env value
	if portFlag != "" {
		port = portFlag
	}

	startServer(port)

}

func startServer(port string) {
	// Create a new Echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	schoolRepo := schools.NewSchoolRepository()
	schoolService := schools.NewSchoolService(schoolRepo)
	schoolHandler := schools.NewSchoolHandler(schoolService)
	schoolHandler.Register(e)

	// Start the server
	addr := ":" + port
	e.Logger.Fatal(e.Start(addr))
}
