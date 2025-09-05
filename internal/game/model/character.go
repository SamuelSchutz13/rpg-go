package model

type Levels int

const (
	Level_1 Levels = 0
	Level_2 Levels = 100
)

type Character struct {
	Name  string `json:"name"`
	Level Levels `json:"level"`
	Hp    int    `json:"hp"`
	Xp    int    `json:"xp"`
	Class Class  `json:"class"`
}
