package app

import (
	"github.com/ijufumi/google-vision-sample/internal/container"
	"github.com/ijufumi/google-vision-sample/internal/http/router"
)

func RunApp() error {
	c := container.NewContainer()
	r := router.NewRouter(c)
	return r.Run()
}
