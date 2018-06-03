package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/cmh2166/elag18apis/taquito/authorization"
	"github.com/cmh2166/elag18apis/taquito/datautils"
	"github.com/cmh2166/elag18apis/taquito/db"
	"github.com/cmh2166/elag18apis/taquito/generated/restapi/operations"
)

// NewRetrieveResource will query DynamoDB with ID for Resource JSON
func NewRetrieveResource(database db.Database) operations.RetrieveResourceHandler {
	return &retrieveResource{database: database}
}

// retrieveResource handles a request for finding & returning an entry
type retrieveResource struct {
	database db.Database
}

// Handle the retrieve resource request
func (d *retrieveResource) Handle(params operations.RetrieveResourceParams, agent *authorization.Agent) middleware.Responder {
	var resource *datautils.Resource
	var err error
	if params.Version != nil {
		resource, err = d.database.RetrieveVersion(params.ID, params.Version)
	} else {
		resource, err = d.database.RetrieveLatest(params.ID)
	}

	if err != nil {
		if _, ok := err.(*db.RecordNotFound); ok {
			return operations.NewRetrieveResourceNotFound()
		}
		panic(err)
	}

	authService := authorization.NewService(agent)
	if !authService.CanRetrieveResource(resource) {
		log.Printf("Agent %s is not permitted to retrieve this resource %s", agent, params.ID)
		return operations.NewDepositResourceUnauthorized()
	}

	return operations.NewRetrieveResourceOK().WithPayload(resource.JSON)
}
