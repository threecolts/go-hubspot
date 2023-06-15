/*
Business Unit

Testing BusinessUnitApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package business_units

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_business_units_BusinessUnitApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BusinessUnitApiService GetBusinessUnitsV3BusinessUnitsUserUserId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.BusinessUnitApi.GetBusinessUnitsV3BusinessUnitsUserUserId(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}