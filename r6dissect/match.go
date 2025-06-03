package r6dissect

type PlayerRoundStats struct {
    Name     string `json:"name"`
    Kills    int    `json:"kills"`
    Deaths   int    `json:"deaths"`
    Assists  int    `json:"assists"`
    Headshot int    `json:"headshot"`
    Plants   int    `json:"plants"`
    Trades   int    `json:"trades"`
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
