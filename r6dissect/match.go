package r6dissect


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
