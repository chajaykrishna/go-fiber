package lead

import (
	"github.com/chajaykrishna/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)

}
func GetLead(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}
func CreateLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}
func DeleteLead() {}
