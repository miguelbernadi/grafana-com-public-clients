/*
GCOM API

Testing OrgsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package gcom

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/grafana/grafana-com-public-clients/go/gcom"
)

func Test_gcom_OrgsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsAPIService DelApiKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var slugOrId string

		httpRes, err := apiClient.OrgsAPI.DelApiKey(context.Background(), name, slugOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAPIService GetApiKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var slugOrId string

		resp, httpRes, err := apiClient.OrgsAPI.GetApiKey(context.Background(), name, slugOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAPIService GetApiKeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var slugOrId string

		resp, httpRes, err := apiClient.OrgsAPI.GetApiKeys(context.Background(), slugOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAPIService GetOrgInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var slug string

		resp, httpRes, err := apiClient.OrgsAPI.GetOrgInstances(context.Background(), slug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAPIService PostApiKeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var slugOrId string

		resp, httpRes, err := apiClient.OrgsAPI.PostApiKeys(context.Background(), slugOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
