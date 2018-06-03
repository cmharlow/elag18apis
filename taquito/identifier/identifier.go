package identifier

import (
	"log"

	"github.com/cmh2166/elag18apis/taquito/config"
	"github.com/cmh2166/elag18apis/taquito/datautils"
)

// Service the interface for a service that mints identifiers
type Service interface {
	Mint(resource *datautils.Resource) (string, error)
}

func NewService(config *config.Config) Service {
	return &TypeSpecificIDService{
		remoteService: remoteService(config),
		localService:  NewUUIDService(),
	}
}

func remoteService(config *config.Config) Service {
	if config.IdentifierServiceHost == "" {
		log.Println("IDENTIFIER_SERVICE_HOST is not set, so using UUID service")
		return NewUUIDService()
	} else {
		return NewRemoteIdentifierService(config)
	}
}
