package models

/*
Contains the models for the User Table
*/
type User struct {
	UserId     string `json:"user_id"`
	UserName   string `json:"user_name"`
	EmailId    string `json:"email_id"`
	ProfilePic string `json:"profile_pic"`
	Role       string `json:"role"`
}

func (user *User) TableName() string {
	return "user"
}
