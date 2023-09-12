package khulnasoft_test

import (
	"testing"

	khulnasoft "github.com/khulnasoft-lab/khulnasoft-go"
	"github.com/stretchr/testify/assert"
)

func TestError_CreateErrors(t *testing.T) {
	baseErr := &khulnasoft.Error{
		StatusCode: 400,
		ErrorCodes: []int{10000},
	}

	requestErr := khulnasoft.NewRequestError(baseErr)
	assert.True(t, requestErr.InternalErrorCodeIs(10000))
	limitError := khulnasoft.NewRatelimitError(baseErr)
	assert.True(t, limitError.InternalErrorCodeIs(10000))
	svcErr := khulnasoft.NewServiceError(baseErr)
	assert.True(t, svcErr.InternalErrorCodeIs(10000))
	authErr := khulnasoft.NewAuthenticationError(baseErr)
	assert.True(t, authErr.InternalErrorCodeIs(10000))
	authzErr := khulnasoft.NewAuthorizationError(baseErr)
	assert.True(t, authzErr.InternalErrorCodeIs(10000))
	notFoundErr := khulnasoft.NewNotFoundError(baseErr)
	assert.True(t, notFoundErr.InternalErrorCodeIs(10000))
}
