package main

import (
	"net/http"
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Person structure
type Person struct {
	ID      string   `json:"id"`
	Name    string   `json:"name" binding:"required"`
	Age     int      `json:"age" binding:"required"`
	Hobbies []string `json:"hobbies"`
}

// Database
var (
	persons = make(map[string]Person)
	mutex   sync.RWMutex
)

func main() {
	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// Routes
	r.GET("/person", getPersons)
	r.GET("/person/:id", getPerson)
	r.POST("/person", createPerson)
	r.PUT("/person/:id", updatePerson)
	r.DELETE("/person/:id", deletePerson)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func getPersons(c *gin.Context) {
	mutex.RLock()
	defer mutex.RUnlock()

	var result []Person
	for _, person := range persons {
		result = append(result, person)
	}
	c.JSON(http.StatusOK, result)
}

func getPerson(c *gin.Context) {
	id := c.Param("id")

	mutex.RLock()
	defer mutex.RUnlock()

	person, exists := persons[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func createPerson(c *gin.Context) {
	var person Person
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person.ID = uuid.New().String()

	mutex.Lock()
	persons[person.ID] = person
	mutex.Unlock()

	c.JSON(http.StatusCreated, person)
}

func updatePerson(c *gin.Context) {
	id := c.Param("id")

	mutex.RLock()
	_, exists := persons[id]
	mutex.RUnlock()
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	var updatedPerson Person
	if err := c.BindJSON(&updatedPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mutex.Lock()
	persons[id] = updatedPerson
	mutex.Unlock()

	c.JSON(http.StatusOK, updatedPerson)
}

func deletePerson(c *gin.Context) {
	id := c.Param("id")

	mutex.RLock()
	_, exists := persons[id]
	mutex.RUnlock()
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	mutex.Lock()
	delete(persons, id)
	mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{"message": "Person deleted successfully"})
}
