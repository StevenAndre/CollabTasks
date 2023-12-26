package model

type User struct {
    
	UserID        		string           `json:"id_user"`
	Name          		string         	 `json:"name"`
	Lastname      		string         	 `json:"lastname"`
	Username      		string         	 `json:"username"`
	Email         		string         	 `json:"email"`
	ProfilePhoto  		string         	 `json:"profile_photo"`
}