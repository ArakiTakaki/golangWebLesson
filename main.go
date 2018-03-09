package main

import (
	"./router"
)

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
