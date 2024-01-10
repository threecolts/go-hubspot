/*
Lists

Testing MembershipsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package lists

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_lists_MembershipsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MembershipsAPIService Add", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId int32

		resp, httpRes, err := apiClient.MembershipsAPI.Add(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MembershipsAPIService AddAllFromList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId int32
		var sourceListId int32

		httpRes, err := apiClient.MembershipsAPI.AddAllFromList(context.Background(), listId, sourceListId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MembershipsAPIService AddAndRemove", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId int32

		resp, httpRes, err := apiClient.MembershipsAPI.AddAndRemove(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MembershipsAPIService GetPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId int32

		resp, httpRes, err := apiClient.MembershipsAPI.GetPage(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MembershipsAPIService Remove", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId int32

		resp, httpRes, err := apiClient.MembershipsAPI.Remove(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MembershipsAPIService RemoveAll", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId int32

		httpRes, err := apiClient.MembershipsAPI.RemoveAll(context.Background(), listId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
