package main

//go:generate go run generate/generate.go

import (
	"log"

	"github.com/cmh2166/elag18apis/simpleDev/generated/models"
	"github.com/cmh2166/elag18apis/simpleDev/generated/restapi"
	"github.com/cmh2166/elag18apis/simpleDev/generated/restapi/operations"
	"github.com/cmh2166/elag18apis/simpleDev/handlers"
	"github.com/go-openapi/loads"
)

func main() {
	server := createServer(8080)
	defer server.Shutdown()

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func createServer(port int) *restapi.Server {
	api := operations.NewTaquitoAPI(swaggerSpec())
	api.RemoteUserAuth = func(identifier string) (*models.Agent, error) {
		return &models.Agent{SunetID: &identifier}, nil
	}
	api.DepositResourceHandler = handlers.NewDepositResource()
	server := restapi.NewServer(api)

	// set the port this service will be run on
	server.Port = port
	return server
}

func swaggerSpec() *loads.Document {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	return swaggerSpec
}
