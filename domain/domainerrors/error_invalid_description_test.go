package domainerrors_test

import (
	"testing"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_invalid_description_error(t *testing.T) {
	err := domainerrors.NewErrorInvalidDescription("test")

	require.NotNil(t, err)
}
