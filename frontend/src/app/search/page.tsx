"use client";

import { ChangeEvent, useState } from "react";

export type UserData = {
  global: {
    name: string
    uid: string
    avatar: any
    platform: string
    level: number
    toNextLevelPercent: number
    internalUpdateCount: number
    bans: {
      isActive: boolean
      remainingSeconds: number
      last_banReason: string
    }
    rank: {
      rankScore: number
      rankName: string
      rankDiv: number
      ladderPosPlatform: number
      rankImg: string
      rankedSeason: string
      ALStopPercent: number
      ALStopInt: number
      ALStopPercentGlobal: number
      ALStopIntGlobal: number
      ALSFlag: boolean
    }
    arena: {
      rankScore: number
      rankName: string
      rankDiv: number
      ladderPosPlatform: number
      rankImg: string
      rankedSeason: string
      ALStopPercent: string
      ALStopInt: string
      ALStopPercentGlobal: string
      ALStopIntGlobal: string
      ALSFlag: boolean
    }
    battlepass: {
      level: number
      history: {
        season1: number
        season2: number
        season3: number
        season4: number
        season5: number
        season6: number
        season7: number
        season8: number
        season9: number
        season10: number
        season11: number
        season12: number
        season13: number
        season14: number
        season15: number
        season16: number
        season17: number
      }
    }
    internalParsingVersion: number
    badges: Array<{
      name: string
      value: number
    }>
    levelPrestige: number
  }
  realtime: {
    lobbyState: string
    isOnline: number
    isInGame: number
    canJoin: number
    partyFull: number
    selectedLegend: string
    currentState: string
    currentStateSinceTimestamp: number
    currentStateAsText: string
  }
  legends: {
    selected: {
      LegendName: string
      data: Array<{
        name: string
        value: number
        key: string
        global: boolean
      }>
      gameInfo: {
        skin: string
        skinRarity: string
        frame: string
        frameRarity: string
        pose: string
        poseRarity: string
        intro: string
        introRarity: string
        badges: Array<{
          name: string
          value: number
          category: string
        }>
      }
      ImgAssets: {
        icon: string
        banner: string
      }
    }
    all: {
      Global: {
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Revenant: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: number
            topPercent: number
          }
        }>
        gameInfo: {
          badges: Array<{
            name: string
            value: number
          }>
        }
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Crypto: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: number
            topPercent: number
          }
        }>
        gameInfo: {
          badges: Array<{
            name: string
            value: number
          }>
        }
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Horizon: {
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Gibraltar: {
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Wattson: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: number
            topPercent: number
          }
        }>
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Fuse: {
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Bangalore: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: number
            topPercent: number
          }
        }>
        gameInfo: {
          badges: Array<{
            name: string
            value: number
          }>
        }
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Wraith: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: string
            topPercent: string
          }
        }>
        gameInfo: {
          badges: Array<{
            name: string
            value: number
          }>
        }
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Octane: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: number
            topPercent: number
          }
        }>
        gameInfo: {
          badges: Array<{
            name: string
            value: number
          }>
        }
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Bloodhound: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: number
            topPercent: number
          }
        }>
        gameInfo: {
          badges: Array<{
            name: string
            value: number
          }>
        }
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Caustic: {
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Lifeline: {
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Pathfinder: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: string
            topPercent: string
          }
        }>
        gameInfo: {
          badges: Array<{
            name: string
            value: number
          }>
        }
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Loba: {
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Mirage: {
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Rampart: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: string
            topPercent: string
          }
        }>
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Valkyrie: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: number
            topPercent: number
          }
        }>
        gameInfo: {
          badges: Array<{
            name: string
            value: number
          }>
        }
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Seer: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: number
            topPercent: number
          }
        }>
        gameInfo: {
          badges: Array<{
            name: string
            value: number
          }>
        }
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Ash: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: number
            topPercent: number
          }
        }>
        gameInfo: {
          badges: Array<{
            name: string
            value: number
          }>
        }
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      "Mad Maggie": {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: number
            topPercent: number
          }
        }>
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Newcastle: {
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Vantage: {
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Catalyst: {
        ImgAssets: {
          icon: string
          banner: string
        }
      }
      Ballistic: {
        data: Array<{
          name: string
          value: number
          key: string
          rank: {
            rankPos: number
            topPercent: number
          }
          rankPlatformSpecific: {
            rankPos: number
            topPercent: number
          }
        }>
        ImgAssets: {
          icon: string
          banner: string
        }
      }
    }
  }
  mozambiquehere_internal: {
    isNewToDB: boolean
    clusterSrv: string
  }
  total: {
    kills: {
      name: string
      value: number
    }
    headshots: {
      name: string
      value: number
    }
    silenced_targets_kills: {
      name: string
      value: number
    }
    shield_damage: {
      name: string
      value: number
    }
    specialEvent_kills: {
      name: string
      value: number
    }
    specialEvent_damage: {
      name: string
      value: number
    }
    specialEvent_wins: {
      name: string
      value: number
    }
    damage: {
      name: string
      value: number
    }
    revives: {
      name: string
      value: number
    }
    games_played: {
      name: string
      value: number
    }
    kills_as_kill_leader: {
      name: string
      value: number
    }
    enemies_scanned: {
      name: string
      value: number
    }
    bullets_amped: {
      name: string
      value: number
    }
    wins_season_12: {
      name: string
      value: number
    }
    ultimate_boosted_travel_distance: {
      name: string
      value: number
    }
    wins_season_13: {
      name: string
      value: number
    }
    tactical_enemies_overheated: {
      name: string
      value: number
    }
    passive_sling_kills: {
      name: string
      value: number
    }
    ultimate_tempest_damage: {
      name: string
      value: number
    }
    kd: {
      value: string
      name: string
    }
  }
  processingTime: number
}

const Search = ({
  children,
}: {
  children: React.ReactNode
}) => {
  const [player, setPlayer] = useState("");
  const [platform, setPlatform] = useState("");
  const [response, setResponse] = useState<UserData | null>(null);

  const handleInputChangePlayer = (event: ChangeEvent<HTMLInputElement>) => setPlayer(event.target.value)
  const handleInputChangePlatform = (event: ChangeEvent<HTMLInputElement>) => setPlatform(event.target.value)

  const getRank = async () => {
    const URL = "http://localhost:8080/search-rank?player=" + player + "&platform=" + platform;
    const res = await fetch(URL);
    setResponse(await res.json())
  }

  return (
    <section>
      <nav>
        <h1>Search Rank</h1>
        <div>
          <input type="text" placeholder="k2font" onChange={handleInputChangePlayer} />
          <input type="text" placeholder="PS4, PC, ..." onChange={handleInputChangePlatform} />
          <button onClick={getRank}>search</button>
        </div>
        <div>
          <p>Rank: {response?.global.rank.rankName}</p>
        </div>
      </nav>
      {children}
    </section>
  )
}

export default Search;