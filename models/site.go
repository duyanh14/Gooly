package models

type Site struct {
	ID          uint
	URL         string
	Title       string `sql:"type:CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Description	string `sql:"type:CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Thubnail    string
	Icon 		string
	Update 		int64 	`gorm:"autoCreateTime"`
}