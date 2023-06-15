/*
HubDB endpoints

Testing RowsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package hubdb

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_hubdb_RowsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RowsApiService CloneDraftTableRow", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableIdOrName string
		var rowId string

		resp, httpRes, err := apiClient.RowsApi.CloneDraftTableRow(context.Background(), tableIdOrName, rowId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RowsApiService CreateTableRow", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableIdOrName string

		resp, httpRes, err := apiClient.RowsApi.CreateTableRow(context.Background(), tableIdOrName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RowsApiService GetDraftTableRowByID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableIdOrName string
		var rowId string

		resp, httpRes, err := apiClient.RowsApi.GetDraftTableRowByID(context.Background(), tableIdOrName, rowId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RowsApiService GetTableRow", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableIdOrName string
		var rowId string

		resp, httpRes, err := apiClient.RowsApi.GetTableRow(context.Background(), tableIdOrName, rowId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RowsApiService GetTableRows", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableIdOrName string

		resp, httpRes, err := apiClient.RowsApi.GetTableRows(context.Background(), tableIdOrName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RowsApiService PurgeDraftTableRow", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableIdOrName string
		var rowId string

		httpRes, err := apiClient.RowsApi.PurgeDraftTableRow(context.Background(), tableIdOrName, rowId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RowsApiService ReadDraftTableRows", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableIdOrName string

		resp, httpRes, err := apiClient.RowsApi.ReadDraftTableRows(context.Background(), tableIdOrName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RowsApiService ReplaceDraftTableRow", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableIdOrName string
		var rowId string

		resp, httpRes, err := apiClient.RowsApi.ReplaceDraftTableRow(context.Background(), tableIdOrName, rowId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RowsApiService UpdateDraftTableRow", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableIdOrName string
		var rowId string

		resp, httpRes, err := apiClient.RowsApi.UpdateDraftTableRow(context.Background(), tableIdOrName, rowId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
