package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
}

func CheckAuth(id int, username string) bool {
	var auth Auth
	db.Select("id").Where(Auth{ID: id, Username: username}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
