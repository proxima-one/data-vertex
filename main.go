package main

import (
	"github.com/gin-gonic/gin"
	"github.com/99designs/gqlgen/handler"
	//Data vertex
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config := nil
	applicationVertex := CreateDataVertex(config)
	r := gin.Default()
	go r.POST("/query", applicationVertex.query())
	go r.GET("/", playgroundHandler())
	r.Run(":4000")
}

func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
