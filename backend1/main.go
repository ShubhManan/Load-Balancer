package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := getServerPort()
	router := gin.Default()
	router.GET("/users", getUsers)

	log.Print("Starting the server on: 8081")
	err := router.Run(port)
	if err != nil {
		log.Fatalf("Server failed to start, %v\n", err)
	}
}

func getUsers(c *gin.Context) {
	log.Printf("The request is coming from %s\n", c.Request.Header.Get("x-forwarded-for"))
	c.JSON(200, gin.H{
		"users":         users,
		"numberOfUsers": len(users),
	})
}

func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8081"
	}
	return port
}

var users = []string{"Test1", "Test2", "Test3"}
