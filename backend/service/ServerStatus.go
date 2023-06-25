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
	SERVER_STATUS_URL = "https://api.mozambiquehe.re/servers"
)

type serverStatus struct {
	OriginLogin struct {
		EUWest struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"EU-West"`
		EUEast struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"EU-East"`
		USWest struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"US-West"`
		USCentral struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"US-Central"`
		USEast struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"US-East"`
		SouthAmerica struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"SouthAmerica"`
		Asia struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Asia"`
	} `json:"Origin_login"`
	EANovafusion struct {
		EUWest struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"EU-West"`
		EUEast struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"EU-East"`
		USWest struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"US-West"`
		USCentral struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"US-Central"`
		USEast struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"US-East"`
		SouthAmerica struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"SouthAmerica"`
		Asia struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Asia"`
	} `json:"EA_novafusion"`
	EAAccounts struct {
		EUWest struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"EU-West"`
		EUEast struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"EU-East"`
		USWest struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"US-West"`
		USCentral struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"US-Central"`
		USEast struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"US-East"`
		SouthAmerica struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"SouthAmerica"`
		Asia struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Asia"`
	} `json:"EA_accounts"`
	ApexOauthCrossplay struct {
		EUWest struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"EU-West"`
		EUEast struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"EU-East"`
		USWest struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"US-West"`
		USCentral struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"US-Central"`
		USEast struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"US-East"`
		SouthAmerica struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"SouthAmerica"`
		Asia struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Asia"`
	} `json:"ApexOauth_Crossplay"`
	SelfCoreTest struct {
		StatusWebsite struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Status-website"`
		StatsAPI struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Stats-API"`
		Overflow1 struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Overflow-#1"`
		Overflow2 struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Overflow-#2"`
		OriginAPI struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Origin-API"`
		PlaystationAPI struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Playstation-API"`
		XboxAPI struct {
			Status         string `json:"Status"`
			HTTPCode       int    `json:"HTTPCode"`
			ResponseTime   int    `json:"ResponseTime"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Xbox-API"`
	} `json:"selfCoreTest"`
	OtherPlatforms struct {
		PlaystationNetwork struct {
			Status         string `json:"Status"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Playstation-Network"`
		XboxLive struct {
			Status         string `json:"Status"`
			QueryTimestamp int    `json:"QueryTimestamp"`
		} `json:"Xbox-Live"`
	} `json:"otherPlatforms"`
}

func ServerStatus(c *gin.Context) {
	// .envから環境変数を読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	API_KEY := os.Getenv("API_KEY")

	// API実行 + データフェッチ
	req, _ := http.NewRequest("GET", SERVER_STATUS_URL, nil)
	req.Header.Set("Authorization", API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("HTTP error.")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// []byte型のJSONデータをGoの構造体に変換する
	var ServerStatus serverStatus
	json.Unmarshal(body, &ServerStatus)
	c.IndentedJSON(http.StatusOK, ServerStatus)
}
