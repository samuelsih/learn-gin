package entity

type Dosen struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
}