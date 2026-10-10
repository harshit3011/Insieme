package models

type User struct {
	ID           string
	Name         string
	Username     string
	Email        string
	Password     string
	ProfilePic   string
	Bio          string
	AccessToken  string
	RefreshToken string
}

type RegisterDetails struct {
	ID         string
	Name       string
	Username   string
	Email      string
	Password   string
	ProfilePic string
	Bio        string
}

type Login struct {
	Email    string
	Password string
}
