/*
Blog Post endpoints

Testing BlogPostsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package blog_posts

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_blog_posts_BlogPostsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BlogPostsApiService Archive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		httpRes, err := apiClient.BlogPostsApi.Archive(context.Background(), objectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService AttachToLanguageGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsApi.AttachToLanguageGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService BatchArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BlogPostsApi.BatchArchive(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService BatchCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsApi.BatchCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService BatchRead", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsApi.BatchRead(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService BatchUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsApi.BatchUpdate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService Clone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsApi.Clone(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService Create", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsApi.Create(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService CreateLanguageVariation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsApi.CreateLanguageVariation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService DetachFromLanguageGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsApi.DetachFromLanguageGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService GetByID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.BlogPostsApi.GetByID(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService GetDraftByID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.BlogPostsApi.GetDraftByID(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService GetPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsApi.GetPage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService GetPreviousVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string
		var revisionId string

		resp, httpRes, err := apiClient.BlogPostsApi.GetPreviousVersion(context.Background(), objectId, revisionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService GetPreviousVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.BlogPostsApi.GetPreviousVersions(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService PushLive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		httpRes, err := apiClient.BlogPostsApi.PushLive(context.Background(), objectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService ResetDraft", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		httpRes, err := apiClient.BlogPostsApi.ResetDraft(context.Background(), objectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService RestorePreviousVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string
		var revisionId string

		resp, httpRes, err := apiClient.BlogPostsApi.RestorePreviousVersion(context.Background(), objectId, revisionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService RestorePreviousVersionToDraft", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string
		var revisionId int64

		resp, httpRes, err := apiClient.BlogPostsApi.RestorePreviousVersionToDraft(context.Background(), objectId, revisionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService Schedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BlogPostsApi.Schedule(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService SetLanguagePrimary", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BlogPostsApi.SetLanguagePrimary(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService Update", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.BlogPostsApi.Update(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService UpdateDraft", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.BlogPostsApi.UpdateDraft(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsApiService UpdateLanguages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsApi.UpdateLanguages(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
