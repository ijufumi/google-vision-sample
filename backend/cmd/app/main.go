package app

import (
	"github.com/ijufumi/google-vision-sample/internal/common/container"
	"github.com/ijufumi/google-vision-sample/internal/presentations/http/router"
)

func RunApp() error {
	c := container.NewContainer()
	r := router.NewRouter(c)
	return r.Run()
}
