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
		items := []MenuItem{
			// Drinks - Coffee
			{Name: "Espresso", Description: "Rich and bold single shot", Category: "Drinks", Price: 3.50},
			{Name: "Flat White", Description: "Velvety microfoam over a double ristretto", Category: "Drinks", Price: 4.50},
			{Name: "Cappuccino", Description: "Espresso with steamed milk and a crown of foam", Category: "Drinks", Price: 4.50},
			{Name: "Latte", Description: "Smooth espresso with silky steamed milk", Category: "Drinks", Price: 4.75},
			{Name: "Mocha", Description: "Espresso, steamed milk, and rich chocolate", Category: "Drinks", Price: 5.25},
			{Name: "Cold Brew", Description: "Slow-steeped for 18 hours, smooth and naturally sweet", Category: "Drinks", Price: 4.25},
			{Name: "Americano", Description: "Espresso lengthened with hot water", Category: "Drinks", Price: 3.75},
			// Drinks - Tea & Other
			{Name: "Matcha Latte", Description: "Ceremonial-grade matcha whisked with steamed oat milk", Category: "Drinks", Price: 5.50},
			{Name: "Iced Chai Latte", Description: "Spiced chai concentrate over ice with oat milk", Category: "Drinks", Price: 5.00},
			{Name: "Earl Grey Tea", Description: "Fragrant black tea with bergamot, served in a pot", Category: "Drinks", Price: 3.50},
			{Name: "Fresh Orange Juice", Description: "Cold-pressed Valencia oranges", Category: "Drinks", Price: 5.00},
			// Food - Lunch
			{Name: "Avocado Toast", Description: "Smashed avocado on sourdough with chili flakes, lime, and feta", Category: "Food", Price: 8.50},
			{Name: "Chicken Pesto Wrap", Description: "Grilled chicken, basil pesto, sun-dried tomatoes, and arugula", Category: "Food", Price: 9.50},
			{Name: "BLT Sandwich", Description: "Crispy bacon, lettuce, tomato, and garlic aioli on toasted ciabatta", Category: "Food", Price: 8.00},
			{Name: "Caprese Panini", Description: "Fresh mozzarella, tomato, and basil pressed on focaccia", Category: "Food", Price: 8.50},
			{Name: "Soup of the Day", Description: "Ask your barista what's simmering", Category: "Food", Price: 6.50},
			// Pastries & Snacks
			{Name: "Croissant", Description: "Buttery, flaky, and baked fresh every morning", Category: "Pastries", Price: 3.75},
			{Name: "Cinnamon Roll", Description: "Warm, gooey, and generously iced", Category: "Pastries", Price: 4.25},
			{Name: "Blueberry Muffin", Description: "Packed with wild blueberries and a crumble top", Category: "Pastries", Price: 3.75},
			{Name: "Banana Bread", Description: "Moist, walnut-studded, sliced thick", Category: "Pastries", Price: 4.00},
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
