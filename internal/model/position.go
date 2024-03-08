package model

type Position struct {
	PositionID    	int16 `json:"position_id"`
	Name string `json:"name"`
	Description string `json:"description"`
}