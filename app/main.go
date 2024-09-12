package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Animal struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var animals = []Animal{}
var nextID = 1

// Create a new animal
func createAnimal(c *gin.Context) {
	var newAnimal Animal

	if err := c.ShouldBindJSON(&newAnimal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAnimal.ID = nextID
	nextID++

	animals = append(animals, newAnimal)
	c.JSON(http.StatusCreated, newAnimal)
}

// Read all animals
func getAnimals(c *gin.Context) {
	c.JSON(http.StatusOK, animals)
}

// Read a specific animal by ID
func getAnimalByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid animal ID"})
		return
	}

	for _, animal := range animals {
		if animal.ID == id {
			c.JSON(http.StatusOK, animal)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Animal not found"})
}

// Update an existing animal by ID
func updateAnimal(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid animal ID"})
		return
	}

	var updatedAnimal Animal
	if err := c.ShouldBindJSON(&updatedAnimal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, animal := range animals {
		if animal.ID == id {
			animals[i].Name = updatedAnimal.Name
			animals[i].Type = updatedAnimal.Type
			c.JSON(http.StatusOK, animals[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Animal not found"})
}

// Delete an animal by ID
func deleteAnimal(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid animal ID"})
		return
	}

	for i, animal := range animals {
		if animal.ID == id {
			animals = append(animals[:i], animals[i+1:]...)
			message := "Animal " + animal.Name + " deleted"
			c.JSON(http.StatusOK, gin.H{"message": message})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Animal not found"})
}

func main() {
	router := gin.Default()

	// Routes
	router.POST("/animals", createAnimal)       // Create
	router.GET("/animals", getAnimals)          // Read all
	router.GET("/animals/:id", getAnimalByID)   // Read by ID
	router.PUT("/animals/:id", updateAnimal)    // Update
	router.DELETE("/animals/:id", deleteAnimal) // Delete

	// Start the server on port 8080
	router.Run(":8080")
}
