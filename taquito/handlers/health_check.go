package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/cmh2166/elag18apis/taquito/generated/models"
	"github.com/cmh2166/elag18apis/taquito/generated/restapi/operations"
)

// NewHealthCheck will return the service health
func NewHealthCheck() operations.HealthCheckHandler {
	return &healthCheck{}
}

type healthCheck struct{}

// Handle the health check request
func (d *healthCheck) Handle(params operations.HealthCheckParams) middleware.Responder {
	return operations.NewHealthCheckOK().WithPayload(&models.HealthCheckResponse{Status: "OK"})
}
