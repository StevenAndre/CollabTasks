package dtos

type UserResisterDto struct {
	Name     string `json:"name" validate:"required"`
	Lastname string `json:"lastname" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
type UserUpdateDto struct {
	Name     string `json:"name" validate:"required"`
	Lastname string `json:"lastname" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type ProjectDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by" validate:"required"`
}

type NotificationDto struct {
	Issue   string `json:"issue" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type TaskDto struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type EvidenceDto struct {
	Comment string `json:"comment" validate:"required"`
}
type PositionDto struct{
	Name string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type TeamDto struct{
	Name string `json:"name" validate:"required"`
}

type TeamMemberDto struct {
	TeamID string `json:"team_id" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
	PositionID int16  `json:"position_id" validate:"required"`
}
