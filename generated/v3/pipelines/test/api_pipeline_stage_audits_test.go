/*
CRM Pipelines

Testing PipelineStageAuditsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package pipelines

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_pipelines_PipelineStageAuditsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PipelineStageAuditsApiService StagesGetAudit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectType string
		var stageId string

		resp, httpRes, err := apiClient.PipelineStageAuditsApi.StagesGetAudit(context.Background(), objectType, stageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
