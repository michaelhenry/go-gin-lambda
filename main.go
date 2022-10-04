package main

import (
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  if isLambda() {
    gateway.ListenAndServe(":8080", r)
  } else {
    http.ListenAndServe(":8080", r)
  }
}

func isLambda() bool {
  if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
    return true
  }
  return false
}