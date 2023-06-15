/*
CMS Source Code

Testing ExtractApiService

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

func Test_source_code_ExtractApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExtractApiService ExtractByPath", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var path string

		httpRes, err := apiClient.ExtractApi.ExtractByPath(context.Background(), path).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
