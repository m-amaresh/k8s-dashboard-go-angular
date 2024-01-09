package main

import (
    "k8s-dashboard-go-angular/backend/k8s"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    router := gin.Default()

    // Configure CORS middleware options as needed
    config := cors.DefaultConfig()
    config.AllowAllOrigins = true  // In production, set to your frontend's URL instead of true
    config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}
    config.AllowHeaders = []string{"Origin", "Content-Type"}

    // Apply CORS middleware to the router
    router.Use(cors.New(config))

    // Set up routes
    router.GET("/namespaces", getNamespaces)
    // Add more routes here

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

func getNamespaces(c *gin.Context) {
    namespaces, err := k8s.ListNamespaces()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, namespaces)
}
