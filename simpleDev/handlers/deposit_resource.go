package handlers

import (
	"time"

	"github.com/cmh2166/elag18apis/simpleDev/generated/models"
	"github.com/cmh2166/elag18apis/simpleDev/generated/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// NewDepositResource -- Accepts requests to create resource.
func NewDepositResource() operations.DepositResourceHandler {
	return &depositResource{}
}

type depositResource struct{}

// Handle the create resource request
func (d *depositResource) Handle(params operations.DepositResourceParams, agent *models.Agent) middleware.Responder {
	// Grab the data passed in from the call
	resource := params.Payload

	// Mint an identifier (UUID)
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	// Add administrative metadata
	var identifier strfmt.UUID
	identifier = strfmt.UUID(uuid.String())
	resource.Identifier = identifier
	resource.Depositor = agent
	var version int64
	version = int64(1)
	resource.Version = version
	var datestamp strfmt.DateTime
	datestamp = strfmt.DateTime(time.Now().UTC())
	resource.Administrative.Created = datestamp
	resource.Administrative.LastUpdated = datestamp

	// Build the API response
	response := &models.ResourceResponse{Identifier: &identifier, Data: &resource}
	return operations.NewDepositResourceCreated().WithPayload(response)
}
