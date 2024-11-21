package main

import (
	"fmt"

	"github.com/chajaykrishna/go-fiber-crm/database"
	"github.com/chajaykrishna/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/getLeads", lead.GetLeads)
	app.Get("/getLead/{id}", lead.GetLead)
	app.Post("/lead", lead.CreateLead)
	// app.Delete("/lead/{id} ")
}

func initDatabases() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("database migrated")
}

func main() {

	app := fiber.New()
	initDatabases()
	setupRoutes(app)
	app.Listen(":3000")
	defer database.DBConn.Close()
}
