/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2023-08-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"strings"
)

// checks if the UploadPnachImageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadPnachImageResponse{}

// UploadPnachImageResponse Response of pnach image upload API.
type UploadPnachImageResponse struct {
	// The payment_id against which the pnach image is uploaded.
	PaymentId *string `json:"payment_id,omitempty"`
	// Authorization status of the subscription.
	AuthorizationStatus *string `json:"authorization_status,omitempty"`
	// Action performed on the file.
	Action *string `json:"action,omitempty"`
	// Message of the API.
	PaymentMessage *string `json:"payment_message,omitempty"`
}


func (o UploadPnachImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadPnachImageResponse) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentId) {
		toSerialize["payment_id"] = o.PaymentId
	}
	if !IsNil(o.AuthorizationStatus) {
		toSerialize["authorization_status"] = o.AuthorizationStatus
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.PaymentMessage) {
		toSerialize["payment_message"] = o.PaymentMessage
	}
	return toSerialize, nil
}


