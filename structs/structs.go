package structs

import (
	"gorm.io/gorm"
)

type MinionUserList struct {
	gorm.Model
	Username            string `gorm:"uniqueIndex:idx_username"`
	Password            string
	ClientUrlIdentifier string
	MinionUrlIdentifier string
}

type MinionList struct {
	gorm.Model
	MinionName          string `gorm:"uniqueIndex:idx_minion_name"`
	MinionUrlIdentifier string `gorm:"uniqueIndex:idx_minion_url"`
}

type HTTPStatusMessage struct {
	Message             string
	MinionUrlIdentifier string
}
