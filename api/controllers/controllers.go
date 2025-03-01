package controllers

import "github.com/JscorpTech/jst-go/bootstrap"

type Controller struct {
	App *bootstrap.App
}

func NewController(app *bootstrap.App) *Controller {
	return &Controller{
		App: app,
	}
}
