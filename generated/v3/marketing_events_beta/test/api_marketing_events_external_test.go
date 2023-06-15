/*
Marketing Events Extension

Testing MarketingEventsExternalApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package marketing_events_beta

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_marketing_events_beta_MarketingEventsExternalApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MarketingEventsExternalApiService ArchiveBatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MarketingEventsExternalApi.ArchiveBatch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketingEventsExternalApiService Create", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MarketingEventsExternalApi.Create(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketingEventsExternalApiService ExternalArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalEventId string

		httpRes, err := apiClient.MarketingEventsExternalApi.ExternalArchive(context.Background(), externalEventId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketingEventsExternalApiService ExternalCancel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalEventId string

		resp, httpRes, err := apiClient.MarketingEventsExternalApi.ExternalCancel(context.Background(), externalEventId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketingEventsExternalApiService ExternalCompleteComplete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalEventId string

		resp, httpRes, err := apiClient.MarketingEventsExternalApi.ExternalCompleteComplete(context.Background(), externalEventId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketingEventsExternalApiService ExternalEmailUpsertByID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalEventId string
		var subscriberState string

		resp, httpRes, err := apiClient.MarketingEventsExternalApi.ExternalEmailUpsertByID(context.Background(), externalEventId, subscriberState).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketingEventsExternalApiService ExternalGetByID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalEventId string

		resp, httpRes, err := apiClient.MarketingEventsExternalApi.ExternalGetByID(context.Background(), externalEventId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketingEventsExternalApiService ExternalReplace", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalEventId string

		resp, httpRes, err := apiClient.MarketingEventsExternalApi.ExternalReplace(context.Background(), externalEventId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketingEventsExternalApiService ExternalUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalEventId string

		resp, httpRes, err := apiClient.MarketingEventsExternalApi.ExternalUpdate(context.Background(), externalEventId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketingEventsExternalApiService ExternalUpsertByID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalEventId string
		var subscriberState string

		resp, httpRes, err := apiClient.MarketingEventsExternalApi.ExternalUpsertByID(context.Background(), externalEventId, subscriberState).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketingEventsExternalApiService Upsert", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MarketingEventsExternalApi.Upsert(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
