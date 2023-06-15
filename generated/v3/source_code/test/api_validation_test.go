/*
CMS Source Code

Testing ValidationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package source_code

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_source_code_ValidationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ValidationApiService Validate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var path string

		resp, httpRes, err := apiClient.ValidationApi.Validate(context.Background(), path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
