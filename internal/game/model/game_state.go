package model

type GameState struct {
	Character Character `json:"character"`
	PosX      int       `json:"pos_x"`
	PosY      int       `json:"pos_y"`
	Filename  string    `json:"filename"`
}
