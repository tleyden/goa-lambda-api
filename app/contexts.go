// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "HelloServerlessGoa": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/tleyden/serverless-forms/design
// --out=$(GOPATH)/src/github.com/tleyden/serverless-forms/goa-generated
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// ShowHelloContext provides the hello show action context.
type ShowHelloContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	WhatToSay string
}

// NewShowHelloContext parses the incoming request URL and body, performs validations and creates the
// context used by the hello controller show action.
func NewShowHelloContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowHelloContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowHelloContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramWhatToSay := req.Params["whatToSay"]
	if len(paramWhatToSay) > 0 {
		rawWhatToSay := paramWhatToSay[0]
		rctx.WhatToSay = rawWhatToSay
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowHelloContext) OK(r *Hello) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.hello+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowHelloContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
