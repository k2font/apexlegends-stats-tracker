package main

import (
	m "apexlegends-stats-tracker/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/map-rotation", m.MapRotation)
	r.GET("/predator", m.Predator)
	r.Run()
}
