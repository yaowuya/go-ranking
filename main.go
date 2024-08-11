package main

import (
	"go-ranking/router"
)

func main() {
	r := router.Router()

	err := r.Run(":9999")
	if err != nil {
		return
	}

}
