package db

import (
	"distributed-chat/master/structs"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MinionUserList = structs.MinionUserList
type MinionList = structs.MinionList

func InitDb() gorm.DB {
	db, err := gorm.Open(sqlite.Open("production.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return *db
}

func CreateDbFromSchema(db gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&MinionUserList{})
	db.AutoMigrate(&MinionList{})
}

func RetrieveUserByName(db gorm.DB, username string) (MinionUserList, error) {
	var minionUserList MinionUserList
	err := db.First(&minionUserList, "username = ?", username).Error
	return minionUserList, err
}

func CreateUser(db gorm.DB, minionUserList MinionUserList) (MinionUserList, error) {
	result := db.Create(&minionUserList)
	if result.Error != nil {
		fmt.Println("Error while creating user")
		return minionUserList, result.Error
	}
	return minionUserList, nil
}

func CreateMinion(db gorm.DB, minionList MinionList) (MinionList, error) {
	result := db.Create(&minionList)
	if result.Error != nil {
		fmt.Println("Error while creating minion")
		return minionList, result.Error
	}
	return minionList, nil
}

func RetrieveAllMinions(db gorm.DB) []MinionList {
	var minions []MinionList
	result := db.Find(&minions)
	if result.Error != nil {
		fmt.Println("Error while retrieving minions")
	}
	return minions
}

func DeleteUser(db gorm.DB, minionUserLists []MinionUserList) {
	db.Delete(&minionUserLists)
}

func UpdateUser(db gorm.DB, minionUserList MinionUserList) MinionUserList {
	db.Save(&minionUserList)
	return minionUserList
}

func RetrieveAllUsersOfMinion(db gorm.DB, minionUrlIdentifier string) []MinionUserList {
	var minionUserList []MinionUserList
	db.Find(&minionUserList, "minion_url_identifier = ?", minionUrlIdentifier)
	return minionUserList
}
