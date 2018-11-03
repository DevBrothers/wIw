package models

type User struct {
	UserPrivate
}

type UserLogin struct {
	Username string
	Password string
}

type UserPublic struct {
	BaseModel
	Name         string  `json:""`
	PublicWishes []Item  `json:"wishes"`
	Groups       []Group `json:""`
}

type UserPrivate struct {
	UserLogin
	UserPublic
}

/*func NewUser() *User{
	var user User
	return &user
}*/
