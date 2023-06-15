/*
Custom Workflow Actions

Testing FunctionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package actions

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_actions_FunctionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FunctionsApiService FunctionsArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var definitionId string
		var functionType string
		var functionId string
		var appId int32

		httpRes, err := apiClient.FunctionsApi.FunctionsArchive(context.Background(), definitionId, functionType, functionId, appId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsApiService FunctionsArchiveByType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var definitionId string
		var functionType string
		var appId int32

		httpRes, err := apiClient.FunctionsApi.FunctionsArchiveByType(context.Background(), definitionId, functionType, appId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsApiService FunctionsCreateOrReplace", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var definitionId string
		var functionType string
		var functionId string
		var appId int32

		resp, httpRes, err := apiClient.FunctionsApi.FunctionsCreateOrReplace(context.Background(), definitionId, functionType, functionId, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsApiService FunctionsCreateOrReplaceByType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var definitionId string
		var functionType string
		var appId int32

		resp, httpRes, err := apiClient.FunctionsApi.FunctionsCreateOrReplaceByType(context.Background(), definitionId, functionType, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsApiService FunctionsGetByID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var definitionId string
		var functionType string
		var functionId string
		var appId int32

		resp, httpRes, err := apiClient.FunctionsApi.FunctionsGetByID(context.Background(), definitionId, functionType, functionId, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsApiService FunctionsGetByType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var definitionId string
		var functionType string
		var appId int32

		resp, httpRes, err := apiClient.FunctionsApi.FunctionsGetByType(context.Background(), definitionId, functionType, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsApiService FunctionsGetPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var definitionId string
		var appId int32

		resp, httpRes, err := apiClient.FunctionsApi.FunctionsGetPage(context.Background(), definitionId, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
