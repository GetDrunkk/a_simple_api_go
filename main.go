package main

import (
	_ "a_simple_api_go/database"
	"a_simple_api_go/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run(":8080")
}
