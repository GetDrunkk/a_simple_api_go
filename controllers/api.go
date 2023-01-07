package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.String(200, "Hello World")
}

func Postal(c *gin.Context) {
	response, err := http.Get("https://geoapi.heartrails.com/api/json?method=searchByPostal&postal=5016121")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	defer response.Body.Close()
	json_data, _ := io.ReadAll(response.Body) //read the json part of the reponse
	var res map[string]interface{}
	json.Unmarshal(json_data, &res)
	//fmt.Print(res["response"])  check the result
	c.JSON(200, res)
}
