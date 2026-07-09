package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// MenuItem represents a menu item in the café menu
type MenuItem struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Category    string  `json:"category" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Available   *bool   `json:"available" gorm:"default:false"`
}

// db is the global database connection used by all handlers
var db *gorm.DB

// initDB opens the SQLite database abd creates the menu_items table
func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("menu.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate creates the table if it doesn't exist
	db.AutoMigrate(&MenuItem{})
}

// seedDB populates the database with the sample cafe items
func seedDB() {
	var count int64
	db.Model(&MenuItem{}).Count(&count)

	// Only seed if the table is empty
	if count == 0 {
		available := true
		items := []MenuItem{
			{Name: "Espresso", Description: "Rich and bold single shot", Category: "Drinks", Price: 3.50, Available: &available},
			{Name: "Cappuccino", Description: "Espresso with steamed milk foam", Category: "Drinks", Price: 4.50, Available: &available},
			{Name: "Croissant", Description: "Buttery, flaky pastry", Category: "Pastries", Price: 3.00, Available: &available},
			{Name: "Blueberry Muffin", Description: "Freshly baked with real blueberries", Category: "Pastries", Price: 3.75, Available: &available},
		}
		db.Create(&items)
		log.Println("Database seeded with sample menu items")
	}
}

// getMenuItems returns all menu items as JSON
func getMenuItems(c *gin.Context) {
	var items []MenuItem
	db.Find(&items)
	c.JSON(200, items)
}

// getMenuItem returns a single menu item by ID
func getMenuItem(c *gin.Context) {
	var item MenuItem
	if err := db.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Menu item not found"})
		return
	}
	c.JSON(200, item)
}

// createMenuItem adds a new menu item to the database
func createMenuItem(c *gin.Context) {
	var item MenuItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&item)
	c.JSON(201, item)
}

// updateMenuItem modifies an existing menu item
func updateMenuItem(c *gin.Context) {
	var item MenuItem
	if err := db.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Menu item not found"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Save(&item)
	c.JSON(200, item)
}

// deleteMenuItem removes a menu item from the database
func deleteMenuItem(c *gin.Context) {
	if err := db.Delete(&MenuItem{}, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Menu item not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Menu item deleted"})
}

func main() {
	// Initialize database and seed sample data
	initDB()
	seedDB()

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Enable CORS for all origins so the frontend and AI service can connect
	r.Use(cors.Default())

	// Group all menu routes under /api/menu
	api := r.Group("/api/menu")
	{
		api.GET("", getMenuItems)          // GET /api/menu
		api.GET("/:id", getMenuItem)       // GET /api/menu/:id
		api.POST("", createMenuItem)       // POST /api/menu
		api.PUT("/:id", updateMenuItem)    // PUT /api/menu/:id
		api.DELETE("/:id", deleteMenuItem) // DELETE /api/menu/:id
	}

	// Start the server on port 8080
	log.Println("Menu Service running on :8080")
	r.Run(":8080")
}
