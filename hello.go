package controller

import (
	"github.com/goadesign/goa"
	"github.com/tleyden/serverless-forms/goa-generated/app"
)

// HelloController implements the hello resource.
type HelloController struct {
	*goa.Controller
}

// NewHelloController creates a hello controller.
func NewHelloController(service *goa.Service) *HelloController {
	return &HelloController{Controller: service.NewController("HelloController")}
}

// Show runs the show action.
func (c *HelloController) Show(ctx *app.ShowHelloContext) error {
	// HelloController_Show: start_implement

	// Put your logic here

	// HelloController_Show: end_implement
	res := &app.Hello{}

	res.Hello = ctx.WhatToSay

	return ctx.OK(res)
}
