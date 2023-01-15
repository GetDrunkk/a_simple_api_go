package controllers

import (
	"a_simple_api_go/database"
	"encoding/json"
	"io"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Return_data struct {
	Postal_code        int     `json:"postal_code"`
	Hit_count          int     `json:"hit_count"`
	Address            string  `json:"address"`
	Tokyo_sta_distance float64 `json:"tokyo_sta_distance"`
}

func Test(c *gin.Context) {
	c.String(200, "Hello World")
}

func Postal(c *gin.Context) {
	code := c.Query("postal_code")
	url := "https://geoapi.heartrails.com/api/json?method=searchByPostal&postal=" + code
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	defer response.Body.Close()
	json_data, _ := io.ReadAll(response.Body) //read the json part of the reponse
	var res map[string]interface{}
	json.Unmarshal(json_data, &res)
	//fmt.Print(res["response"])  check the result

	var re_data Return_data
	all_locs := res["response"].(map[string]interface{})["location"].([]interface{})

	var x float64
	var y float64
	var dis float64
	var dis_n float64
	var add string
	var add_fin string
	dis = 0.0
	for i, k := range all_locs {
		x, _ = strconv.ParseFloat(k.(map[string]interface{})["x"].(string), 64)
		y, _ = strconv.ParseFloat(k.(map[string]interface{})["y"].(string), 64)
		add = k.(map[string]interface{})["prefecture"].(string) + k.(map[string]interface{})["city"].(string) + k.(map[string]interface{})["town"].(string)
		if i != 0 {
			add_fin = Take_address(add, add_fin)
		} else {
			add_fin = add
		}
		dis_n = Cal_dis(x, y)
		if dis_n > dis {
			dis = dis_n
		}
	}
	int_code, _ := strconv.Atoi(code)
	re_data.Postal_code = int_code
	re_data.Hit_count = len(all_locs)
	re_data.Address = add_fin
	re_data.Tokyo_sta_distance = float64(int64(dis*10+0.5)) / 10
	//fmt.Print(re_data)
	database.Insert(int_code)
	c.JSON(200, re_data)
}

func Cal_dis(x, y float64) (d float64) {
	x_t := 139.7673068
	y_t := 35.6809591
	d = math.Sqrt(math.Pow((x-x_t)*math.Cos(math.Pi*(y+y_t)/360), 2)+math.Pow((y-y_t), 2)) * math.Pi * 6371 / 180

	return
}

func Take_address(x, y string) (res string) {
	res = ""
	x_r := []rune(x)
	y_r := []rune(y)
	if len(x_r) < len(y_r) {
		for i := 0; i < len(x_r); i++ {
			if x_r[i] == y_r[i] {
				res += string(x_r[i])
			} else {
				return
			}
		}
		return
	} else {
		for i := 0; i < len(y_r); i++ {
			if x_r[i] == y_r[i] {
				res += string(y_r[i])
			} else {
				return
			}
		}
		return
	}
}

func Access_log(c *gin.Context) {
	res := database.Read_log()
	c.JSON(200, res)
}
