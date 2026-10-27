package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"math/rand"
	"recommendation/data"  // replace with your actual project name
	"time"
)

func GetOrigamiOfTheDay(c *gin.Context) {
	origamis := data.GetDailyOrigami()
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	selectedOrigami := origamis[rand.Intn(len(origamis))] // get a random origami
	
	c.JSON(200, selectedOrigami)
}

func GetRecommendationStatus(c *gin.Context) {
	// Here you would typically check some aspects of your service to determine its status.
	// If everything's ok, return operational. Otherwise, return a different status.
	// This is a simple example without real checks, adjust according to your needs.

	// Example checks might include:
	// - Database connectivity
	// - External API/service availability
	// - Disk space, memory usage, etc.

	status := "operational"  // or "down", "maintenance", etc.

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}

func StartAPI() {
	r := gin.Default()
	r.GET("/origami-of-the-day", GetOrigamiOfTheDay)

	// TODO: Add more routes as needed
	
	r.Run(":8080")  // start the server on port 8080
}
