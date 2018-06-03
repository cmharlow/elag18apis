package handlers

import (
	"log"
	"time"

	"github.com/cmh2166/elag18apis/taquito/authorization"
	"github.com/cmh2166/elag18apis/taquito/datautils"
	"github.com/cmh2166/elag18apis/taquito/db"
	"github.com/cmh2166/elag18apis/taquito/generated/models"
	"github.com/cmh2166/elag18apis/taquito/generated/restapi/operations"
	"github.com/cmh2166/elag18apis/taquito/validators"
	"github.com/go-openapi/runtime/middleware"

	"github.com/cmh2166/elag18apis/taquito/identifier"
)

// NewDepositResource -- Accepts requests to create resource and pushes them to Kinesis.
func NewDepositResource(database db.Database, validator validators.ResourceValidator, identifierService identifier.Service) operations.DepositResourceHandler {
	return &depositResource{
		database:          database,
		validator:         validator,
		identifierService: identifierService,
	}
}

type depositResource struct {
	database          db.Database
	validator         validators.ResourceValidator
	identifierService identifier.Service
}

// Handle the create resource request
func (d *depositResource) Handle(params operations.DepositResourceParams, agent *authorization.Agent) middleware.Responder {
	resource := datautils.NewResource(params.Payload.(map[string]interface{}))
	if errors := d.validator.ValidateResource(resource); errors != nil {
		return operations.NewDepositResourceUnprocessableEntity().
			WithPayload(&models.ErrorResponse{Errors: *errors})
	}

	authService := authorization.NewService(agent)
	if !authService.CanCreateResourceOfType(resource.Type()) {
		log.Printf("Agent %s is not permitted to create a resource of type %s", agent, resource.Type())
		return operations.NewDepositResourceUnauthorized()
	}

	externalID, err := d.identifierService.Mint(resource)
	if err != nil {
		panic(err)
	}

	uuid, err := identifier.NewUUIDService().Mint(resource)
	if err != nil {
		panic(err)
	}

	resource = resource.
		WithID(uuid).
		WithExternalIdentifier(externalID).
		WithVersion(1)
	(*resource.Administrative())["created"] = time.Now().UTC().Format(time.RFC3339)

	if err := d.database.Insert(resource); err != nil {
		panic(err)
	}

	response := datautils.JSONObject{"id": externalID}
	return operations.NewDepositResourceCreated().WithPayload(response)
}
