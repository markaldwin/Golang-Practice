package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type profile struct {
	ID      string `json: "id"`
	Name    string `json: "name"`
	Address string `json: "address"`
	Motto   string `json: "motto"`
}

var profiledata = []profile{}

func getProfile(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, profiledata)
}

func postProfile(c *gin.Context) {
	var newProfile profile

	if err := c.BindJSON(&newProfile); err != nil {
		return
	}

	profiledata = append(profiledata, newProfile)

	c.IndentedJSON(http.StatusCreated, newProfile)
}

func main() {
	router := gin.Default()
	router.GET("/profile", getProfile)
	router.POST("/profile", postProfile)
	router.Run("localhost:5050")
}
