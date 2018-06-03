package handlers

import (
	"log"

	"github.com/cmh2166/elag18apis/taquito/authorization"
	"github.com/cmh2166/elag18apis/taquito/db"
	"github.com/cmh2166/elag18apis/taquito/generated/restapi"
	"github.com/cmh2166/elag18apis/taquito/generated/restapi/operations"
	"github.com/cmh2166/elag18apis/taquito/identifier"
	"github.com/cmh2166/elag18apis/taquito/validators"
	"github.com/go-openapi/loads"
)

// BuildAPI create new service API
func BuildAPI(database db.Database, identifierService identifier.Service) *operations.TacoAPI {
	api := operations.NewTacoAPI(swaggerSpec())
	api.RemoteUserAuth = func(identifier string) (*authorization.Agent, error) {
		return &authorization.Agent{Identifier: identifier}, nil
	}
	api.RetrieveResourceHandler = NewRetrieveResource(database)
	api.DeleteResourceHandler = NewDeleteResource(database)
	api.DepositResourceHandler = NewDepositResource(database, depositResourceValidator(database), identifierService)
	api.UpdateResourceHandler = NewUpdateResource(database, updateResourceValidator(database))
	api.HealthCheckHandler = NewHealthCheck()
	return api
}

// Builds the validator for deposit resource requests
func depositResourceValidator(database db.Database) validators.ResourceValidator {
	return validators.NewCompositeResourceValidator(
		[]validators.ResourceValidator{
			validators.NewDepositResourceValidator(database),
			structuralValidator(database),
		})
}

// Builds the validator for update resource requests
func updateResourceValidator(database db.Database) validators.ResourceValidator {
	return validators.NewCompositeResourceValidator(
		[]validators.ResourceValidator{
			validators.NewUpdateResourceValidator(database),
			structuralValidator(database),
		})
}

// Builds the validator for structural validations.
// This is suitable for both create and update requests
func structuralValidator(database db.Database) validators.ResourceValidator {
	return validators.NewCompositeResourceValidator(
		[]validators.ResourceValidator{
			validators.NewFilesetStructuralValidator(database),
			validators.NewDROStructuralValidator(database),
			validators.NewCollectionStructuralValidator(database),
			validators.NewSequenceValidator(),
		})
}

func swaggerSpec() *loads.Document {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	return swaggerSpec
}
