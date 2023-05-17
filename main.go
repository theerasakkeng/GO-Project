package main

import (
	"project-study/routes"
	"time"
)

func main() {
	initTimeZone()
	routes.InitRoute()
	routes.Router.Run()
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
