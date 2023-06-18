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
	PREDATOR_URL = "https://api.mozambiquehe.re/predator"
)

type predator struct {
	Rp struct {
		Pc struct {
			FoundRank            int    `json:"foundRank"`
			Val                  int    `json:"val"`
			UID                  string `json:"uid"`
			UpdateTimestamp      int    `json:"updateTimestamp"`
			TotalMastersAndPreds int    `json:"totalMastersAndPreds"`
		} `json:"PC"`
		Ps4 struct {
			FoundRank            int    `json:"foundRank"`
			Val                  int    `json:"val"`
			UID                  string `json:"uid"`
			UpdateTimestamp      int    `json:"updateTimestamp"`
			TotalMastersAndPreds int    `json:"totalMastersAndPreds"`
		} `json:"PS4"`
		X1 struct {
			FoundRank            int    `json:"foundRank"`
			Val                  int    `json:"val"`
			UID                  string `json:"uid"`
			UpdateTimestamp      int    `json:"updateTimestamp"`
			TotalMastersAndPreds int    `json:"totalMastersAndPreds"`
		} `json:"X1"`
		Switch struct {
			FoundRank            int    `json:"foundRank"`
			Val                  int    `json:"val"`
			UID                  string `json:"uid"`
			UpdateTimestamp      int    `json:"updateTimestamp"`
			TotalMastersAndPreds int    `json:"totalMastersAndPreds"`
		} `json:"SWITCH"`
	} `json:"RP"`
	Ap struct {
		Pc struct {
			FoundRank            int    `json:"foundRank"`
			Val                  int    `json:"val"`
			UID                  string `json:"uid"`
			UpdateTimestamp      int    `json:"updateTimestamp"`
			TotalMastersAndPreds int    `json:"totalMastersAndPreds"`
		} `json:"PC"`
		Ps4 struct {
			FoundRank            int    `json:"foundRank"`
			Val                  int    `json:"val"`
			UID                  string `json:"uid"`
			UpdateTimestamp      int    `json:"updateTimestamp"`
			TotalMastersAndPreds int    `json:"totalMastersAndPreds"`
		} `json:"PS4"`
		X1 struct {
			FoundRank            int    `json:"foundRank"`
			Val                  int    `json:"val"`
			UID                  string `json:"uid"`
			UpdateTimestamp      int    `json:"updateTimestamp"`
			TotalMastersAndPreds int    `json:"totalMastersAndPreds"`
		} `json:"X1"`
		Switch struct {
			FoundRank            int    `json:"foundRank"`
			Val                  int    `json:"val"`
			UID                  string `json:"uid"`
			UpdateTimestamp      int    `json:"updateTimestamp"`
			TotalMastersAndPreds int    `json:"totalMastersAndPreds"`
		} `json:"SWITCH"`
	} `json:"AP"`
}

func Predator(c *gin.Context) {
	// .envから環境変数を読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	API_KEY := os.Getenv("API_KEY")

	// API実行 + データフェッチ
	req, _ := http.NewRequest("GET", PREDATOR_URL, nil)
	req.Header.Set("Authorization", API_KEY)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("HTTP error.")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// []byte型のJSONデータをGoの構造体に変換する
	var predator predator
	json.Unmarshal(body, &predator)
	c.IndentedJSON(http.StatusOK, predator)
}
