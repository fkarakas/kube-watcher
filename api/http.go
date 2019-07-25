package api

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Repository interface {
	GetEndpoint(key string) (interface{}, bool)
}

func Start(repo Repository) {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"response": "ok",
		})
	})
	router.GET("/endpoint/:key", func(c *gin.Context) {
		key := c.Param("key")

		value, ok := repo.GetEndpoint(key)
		if !ok {
			c.AbortWithStatus(http.StatusNotFound)
		}
		c.JSON(http.StatusOK, value)
	})
	//go r.Run() // listen and serve on 0.0.0.0:8080
	go endless.ListenAndServe(":8080", router)
}
