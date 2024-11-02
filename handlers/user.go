package handlers

import (
	"golang-fiber-cicd/database"
	"golang-fiber-cicd/models"

	"github.com/gofiber/fiber/v2"
)

// Get all users
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

// Get single user by ID
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	return c.JSON(user)
}

// Create a new user
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Create(&user)
	return c.JSON(user)
}

// Update a user
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Save(&user)
	return c.JSON(user)
}

// Delete a user
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	database.DB.Delete(&user)
	return c.JSON(fiber.Map{"message": "User deleted"})
}
