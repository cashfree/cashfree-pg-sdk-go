/*
Cashfree Payment Gateway APIs

Testing PaymentsAPIService

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

func Test_cashfree_pg_PaymentsAPIService(t *testing.T) {

	configuration := cashfree.NewConfiguration()
	apiClient := cashfree.NewAPIClient(configuration)

	t.Run("Test PaymentsAPIService PGAuthorizeOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId string

		resp, httpRes, err := apiClient.PaymentsAPI.PGAuthorizeOrder(context.Background(), orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentsAPIService PGOrderAuthenticatePayment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cfPaymentId string

		resp, httpRes, err := apiClient.PaymentsAPI.PGOrderAuthenticatePayment(context.Background(), cfPaymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentsAPIService PGOrderFetchPayment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId string
		var cfPaymentId string

		resp, httpRes, err := apiClient.PaymentsAPI.PGOrderFetchPayment(context.Background(), orderId, cfPaymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentsAPIService PGOrderFetchPayments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId string

		resp, httpRes, err := apiClient.PaymentsAPI.PGOrderFetchPayments(context.Background(), orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentsAPIService PGPayOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PaymentsAPI.PGPayOrder(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}