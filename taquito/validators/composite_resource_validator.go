package validators

import (
	"github.com/cmh2166/elag18apis/taquito/datautils"
	"github.com/cmh2166/elag18apis/taquito/generated/models"
)

type compositeResourceValidator struct {
	validators []ResourceValidator
}

// NewCompositeResourceValidator creates a new validator that combines several sub-validators
func NewCompositeResourceValidator(validators []ResourceValidator) ResourceValidator {
	return &compositeResourceValidator{validators: validators}
}

func (d *compositeResourceValidator) ValidateResource(resource *datautils.Resource) *models.ErrorResponseErrors {
	for _, validator := range d.validators {
		if errors := validator.ValidateResource(resource); errors != nil {
			return errors
		}
	}
	return nil
}
