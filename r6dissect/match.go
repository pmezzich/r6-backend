package r6dissect

type PlayerRoundStats struct {
	Name     string `json:"name"`      // Player's name
	Kills    int    `json:"kills"`     // Number of kills
	Deaths   int    `json:"deaths"`    // Number of deaths
	Assists  int    `json:"assists"`   // Number of assists
	Headshot int    `json:"headshot"`  // Number of headshots
	Plants   int    `json:"plants"`    // Number of plants
	Trades   int    `json:"trades"`    // Number of trades
	Score    int    `json:"score"`     // Optional: Total round score
	Team     int    `json:"team"`      // Optional: Team index
	Operator string `json:"operator"`  // Optional: Operator used
}


type Round struct {
    Map         string             `json:"map"`
    Operator    string             `json:"operator"`
    Result      string             `json:"result"`
    PlayerStats []PlayerRoundStats `json:"playerStats"` // Added this field
}

type Match struct {
    ID     string  `json:"id"`
    Rounds []Round `json:"rounds"`
}
