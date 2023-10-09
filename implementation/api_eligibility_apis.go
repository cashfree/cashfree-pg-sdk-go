/*
New Payment Gateway APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2022-09-01
Contact: nextgenapi@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg_sdk_go

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// EligibilityAPIsApiService EligibilityAPIsApi service
type EligibilityAPIsApiService service

type ApiEligibilityCardlessEMIRequest struct {
	ctx context.Context
	ApiService *EligibilityAPIsApiService
	xClientId *string
	xClientSecret *string
	xApiVersion *string
	eligibilityCardlessEMIRequest *EligibilityCardlessEMIRequest
}

func (r ApiEligibilityCardlessEMIRequest) XClientId(xClientId string) ApiEligibilityCardlessEMIRequest {
	r.xClientId = &xClientId
	return r
}

func (r ApiEligibilityCardlessEMIRequest) XClientSecret(xClientSecret string) ApiEligibilityCardlessEMIRequest {
	r.xClientSecret = &xClientSecret
	return r
}

func (r ApiEligibilityCardlessEMIRequest) XApiVersion(xApiVersion string) ApiEligibilityCardlessEMIRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiEligibilityCardlessEMIRequest) EligibilityCardlessEMIRequest(eligibilityCardlessEMIRequest EligibilityCardlessEMIRequest) ApiEligibilityCardlessEMIRequest {
	r.eligibilityCardlessEMIRequest = &eligibilityCardlessEMIRequest
	return r
}

func (r ApiEligibilityCardlessEMIRequest) Execute() ([]EligibleCardlessEMIEntity, *http.Response, error) {
	return r.ApiService.EligibilityCardlessEMIExecute(r)
}

/*
EligibilityCardlessEMI Get eligible Cardless EMI

Use this API to get eligible Cardless EMI Payment Methods for a customer on an order.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEligibilityCardlessEMIRequest
*/
func (a *EligibilityAPIsApiService) EligibilityCardlessEMI(ctx context.Context) ApiEligibilityCardlessEMIRequest {
	return ApiEligibilityCardlessEMIRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []EligibleCardlessEMIEntity
func (a *EligibilityAPIsApiService) EligibilityCardlessEMIExecute(r ApiEligibilityCardlessEMIRequest) ([]EligibleCardlessEMIEntity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EligibleCardlessEMIEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EligibilityAPIsApiService.EligibilityCardlessEMI")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eligibility/cardlessemi"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xClientId == nil {
		return localVarReturnValue, nil, reportError("xClientId is required and must be specified")
	}
	if r.xClientSecret == nil {
		return localVarReturnValue, nil, reportError("xClientSecret is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-client-id", r.xClientId, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-client-secret", r.xClientSecret, "")
	if r.xApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", r.xApiVersion, "")
	}
	// body params
	localVarPostBody = r.eligibilityCardlessEMIRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEligibilityOfferRequest struct {
	ctx context.Context
	ApiService *EligibilityAPIsApiService
	xClientId *string
	xClientSecret *string
	xApiVersion *string
	xRequestId *string
	xIdempotencyKey *string
	eligibilityOffersRequest *CFEligibilityOffersRequest
}

func (r ApiEligibilityOfferRequest) XClientId(xClientId string) ApiEligibilityOfferRequest {
	r.xClientId = &xClientId
	return r
}

func (r ApiEligibilityOfferRequest) XClientSecret(xClientSecret string) ApiEligibilityOfferRequest {
	r.xClientSecret = &xClientSecret
	return r
}

func (r ApiEligibilityOfferRequest) XApiVersion(xApiVersion string) ApiEligibilityOfferRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiEligibilityOfferRequest) XRequestId(xRequestId string) ApiEligibilityOfferRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiEligibilityOfferRequest) XIdempotencyKey(xIdempotencyKey string) ApiEligibilityOfferRequest {
	r.xIdempotencyKey = &xIdempotencyKey
	return r
}

func (r ApiEligibilityOfferRequest) EligibilityOffersRequest(eligibilityOffersRequest CFEligibilityOffersRequest) ApiEligibilityOfferRequest {
	r.eligibilityOffersRequest = &eligibilityOffersRequest
	return r
}

func (r ApiEligibilityOfferRequest) Execute() ([]EligibleOffersEntity, *http.Response, error) {
	return r.ApiService.EligibilityOfferExecute(r)
}

/*
EligibilityOffer Get eligible Offers

Use this API to get eligible offers for an order or amount.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEligibilityOfferRequest
*/
func (a *EligibilityAPIsApiService) EligibilityOffer(ctx context.Context) ApiEligibilityOfferRequest {
	return ApiEligibilityOfferRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []EligibleOffersEntity
func (a *EligibilityAPIsApiService) EligibilityOfferExecute(r ApiEligibilityOfferRequest) ([]EligibleOffersEntity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EligibleOffersEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EligibilityAPIsApiService.EligibilityOffer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eligibility/offers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xClientId == nil {
		return localVarReturnValue, nil, reportError("xClientId is required and must be specified")
	}
	if r.xClientSecret == nil {
		return localVarReturnValue, nil, reportError("xClientSecret is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-client-id", r.xClientId, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-client-secret", r.xClientSecret, "")
	if r.xApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", r.xApiVersion, "")
	}
	// body params
	localVarPostBody = r.eligibilityOffersRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
