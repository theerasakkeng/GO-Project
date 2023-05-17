package main

import "project-study/routes"

func main() {
	routes.InitRoute()
	routes.Router.Run()
}
