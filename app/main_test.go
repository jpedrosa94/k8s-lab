package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// SetupRouter initializes the Gin engine and routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/animals", createAnimal)
	router.GET("/animals", getAnimals)
	router.GET("/animals/:id", getAnimalByID)
	router.PUT("/animals/:id", updateAnimal)
	router.DELETE("/animals/:id", deleteAnimal)
	return router
}

// Test creating an animal
func TestCreateAnimal(t *testing.T) {
	router := SetupRouter()

	animal := Animal{
		Name: "Giraffe",
		Type: "Mammal",
	}

	body, _ := json.Marshal(animal)
	req, _ := http.NewRequest("POST", "/animals", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Check if response contains the created animal
	var createdAnimal Animal
	err := json.Unmarshal(w.Body.Bytes(), &createdAnimal)
	assert.Nil(t, err)
	assert.Equal(t, "Giraffe", createdAnimal.Name)
	assert.Equal(t, "Mammal", createdAnimal.Type)
}

// Test getting all animals
func TestGetAnimals(t *testing.T) {
	router := SetupRouter()

	req, _ := http.NewRequest("GET", "/animals", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Check if response contains an array (even if empty)
	var animals []Animal
	err := json.Unmarshal(w.Body.Bytes(), &animals)
	assert.Nil(t, err)
}

// Test getting an animal by ID
func TestGetAnimalByID(t *testing.T) {
	router := SetupRouter()

	// Create an animal first
	animal := Animal{
		Name: "Lion",
		Type: "Mammal",
	}
	body, _ := json.Marshal(animal)
	req, _ := http.NewRequest("POST", "/animals", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Get the created animal
	createdAnimal := Animal{}
	json.Unmarshal(w.Body.Bytes(), &createdAnimal)

	// Get the animal by ID
	req, _ = http.NewRequest("GET", "/animals/"+strconv.Itoa(createdAnimal.ID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var fetchedAnimal Animal
	err := json.Unmarshal(w.Body.Bytes(), &fetchedAnimal)
	assert.Nil(t, err)
	assert.Equal(t, "Lion", fetchedAnimal.Name)
}

// Test updating an animal by ID
func TestUpdateAnimal(t *testing.T) {
	router := SetupRouter()

	// Create an animal first
	animal := Animal{
		Name: "Tiger",
		Type: "Mammal",
	}
	body, _ := json.Marshal(animal)
	req, _ := http.NewRequest("POST", "/animals", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	createdAnimal := Animal{}
	json.Unmarshal(w.Body.Bytes(), &createdAnimal)

	// Update the animal
	updatedAnimal := Animal{
		Name: "White Tiger",
		Type: "Mammal",
	}
	body, _ = json.Marshal(updatedAnimal)
	req, _ = http.NewRequest("PUT", "/animals/"+strconv.Itoa(createdAnimal.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Check if the animal was updated correctly
	var animalAfterUpdate Animal
	err := json.Unmarshal(w.Body.Bytes(), &animalAfterUpdate)
	assert.Nil(t, err)
	assert.Equal(t, "White Tiger", animalAfterUpdate.Name)
}

// Test deleting an animal by ID
func TestDeleteAnimal(t *testing.T) {
	router := SetupRouter()

	// Create an animal first
	animal := Animal{
		Name: "Elephant",
		Type: "Mammal",
	}
	body, _ := json.Marshal(animal)
	req, _ := http.NewRequest("POST", "/animals", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	createdAnimal := Animal{}
	json.Unmarshal(w.Body.Bytes(), &createdAnimal)

	// Delete the animal
	req, _ = http.NewRequest("DELETE", "/animals/"+strconv.Itoa(createdAnimal.ID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Check if animal is deleted
	req, _ = http.NewRequest("GET", "/animals/"+strconv.Itoa(createdAnimal.ID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
