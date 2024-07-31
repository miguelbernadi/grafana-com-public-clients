/*
GCOM API

Testing StackRegionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package gcom

import (
	"context"
	"testing"

	openapiclient "github.com/grafana/grafana-com-public-clients/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_gcom_StackRegionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StackRegionsAPIService GetStackRegions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StackRegionsAPI.GetStackRegions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
