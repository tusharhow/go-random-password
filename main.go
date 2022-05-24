package main

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/randomPass/:count", randomPassword)
	log.Fatal(r.Run(":8080"))

}

func randomPassword(c *gin.Context) {
	count := c.Param("count")

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" + "@$%&#@%$$^&$")

	intVar, err := strconv.Atoi(count)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	length := intVar
	var p strings.Builder
	for i := 0; i < length; i++ {
		p.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := p.String()
	c.JSON(200, gin.H{
		"password": str,
	})

}