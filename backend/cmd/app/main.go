package app

import (
	"github.com/ijufumi/google-vision-sample/pkg/container"
	"github.com/ijufumi/google-vision-sample/pkg/http/router"
)

func RunApp() error {
	c := container.NewContainer()
	r := router.NewRouter(c)
	return r.Run()
}
