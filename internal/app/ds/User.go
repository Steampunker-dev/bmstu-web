package ds

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Login    string `json:"login" gorm:"type:varchar(255)"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
}
