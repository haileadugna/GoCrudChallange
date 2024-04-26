package handlers

import (
	"net/http"
	"gocrudchallange/app/models"

	"github.com/gin-gonic/gin"
)

func RegisterPersonRoutes(router *gin.Engine) {
	router.GET("/person", getPersons)
	router.GET("/person/:id", getPerson)
	router.POST("/person", createPerson)
	router.PUT("/person/:id", updatePerson)
	router.DELETE("/person/:id", deletePerson)
}

func getPersons(c *gin.Context) {
	persons := models.GetAllPersons()
	c.JSON(http.StatusOK, persons)
}

func getPerson(c *gin.Context) {
	id := c.Param("id")
	person, exists := models.GetPerson(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func createPerson(c *gin.Context) {
	var person models.Person
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdPerson, _ := models.CreatePerson(person)
	c.JSON(http.StatusCreated, createdPerson)
}

func updatePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedPerson, exists := models.UpdatePerson(id, person)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, updatedPerson)
}

func deletePerson(c *gin.Context) {
	id := c.Param("id")
	deletedPerson, exists := models.DeletePerson(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, deletedPerson)
}
