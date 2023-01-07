package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		response, err := http.Get("https://geoapi.heartrails.com/api/json?method=searchByPostal&postal=5016121")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		defer response.Body.Close()
		b, _ := io.ReadAll(response.Body) //read the json part of the reponse
		var res map[string]interface{}
		json.Unmarshal(b, &res)
		//fmt.Print(res["response"])  check the result
		c.JSON(200, res)
	})

	router.Run(":8080")
}
