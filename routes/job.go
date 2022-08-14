package routes

import (
	"jobportal/database"
	"jobportal/models"

	"github.com/gofiber/fiber/v2"
)

func AddJob(c *fiber.Ctx) error {
	job := new(models.Job)

	if err := c.BodyParser(job); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&job)

	return c.Status(200).JSON(job)
}

func ListJob(c *fiber.Ctx) error {
	jobs := []models.Job{}
	database.DB.Db.Find(&jobs)

	return c.Status(200).JSON(jobs)
}
