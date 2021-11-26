package entity

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	Semester int16  `gorm:"type:int" json:"semester"`
	URL      string `gorm:"type:varchar(255)" json:"url"`
	DosenID  uint64 `gorm:"not null" json:"dosen_id"`
	Dosen     Dosen `gorm:"foreignkey:DosenID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"dosen"`
}