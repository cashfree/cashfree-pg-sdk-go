/*
Cashfree Payment Gateway APIs

Testing SettlementsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cashfree_pg

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	cashfree "github.com/cashfree/cashfree_pg"
)

func Test_cashfree_pg_SettlementsAPIService(t *testing.T) {

	configuration := cashfree.NewConfiguration()
	apiClient := cashfree.NewAPIClient(configuration)

	t.Run("Test SettlementsAPIService PGOrderFetchSettlement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId string

		resp, httpRes, err := apiClient.SettlementsAPI.PGOrderFetchSettlement(context.Background(), orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}