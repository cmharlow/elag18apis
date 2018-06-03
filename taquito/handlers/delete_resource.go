package handlers

import (
	"github.com/cmh2166/elag18apis/taquito/db"
	"github.com/cmh2166/elag18apis/taquito/generated/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// NewDeleteResource -- Accepts requests to remove a resource and pushes them to Kinesis.
func NewDeleteResource(repository db.Database) operations.DeleteResourceHandler {
	return &deleteResourceEntry{repository: repository}
}

type deleteResourceEntry struct {
	repository db.Database
}

// Handle the delete entry request. If the resource being deleted is a file,
// Also delete the associated binary from s3
func (d *deleteResourceEntry) Handle(params operations.DeleteResourceParams) middleware.Responder {
	if err := d.DeleteAllVersions(params.ID); err != nil {
		panic(err)
	}

	return operations.NewDeleteResourceNoContent()
}

// DeleteAllVersions removes all versions with the given external id
func (d *deleteResourceEntry) DeleteAllVersions(externalID string) error {
	resource, err := d.repository.RetrieveLatest(externalID)
	if err != nil {
		if _, ok := err.(*db.RecordNotFound); !ok {
			panic(err)
		}
	}
	// Delete all versions of the resource
	for resource != nil {
		if err = d.repository.DeleteByID(resource.ID()); err != nil {
			panic(err)
		}

		// retrieve the next resource
		resource, err = d.repository.RetrieveLatest(externalID)
		if err != nil {
			if _, ok := err.(*db.RecordNotFound); !ok {
				panic(err)
			}
		}

	}
	return nil
}
