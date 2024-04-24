package main

import (
	"restApi/router"
)

func main() {
	r := router.Router()
	r.Run(":9999")
}
