package registry

import (
	"fmt"
	"reflect"

	"github.com/Exca-DK/pegism/core/log"
)

type Service interface {
	// Start spawns any goroutines required by the service.
	Start() error
	// Stop terminates all goroutines belonging to the service,
	// blocking until they are all terminated.
	Stop() error
	// Status returns error if the service is not considered healthy.
	Status() error
}

type ServiceRegistry struct {
	services     map[reflect.Type]Service // map of types to services.
	serviceTypes []reflect.Type           // keep an ordered slice of registered service types.
	log          log.Logger
}

// NewServiceRegistry starts a registry instance for convenience
func NewServiceRegistry(applicationName string) *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[reflect.Type]Service),
		log:      log.Root().With(applicationName, "registry"),
	}
}

// StartAll initialized each service in order of registration.
func (s *ServiceRegistry) Start() {
	s.log.Debug(
		"Starting services",
		"amount", s.services,
	)
	errChs := make(chan error, len(s.serviceTypes))
	for _, kind := range s.serviceTypes {
		s.log.Debug("Starting service", "type", kind.Elem().Name())
		go func(kind reflect.Type) {
			err := s.services[kind].Start()
			if err != nil {
				s.log.Fatal(
					"Failed starting service",
					"type", kind.Elem().Name(),
					"err", err,
				)
			}
		}(kind)
	}

	for i := 0; i < len(s.serviceTypes); i++ {
		err := <-errChs
		if err != nil {
			s.log.Info("Stopping service due to bootstrap error")
			s.Stop()
		}
	}
}

// StopAll ends every service in reverse order of registration, logging a
// panic if any of them fail to stop.
func (s *ServiceRegistry) Stop() {
	for i := len(s.serviceTypes) - 1; i >= 0; i-- {
		kind := s.serviceTypes[i]
		service := s.services[kind]
		if err := service.Stop(); err != nil {
			s.log.Error(
				"Could not stop service:",
				"type", kind.Elem().Name(),
				"err", err,
			)
		}
	}
}

// Statuses returns a map of Service type -> error. The map will be populated
// with the results of each service.Status() method call.
func (s *ServiceRegistry) Statuses() map[reflect.Type]error {
	m := make(map[reflect.Type]error, len(s.serviceTypes))
	for _, kind := range s.serviceTypes {
		m[kind] = s.services[kind].Status()
	}
	return m
}

// RegisterService appends a service constructor function to the service
// registry.
func (s *ServiceRegistry) RegisterService(service Service) error {
	kind := reflect.TypeOf(service)
	if kind.Kind() != reflect.Ptr {
		return fmt.Errorf("input must be of pointer type, received value type instead: %T", service)
	}
	if _, exists := s.services[kind]; exists {
		return fmt.Errorf("service already exists: %v", kind)
	}
	s.log.Debug("registering service", "type", kind.Elem().Name())
	s.services[kind] = service
	s.serviceTypes = append(s.serviceTypes, kind)
	return nil
}

// FetchService takes in a struct pointer and sets the value of that pointer
// to a service currently stored in the service registry. This ensures the input argument is
// set to the right pointer that refers to the originally registered service.
func (s *ServiceRegistry) FetchService(service interface{}) error {
	if reflect.TypeOf(service).Kind() != reflect.Ptr {
		return fmt.Errorf("input must be of pointer type, received value type instead: %T", service)
	}
	element := reflect.ValueOf(service).Elem()
	if running, ok := s.services[element.Type()]; ok {
		element.Set(reflect.ValueOf(running))
		return nil
	}
	return fmt.Errorf("unknown service: %T", service)
}
