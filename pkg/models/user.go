package models

type User struct {
	Username string `json:"username" bson:"_id"`
	Password string `json:"password" bson:"password"`
}
