package registry

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Exca-DK/pegism/core/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockService struct {
	status error
}

func (*mockService) Start() error {
	return nil
}

func (*mockService) Stop() error {
	return nil
}

func (m *mockService) Status() error {
	return m.status
}

type secondMockService struct {
	status error
}

func (*secondMockService) Start() error {
	return nil
}

func (*secondMockService) Stop() error {
	return nil
}

func (s *secondMockService) Status() error {
	return s.status
}

func TestRegisterService_Twice(t *testing.T) {
	registry := &ServiceRegistry{
		services: make(map[reflect.Type]Service),
		log:      log.Root(),
	}

	m := &mockService{}
	require.NoError(t, registry.RegisterService(m), "Failed to register first service")

	// Checks if first service was indeed registered.
	require.Equal(t, 1, len(registry.serviceTypes))
	assert.ErrorContains(t, registry.RegisterService(m), "service already exists")
}

func TestRegisterService_Different(t *testing.T) {
	registry := &ServiceRegistry{
		services: make(map[reflect.Type]Service),
		log:      log.Root(),
	}

	m := &mockService{}
	s := &secondMockService{}
	require.NoError(t, registry.RegisterService(m), "Failed to register first service")
	require.NoError(t, registry.RegisterService(s), "Failed to register second service")

	require.Equal(t, 2, len(registry.serviceTypes))

	_, exists := registry.services[reflect.TypeOf(m)]
	assert.Equal(t, true, exists, "service of type %v not registered", reflect.TypeOf(m))

	_, exists = registry.services[reflect.TypeOf(s)]
	assert.Equal(t, true, exists, "service of type %v not registered", reflect.TypeOf(s))
}

func TestFetchService_OK(t *testing.T) {
	registry := &ServiceRegistry{
		services: make(map[reflect.Type]Service),
		log:      log.Root(),
	}

	m := &mockService{}
	require.NoError(t, registry.RegisterService(m), "Failed to register first service")

	assert.ErrorContains(
		t,
		registry.FetchService(*m),
		"input must be of pointer type, received value type instead",
	)

	var s *secondMockService
	assert.ErrorContains(t, registry.FetchService(&s), "unknown service")

	var m2 *mockService
	require.NoError(t, registry.FetchService(&m2), "Failed to fetch service")
	require.Equal(t, m, m2)
}

func TestServiceStatus_OK(t *testing.T) {
	registry := &ServiceRegistry{
		services: make(map[reflect.Type]Service),
		log:      log.Root(),
	}

	m := &mockService{}
	require.NoError(t, registry.RegisterService(m), "Failed to register first service")

	s := &secondMockService{}
	require.NoError(t, registry.RegisterService(s), "Failed to register first service")

	m.status = errors.New("something bad has happened")
	s.status = errors.New("woah, horsee")

	statuses := registry.Statuses()

	assert.ErrorContains(t, statuses[reflect.TypeOf(m)], "something bad has happened")
	assert.ErrorContains(t, statuses[reflect.TypeOf(s)], "woah, horsee")
}
