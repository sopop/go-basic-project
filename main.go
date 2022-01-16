package main

import (
	"demo/router"
)

func main() {
	router := router.Strat()
	router.Run(":8089")
}
