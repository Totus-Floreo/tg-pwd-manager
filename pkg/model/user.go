package model

type User struct {
	UserID  int64
	Storage map[string]*Credentials
}
