package MapRotation

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	MAP_ROTATION_URL = "https://api.mozambiquehe.re/maprotation"
)

type mapRotation struct {
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

func MapRotation(c *gin.Context) {
	// .envから環境変数を読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	API_KEY := os.Getenv("API_KEY")

	// API実行 + データフェッチ
	req, _ := http.NewRequest("GET", MAP_ROTATION_URL, nil)
	req.Header.Set("Authorization", API_KEY)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("HTTP error.")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// []byte型のJSONデータをGoの構造体に変換する
	var mapRotation mapRotation
	json.Unmarshal(body, &mapRotation)
	c.IndentedJSON(http.StatusOK, mapRotation)
}
