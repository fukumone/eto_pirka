package main

import "github.com/t-fukui/eto_pirka/routes"

func main() {
	router := routes.Init()

	router.Run(":3000")
}
