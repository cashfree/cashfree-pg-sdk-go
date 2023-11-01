/*
Cashfree Payment Gateway APIs

Testing PGReconciliationAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cashfree_pg

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	cashfree "github.com/cashfree/cashfree-pg-sdk-go"
)

func Test_cashfree_pg_PGReconciliationAPIService(t *testing.T) {

	configuration := cashfree.NewConfiguration()
	apiClient := cashfree.NewAPIClient(configuration)

	t.Run("Test PGReconciliationAPIService PGFetchRecon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PGReconciliationAPI.PGFetchRecon(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
