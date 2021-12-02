package main

import (
	"Server/Routes"
)

func main() {
	r := Routes.SetupRouter()
	r.Run(":8081")
}
