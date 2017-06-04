//go:generate goagen bootstrap -d github.com/tleyden/serverless-forms/design

package main

import (
	"net/http"

	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	controller "github.com/tleyden/serverless-forms/controllers"
	"github.com/tleyden/serverless-forms/goa-generated/app"
)

func createGoaService() *goa.Service {

	// Create service
	service := goa.New("HelloServerlessGoa")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "hello" controller
	c := controller.NewHelloController(service)
	app.MountHelloController(service, c)

	return service
}

func main() {

	service := createGoaService()

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

// Handle is the exported handler called by AWS Lambda.
var Handle apigatewayproxy.Handler

func init() {

	ln := net.Listen()

	// Amazon API Gateway Binary support out of the box.
	Handle = apigatewayproxy.New(ln, nil).Handle

	service := createGoaService()

	// Any Go framework complying with the Go http.Handler interface can be used.
	// This includes, but is not limited to, Vanilla Go, Gin, Echo, Gorrila, etc.
	go http.Serve(ln, service.Mux)

}
