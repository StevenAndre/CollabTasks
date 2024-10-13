package model



type TeamMember struct{
	User 		User  `json:"member"`
	TeamID     string `json:"team_id"`
	PositionID int16 `json:"position_id"`

}