package validators

import (
	"github.com/cmh2166/elag18apis/taquito/datautils"
	"github.com/cmh2166/elag18apis/taquito/generated/models"
)

// ResourceValidator is the interface for validators that check the resources format
type ResourceValidator interface {
	ValidateResource(resource *datautils.Resource) *models.ErrorResponseErrors
}
