package handlers

import (
	"fmt"
	"github.com/3xlerman/web-app/app/database"
	"github.com/3xlerman/web-app/app/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Title": "Homepage",
	})
}

func Person(c *gin.Context) {

	var status int
	var message string

	jsonInput := struct {
		Firstname string `json:"firstname" binding:"required"`
		Lastname  string `json:"lastname" binging:"required"`
		Sex       string `json:"sex" binding:"required"`
		Age       string `json:"age" binding:"required"`
		Language  string `json:"language" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified: " + err.Error(),
		})
		return
	}

	user := model.User{
		ID:        primitive.NewObjectID(),
		Firstname: jsonInput.Firstname,
		Lastname:  jsonInput.Lastname,
		Sex:       jsonInput.Sex,
		Age:       jsonInput.Age,
		Language:  jsonInput.Language,
	}

	err := database.InsertUser(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}

	if status == 0 {
		message = "ok"
		status = http.StatusOK
	}

	c.JSON(status, gin.H{
		"message": message,
	})
}
