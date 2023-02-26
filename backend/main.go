package main

import (
	"github.com/ijufumi/google-vision-sample/cmd/app"
)

func main() {
	err := app.RunApp()
	if err != nil {
		panic(err)
	}
}
