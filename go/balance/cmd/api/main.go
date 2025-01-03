package main

import (
	// "net/http"

	"github.com/belone0/stack-trial/go/balance/internal/controllers"
	"github.com/belone0/stack-trial/go/balance/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    db.ConnectToDatabase()

    r.GET("/books", controllers.FindBooks) 

    r.Run()
}
