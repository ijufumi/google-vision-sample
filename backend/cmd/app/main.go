package main

import (
	"github.com/ijufumi/google-vision-sample/pkg/container"
	"github.com/ijufumi/google-vision-sample/pkg/http/router"
)

func main() {
	c := container.NewContainer()
	r := router.NewRouter(c)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
