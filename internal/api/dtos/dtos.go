package dtos


type UserResisterDto struct {
	Name          		string         	 `json:"name" validate:"required"`
	Lastname      		string         	 `json:"lastname" validate:"required"`
	Username      		string         	 `json:"username" validate:"required"`
	Email         		string         	 `json:"email" validate:"required,email"`
	Password      		string         	 `json:"password" validate:"required,min=8"`
}
type UserUpdateDto struct {
	Name          		string         	 `json:"name" validate:"required"`
	Lastname      		string         	 `json:"lastname" validate:"required"`
	Username      		string         	 `json:"username" validate:"required"`
	Email         		string         	 `json:"email" validate:"required,email"`
	Password      		string         	 `json:"password" validate:"required,min=8"`
}