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
	SEARCH_RANK_URL = "https://api.mozambiquehe.re/bridge"
)

type searchRank struct {
	Global struct {
		Name                string `json:"name"`
		UID                 string `json:"uid"`
		Avatar              any    `json:"avatar"`
		Platform            string `json:"platform"`
		Level               int    `json:"level"`
		ToNextLevelPercent  int    `json:"toNextLevelPercent"`
		InternalUpdateCount int    `json:"internalUpdateCount"`
		Bans                struct {
			IsActive         bool   `json:"isActive"`
			RemainingSeconds int    `json:"remainingSeconds"`
			LastBanReason    string `json:"last_banReason"`
		} `json:"bans"`
		Rank struct {
			RankScore           int     `json:"rankScore"`
			RankName            string  `json:"rankName"`
			RankDiv             int     `json:"rankDiv"`
			LadderPosPlatform   int     `json:"ladderPosPlatform"`
			RankImg             string  `json:"rankImg"`
			RankedSeason        string  `json:"rankedSeason"`
			ALStopPercent       float64 `json:"ALStopPercent"`
			ALStopInt           int     `json:"ALStopInt"`
			ALStopPercentGlobal float64 `json:"ALStopPercentGlobal"`
			ALStopIntGlobal     int     `json:"ALStopIntGlobal"`
			ALSFlag             bool    `json:"ALSFlag"`
		} `json:"rank"`
		Arena struct {
			RankScore           int    `json:"rankScore"`
			RankName            string `json:"rankName"`
			RankDiv             int    `json:"rankDiv"`
			LadderPosPlatform   int    `json:"ladderPosPlatform"`
			RankImg             string `json:"rankImg"`
			RankedSeason        string `json:"rankedSeason"`
			ALStopPercent       string `json:"ALStopPercent"`
			ALStopInt           string `json:"ALStopInt"`
			ALStopPercentGlobal string `json:"ALStopPercentGlobal"`
			ALStopIntGlobal     string `json:"ALStopIntGlobal"`
			ALSFlag             bool   `json:"ALSFlag"`
		} `json:"arena"`
		Battlepass struct {
			Level   int `json:"level"`
			History struct {
				Season1  int `json:"season1"`
				Season2  int `json:"season2"`
				Season3  int `json:"season3"`
				Season4  int `json:"season4"`
				Season5  int `json:"season5"`
				Season6  int `json:"season6"`
				Season7  int `json:"season7"`
				Season8  int `json:"season8"`
				Season9  int `json:"season9"`
				Season10 int `json:"season10"`
				Season11 int `json:"season11"`
				Season12 int `json:"season12"`
				Season13 int `json:"season13"`
				Season14 int `json:"season14"`
				Season15 int `json:"season15"`
				Season16 int `json:"season16"`
				Season17 int `json:"season17"`
			} `json:"history"`
		} `json:"battlepass"`
		InternalParsingVersion int `json:"internalParsingVersion"`
		Badges                 []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"badges"`
		LevelPrestige int `json:"levelPrestige"`
	} `json:"global"`
	Realtime struct {
		LobbyState                 string `json:"lobbyState"`
		IsOnline                   int    `json:"isOnline"`
		IsInGame                   int    `json:"isInGame"`
		CanJoin                    int    `json:"canJoin"`
		PartyFull                  int    `json:"partyFull"`
		SelectedLegend             string `json:"selectedLegend"`
		CurrentState               string `json:"currentState"`
		CurrentStateSinceTimestamp int    `json:"currentStateSinceTimestamp"`
		CurrentStateAsText         string `json:"currentStateAsText"`
	} `json:"realtime"`
	Legends struct {
		Selected struct {
			LegendName string `json:"LegendName"`
			Data       []struct {
				Name   string `json:"name"`
				Value  int    `json:"value"`
				Key    string `json:"key"`
				Global bool   `json:"global"`
			} `json:"data"`
			GameInfo struct {
				Skin        string `json:"skin"`
				SkinRarity  string `json:"skinRarity"`
				Frame       string `json:"frame"`
				FrameRarity string `json:"frameRarity"`
				Pose        string `json:"pose"`
				PoseRarity  string `json:"poseRarity"`
				Intro       string `json:"intro"`
				IntroRarity string `json:"introRarity"`
				Badges      []struct {
					Name     string `json:"name"`
					Value    int    `json:"value"`
					Category string `json:"category"`
				} `json:"badges"`
			} `json:"gameInfo"`
			ImgAssets struct {
				Icon   string `json:"icon"`
				Banner string `json:"banner"`
			} `json:"ImgAssets"`
		} `json:"selected"`
		All struct {
			Global struct {
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Global"`
			Revenant struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				GameInfo struct {
					Badges []struct {
						Name  string `json:"name"`
						Value int    `json:"value"`
					} `json:"badges"`
				} `json:"gameInfo"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Revenant"`
			Crypto struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				GameInfo struct {
					Badges []struct {
						Name  string `json:"name"`
						Value int    `json:"value"`
					} `json:"badges"`
				} `json:"gameInfo"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Crypto"`
			Horizon struct {
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Horizon"`
			Gibraltar struct {
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Gibraltar"`
			Wattson struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Wattson"`
			Fuse struct {
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Fuse"`
			Bangalore struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				GameInfo struct {
					Badges []struct {
						Name  string `json:"name"`
						Value int    `json:"value"`
					} `json:"badges"`
				} `json:"gameInfo"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Bangalore"`
			Wraith struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    string `json:"rankPos"`
						TopPercent string `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				GameInfo struct {
					Badges []struct {
						Name  string `json:"name"`
						Value int    `json:"value"`
					} `json:"badges"`
				} `json:"gameInfo"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Wraith"`
			Octane struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				GameInfo struct {
					Badges []struct {
						Name  string `json:"name"`
						Value int    `json:"value"`
					} `json:"badges"`
				} `json:"gameInfo"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Octane"`
			Bloodhound struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				GameInfo struct {
					Badges []struct {
						Name  string `json:"name"`
						Value int    `json:"value"`
					} `json:"badges"`
				} `json:"gameInfo"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Bloodhound"`
			Caustic struct {
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Caustic"`
			Lifeline struct {
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Lifeline"`
			Pathfinder struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int `json:"rankPos"`
						TopPercent int `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    string `json:"rankPos"`
						TopPercent string `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				GameInfo struct {
					Badges []struct {
						Name  string `json:"name"`
						Value int    `json:"value"`
					} `json:"badges"`
				} `json:"gameInfo"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Pathfinder"`
			Loba struct {
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Loba"`
			Mirage struct {
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Mirage"`
			Rampart struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    string `json:"rankPos"`
						TopPercent string `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Rampart"`
			Valkyrie struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				GameInfo struct {
					Badges []struct {
						Name  string `json:"name"`
						Value int    `json:"value"`
					} `json:"badges"`
				} `json:"gameInfo"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Valkyrie"`
			Seer struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				GameInfo struct {
					Badges []struct {
						Name  string `json:"name"`
						Value int    `json:"value"`
					} `json:"badges"`
				} `json:"gameInfo"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Seer"`
			Ash struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				GameInfo struct {
					Badges []struct {
						Name  string `json:"name"`
						Value int    `json:"value"`
					} `json:"badges"`
				} `json:"gameInfo"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Ash"`
			MadMaggie struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Mad Maggie"`
			Newcastle struct {
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Newcastle"`
			Vantage struct {
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Vantage"`
			Catalyst struct {
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Catalyst"`
			Ballistic struct {
				Data []struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
					Key   string `json:"key"`
					Rank  struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rank"`
					RankPlatformSpecific struct {
						RankPos    int     `json:"rankPos"`
						TopPercent float64 `json:"topPercent"`
					} `json:"rankPlatformSpecific"`
				} `json:"data"`
				ImgAssets struct {
					Icon   string `json:"icon"`
					Banner string `json:"banner"`
				} `json:"ImgAssets"`
			} `json:"Ballistic"`
		} `json:"all"`
	} `json:"legends"`
	MozambiquehereInternal struct {
		IsNewToDB  bool   `json:"isNewToDB"`
		ClusterSrv string `json:"clusterSrv"`
	} `json:"mozambiquehere_internal"`
	Total struct {
		Kills struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"kills"`
		Headshots struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"headshots"`
		SilencedTargetsKills struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"silenced_targets_kills"`
		ShieldDamage struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"shield_damage"`
		SpecialEventKills struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"specialEvent_kills"`
		SpecialEventDamage struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"specialEvent_damage"`
		SpecialEventWins struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"specialEvent_wins"`
		Damage struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"damage"`
		Revives struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"revives"`
		GamesPlayed struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"games_played"`
		KillsAsKillLeader struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"kills_as_kill_leader"`
		EnemiesScanned struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"enemies_scanned"`
		BulletsAmped struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"bullets_amped"`
		WinsSeason12 struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"wins_season_12"`
		UltimateBoostedTravelDistance struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"ultimate_boosted_travel_distance"`
		WinsSeason13 struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"wins_season_13"`
		TacticalEnemiesOverheated struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"tactical_enemies_overheated"`
		PassiveSlingKills struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"passive_sling_kills"`
		UltimateTempestDamage struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"ultimate_tempest_damage"`
		Kd struct {
			Value string `json:"value"`
			Name  string `json:"name"`
		} `json:"kd"`
	} `json:"total"`
	ProcessingTime float64 `json:"processingTime"`
}

func SearchRank(c *gin.Context) {
	// .envから環境変数を読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	API_KEY := os.Getenv("API_KEY")

	// クエリパラメータの取得
	player := c.Query("player")
	platform := c.Query("platform")

	// API実行 + データフェッチ
	req, _ := http.NewRequest("GET", SEARCH_RANK_URL+"?player="+player+"&platform="+platform, nil)
	req.Header.Set("Authorization", API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("HTTP error.")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// []byte型のJSONデータをGoの構造体に変換する
	var searchRank searchRank
	json.Unmarshal(body, &searchRank)
	c.IndentedJSON(http.StatusOK, searchRank)
}
