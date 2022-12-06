package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
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
	var content, err = json.Marshal(profiledata)
	content, err = ioutil.ReadFile("profile.json")
	if err != nil {
		// 	log.Fatal(err)
		fmt.Println("File Does Not Exist")
	} else {
		err = json.Unmarshal(content, &profiledata)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", content)
	}

}

func postProfile(c *gin.Context) {
	var newProfile profile

	if err := c.BindJSON(&newProfile); err != nil {
		return
	}

	finalprofile := append(profiledata, newProfile)

	c.IndentedJSON(http.StatusCreated, newProfile)

	finalJson, err := json.Marshal(finalprofile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

	err = ioutil.WriteFile("profile.json", finalJson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := gin.Default()
	router.GET("/profile", getProfile)
	router.POST("/profile", postProfile)
	router.Run("localhost:5050")
}
