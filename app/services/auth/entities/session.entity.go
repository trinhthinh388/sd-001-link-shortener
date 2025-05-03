package entities

type Session struct {
	Id     string `gorm:"primary_key"`
	UserId string	`gorm:"type:varchar(22);not null;unique"`
	User 	 User   `gorm:"foreignKey:UserId"` 
}