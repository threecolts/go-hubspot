/*
Feedback Submissions

Testing BasicApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package feedback_submissions

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_feedback_submissions_BasicApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BasicApiService Archive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var feedbackSubmissionId string

		httpRes, err := apiClient.BasicApi.Archive(context.Background(), feedbackSubmissionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicApiService Create", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BasicApi.Create(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicApiService GetByID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var feedbackSubmissionId string

		resp, httpRes, err := apiClient.BasicApi.GetByID(context.Background(), feedbackSubmissionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicApiService GetPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BasicApi.GetPage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicApiService Update", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var feedbackSubmissionId string

		resp, httpRes, err := apiClient.BasicApi.Update(context.Background(), feedbackSubmissionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}