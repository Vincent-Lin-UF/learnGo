package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var articles = map[int]Article{}
var nextID = 1

func main() {
	router := gin.Default()

	// Create
	router.POST("/articles", func(c *gin.Context) {
		var newArticle Article
		if err := c.ShouldBindJSON(&newArticle); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newArticle.ID = nextID
		nextID++
		articles[newArticle.ID] = newArticle
		c.JSON(http.StatusCreated, newArticle)
	})

	// Listing all the articles
	router.GET("articles", func(c *gin.Context) {
		list := make([]Article, 0, len(articles))
		for _, a := range articles {
			list = append(list, a)
		}
		c.JSON(http.StatusOK, list)
	})

	// Fetching a single article by ID
	router.GET("/articles/:id", func(c *gin.Context) {
		// 1. Parse ID from URL
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
			return
		}

		// 2. Look up in the in-memory map
		article, exists := articles[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
			return
		}

		c.JSON(http.StatusOK, article)
	})

	router.Run(":8080")
}
