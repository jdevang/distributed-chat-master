package main

import (
	"distributed-chat/master/db"
	"distributed-chat/master/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MinionUserList = structs.MinionUserList
type MinionList = structs.MinionList
type HTTPStatusMessage = structs.HTTPStatusMessage

var dbInstance = db.InitDb()

func main() {
	router := gin.Default()
	db.CreateDbFromSchema(dbInstance)

	router.POST("/getMinionUrlIdentifier", getMinionUrlIdentifier)

	router.POST("/registerUser", registerUser)

	router.POST("/registerMinion", registerMinion)

	router.GET("/alive", alive)

	router.SetTrustedProxies(nil)

	router.Run(":8080")
}

func getMinionUrlIdentifier(c *gin.Context) {
	var minionUserList MinionUserList
	err := c.BindJSON(&minionUserList)
	if err != nil {
		fmt.Println("Error in reading json")
		c.IndentedJSON(http.StatusBadRequest, HTTPStatusMessage{Message: "faulty request"})
		return
	}

	minionUserList, err = db.RetrieveUserByName(dbInstance, minionUserList.Username)

	if err != nil {
		fmt.Println("User doesn't exist")
		c.IndentedJSON(http.StatusConflict, HTTPStatusMessage{Message: "Username can't be found"})
		return
	}

	c.IndentedJSON(http.StatusOK, HTTPStatusMessage{MinionUrlIdentifier: minionUserList.MinionUrlIdentifier})
}

func registerUser(c *gin.Context) {
	var minionUserList MinionUserList
	err := c.BindJSON(&minionUserList)
	if err != nil {
		fmt.Println("Error in reading json")
		c.IndentedJSON(http.StatusBadRequest, HTTPStatusMessage{Message: "faulty request"})
		return
	}

	_, err = db.CreateUser(dbInstance, minionUserList)

	if err != nil {
		fmt.Println("User couldn't be created")
		c.IndentedJSON(http.StatusConflict, HTTPStatusMessage{Message: "Username can't be added to list"})
		return
	}
	c.IndentedJSON(http.StatusCreated, HTTPStatusMessage{Message: "Username added to list"})
}

func registerMinion(c *gin.Context) {
	var minionList MinionList
	err := c.BindJSON(&minionList)
	if err != nil {
		fmt.Println("Error in reading json")
		c.IndentedJSON(http.StatusBadRequest, HTTPStatusMessage{Message: "faulty request"})
		return
	}

	_, err = db.CreateMinion(dbInstance, minionList)

	if err != nil {
		fmt.Println("Minion name exists")
		c.IndentedJSON(http.StatusConflict, HTTPStatusMessage{Message: "Minion name exists"})
		return
	}
	c.IndentedJSON(http.StatusCreated, HTTPStatusMessage{Message: "Minion added to list"})
}

func alive(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, HTTPStatusMessage{Message: "I'm Alive!"})
}

// TODO: client propagation in event of minion failure
// TODO: Master fault tolerance
