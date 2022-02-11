package domainerrors_test

import (
	"testing"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_internal_server_error(t *testing.T) {
	err := domainerrors.NewErrorInternalServer("process", "error message")
	require.NotNil(t, err)
}

func Test_should_be_create_an_internal_server_error_text(t *testing.T) {
	err := domainerrors.NewErrorInternalServer("process", `"error message"`)

	require.Equal(t, `{"process": "error message"}`, err.Error())
}
