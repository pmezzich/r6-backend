package r6dissect

type PlayerRoundStats struct {
    Username           string  `json:"username"`
    TeamIndex          int     `json:"teamIndex"`
    Operator           string  `json:"operator"`
    Score              int     `json:"score"`
    Kills              int     `json:"kills"`
    Deaths             int     `json:"deaths"`
    Assists            int     `json:"assists"`
    Headshots          int     `json:"headshots"`
    HeadshotPercentage float64 `json:"headshotPercentage"`
    Plants             int     `json:"plants"`
    Trades             int     `json:"trades"`
    Died               bool    `json:"died"`
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
