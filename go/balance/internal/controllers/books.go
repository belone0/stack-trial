package controllers

import (
	"net/http"

	"github.com/belone0/stack-trial/go/balance/internal/db"
	"github.com/belone0/stack-trial/go/balance/internal/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	db.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}