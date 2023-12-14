package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todo represents the structure of the data received from the API
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func ConsumeApi(c *gin.Context) {
	// Replace the URL with the API endpoint you want to consume
	apiURL := "https://jsonplaceholder.typicode.com/todos/1"

	// Make HTTP GET request
	response, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer response.Body.Close()

	// Check the status code
	if response.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Unmarshal JSON response into the Todo struct
	var todo Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	// You can process the data or send it directly to the client
	c.JSON(http.StatusOK, gin.H{"data": todo})
}
