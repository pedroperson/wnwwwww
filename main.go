package main

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const devPort = "5000"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", renderHome)

	port := os.Getenv("PORT")
	if port == "" {
		port = devPort
	}
	router.Run(":" + port)
}

func renderHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Main website",
		"Version": rand.Intn(1000),
	})
}
