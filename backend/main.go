package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MAP_ROTATION_URL = "https://api.mozambiquehe.re/maprotation"
)

type MapRotation struct {
	Current struct {
		Start             int    `json:"start"`
		End               int    `json:"end"`
		ReadableDateStart string `json:"readableDate_start"`
		ReadableDateEnd   string `json:"readableDate_end"`
		Map               string `json:"map"`
		Code              string `json:"code"`
		DurationInSecs    int    `json:"DurationInSecs"`
		DurationInMinutes int    `json:"DurationInMinutes"`
		Asset             string `json:"asset"`
		RemainingSecs     int    `json:"remainingSecs"`
		RemainingMins     int    `json:"remainingMins"`
		RemainingTimer    string `json:"remainingTimer"`
	} `json:"current"`
	Next struct {
		Start             int    `json:"start"`
		End               int    `json:"end"`
		ReadableDateStart string `json:"readableDate_start"`
		ReadableDateEnd   string `json:"readableDate_end"`
		Map               string `json:"map"`
		Code              string `json:"code"`
		DurationInSecs    int    `json:"DurationInSecs"`
		DurationInMinutes int    `json:"DurationInMinutes"`
	} `json:"next"`
}

func mapRotation(c *gin.Context) {
	req, _ := http.NewRequest("GET", MAP_ROTATION_URL, nil)
	req.Header.Set("Authorization", "ff9fd10daf2611e4a4fec89efb26d368")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var mapData MapRotation
	json.Unmarshal(body, &mapData)
	c.IndentedJSON(http.StatusOK, mapData)
}

func main() {
	r := gin.Default()
	r.GET("/map-rotation", mapRotation)
	r.Run()
}
