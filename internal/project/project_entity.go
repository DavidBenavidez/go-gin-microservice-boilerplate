package project

import "time"

// Project table
type Project struct {
	ID              uint       `gorm:"primary_key;column:ID"`
	UUID            string     `gorm:"not null;column:PROJECT_ID"`
	Name            string     `gorm:"not null;column:PROJECT_NAME"`
	Description     string     `gorm:"not null;column:PROJECT_DESC"`
	Creator         string     `gorm:"not null;column:CREATOR"`
	StartDate       *time.Time `gorm:"not null;column:START_DATETIME"`
	EndDate         *time.Time `gorm:"column:END_DATETIME"`
	CreatedDatetime time.Time  `gorm:"not null;column:CREATED_DATETIME"`
	UpdatedDatetime time.Time  `gorm:"column:UPDATED_DATETIME"`
	CreatedBy       string     `gorm:"not null;column:CREATED_BY"`
	UpdateBy        string     `gorm:"column:UPDATED_BY"`
}
