package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	router.GET("/os", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, runtime.GOOS)
	})

	router.GET("/products", endPointHandler)
	router.GET("/products/:productId", endPointHandler)
	router.POST("/products", endPointHandler)
	router.PUT("/products/:productId", endPointHandler)
	router.DELETE("/products/:productId", endPointHandler)

	router.GET("/addStr/:x/:y", addStrHandler)
	router.GET("/addJson", addJsonHandler)

	router.GET("/productJSON", productJSONHandler)
	router.GET("/productXML", productXMLHandler)
	router.GET("/productYAML", productYAMLHandler)

	router.GET("/print", printHandler)

	v1 := router.Group("/v1")
	v1.GET("/products", v1EndPointHandler)
	v1.GET("/products/:productId", v1EndPointHandler)
	v1.POST("/products", v1EndPointHandler)
	v1.PUT("/products/:productId", v1EndPointHandler)
	v1.DELETE("/products/:productId", v1EndPointHandler)

	v2 := router.Group("/v2")
	v2.GET("/products", v2EndPointHandler)
	v2.GET("/products/:productId", v2EndPointHandler)
	v2.POST("/products", v2EndPointHandler)
	v2.PUT("/products/:productId", v2EndPointHandler)
	v2.DELETE("/products/:productId", v2EndPointHandler)

	router.Run(":8080")
}

func endPointHandler(c *gin.Context) {
	c.String(http.StatusOK, "%s %s", c.Request.Method, c.Request.URL.Path)
}

func v1EndPointHandler(c *gin.Context) {
	c.String(http.StatusOK, "%s %s", c.Request.Method, c.Request.URL.Path)
}

func v2EndPointHandler(c *gin.Context) {
	c.String(http.StatusOK, "%s %s", c.Request.Method, c.Request.URL.Path)
}

func addStrHandler(c *gin.Context) {
	x, _ := strconv.ParseFloat(c.Param("x"), 64)
	y, _ := strconv.ParseFloat(c.Param("y"), 64)
	c.String(http.StatusOK, fmt.Sprintf("%f", x+y))
}

type AddBodyJson struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func addJsonHandler(c *gin.Context) {
	var aBJ AddBodyJson
	if err := c.ShouldBindJSON(&aBJ); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Calculation error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"answer": aBJ.X + aBJ.Y,
	})
}

type Product struct {
	Id   int    `json:"id" xml:"Id" yaml:"id"`
	Name string `json:"name" xml:"Name" yaml:"name"`
}

func productJSONHandler(c *gin.Context) {
	product := Product{1, "Apple"}
	c.JSON(http.StatusOK, product)
}

func productXMLHandler(c *gin.Context) {
	product := Product{2, "Banana"}
	c.XML(http.StatusOK, product)
}

func productYAMLHandler(c *gin.Context) {
	product := Product{3, "Mango"}
	c.YAML(http.StatusOK, product)
}

type PrintJob struct {
	JobId int `json:"jobId" binding:"required,gte=10000"`
	Pages int `json:"pages" binding:"required,gte=1,lte=100"`
}

func printHandler(c *gin.Context) {
	var p PrintJob
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("PrintJob #%v started!", p.JobId),
	})
}
