package main

import (
	"github.com/SnrMatt/AmmorApparel/router"
)

func main() {
	e := router.New()

	e.Logger.Fatal(e.Start(":1323"))
}
