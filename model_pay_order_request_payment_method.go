/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2022-09-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"fmt"
)

// PayOrderRequestPaymentMethod - struct for PayOrderRequestPaymentMethod
type PayOrderRequestPaymentMethod struct {
	AppPaymentMethod *AppPaymentMethod
	CardEMIPaymentMethod *CardEMIPaymentMethod
	CardPaymentMethod *CardPaymentMethod
	CardlessEMIPaymentMethod *CardlessEMIPaymentMethod
	NetBankingPaymentMethod *NetBankingPaymentMethod
	PaylaterPaymentMethod *PaylaterPaymentMethod
	UPIPaymentMethod *UPIPaymentMethod
}

// AppPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns AppPaymentMethod wrapped in PayOrderRequestPaymentMethod
func AppPaymentMethodAsPayOrderRequestPaymentMethod(v *AppPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		AppPaymentMethod: v,
	}
}

// CardEMIPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns CardEMIPaymentMethod wrapped in PayOrderRequestPaymentMethod
func CardEMIPaymentMethodAsPayOrderRequestPaymentMethod(v *CardEMIPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		CardEMIPaymentMethod: v,
	}
}

// CardPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns CardPaymentMethod wrapped in PayOrderRequestPaymentMethod
func CardPaymentMethodAsPayOrderRequestPaymentMethod(v *CardPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		CardPaymentMethod: v,
	}
}

// CardlessEMIPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns CardlessEMIPaymentMethod wrapped in PayOrderRequestPaymentMethod
func CardlessEMIPaymentMethodAsPayOrderRequestPaymentMethod(v *CardlessEMIPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		CardlessEMIPaymentMethod: v,
	}
}

// NetBankingPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns NetBankingPaymentMethod wrapped in PayOrderRequestPaymentMethod
func NetBankingPaymentMethodAsPayOrderRequestPaymentMethod(v *NetBankingPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		NetBankingPaymentMethod: v,
	}
}

// PaylaterPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns PaylaterPaymentMethod wrapped in PayOrderRequestPaymentMethod
func PaylaterPaymentMethodAsPayOrderRequestPaymentMethod(v *PaylaterPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		PaylaterPaymentMethod: v,
	}
}

// UPIPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns UPIPaymentMethod wrapped in PayOrderRequestPaymentMethod
func UPIPaymentMethodAsPayOrderRequestPaymentMethod(v *UPIPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		UPIPaymentMethod: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PayOrderRequestPaymentMethod) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'AppPaymentMethod'
	if jsonDict["type"] == "AppPaymentMethod" {
		// try to unmarshal JSON data into AppPaymentMethod
		err = json.Unmarshal(data, &dst.AppPaymentMethod)
		if err == nil {
			return nil // data stored in dst.AppPaymentMethod, return on the first match
		} else {
			dst.AppPaymentMethod = nil
			return fmt.Errorf("failed to unmarshal PayOrderRequestPaymentMethod as AppPaymentMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CardEMIPaymentMethod'
	if jsonDict["type"] == "CardEMIPaymentMethod" {
		// try to unmarshal JSON data into CardEMIPaymentMethod
		err = json.Unmarshal(data, &dst.CardEMIPaymentMethod)
		if err == nil {
			return nil // data stored in dst.CardEMIPaymentMethod, return on the first match
		} else {
			dst.CardEMIPaymentMethod = nil
			return fmt.Errorf("failed to unmarshal PayOrderRequestPaymentMethod as CardEMIPaymentMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CardPaymentMethod'
	if jsonDict["type"] == "CardPaymentMethod" {
		// try to unmarshal JSON data into CardPaymentMethod
		err = json.Unmarshal(data, &dst.CardPaymentMethod)
		if err == nil {
			return nil // data stored in dst.CardPaymentMethod, return on the first match
		} else {
			dst.CardPaymentMethod = nil
			return fmt.Errorf("failed to unmarshal PayOrderRequestPaymentMethod as CardPaymentMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CardlessEMIPaymentMethod'
	if jsonDict["type"] == "CardlessEMIPaymentMethod" {
		// try to unmarshal JSON data into CardlessEMIPaymentMethod
		err = json.Unmarshal(data, &dst.CardlessEMIPaymentMethod)
		if err == nil {
			return nil // data stored in dst.CardlessEMIPaymentMethod, return on the first match
		} else {
			dst.CardlessEMIPaymentMethod = nil
			return fmt.Errorf("failed to unmarshal PayOrderRequestPaymentMethod as CardlessEMIPaymentMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'NetBankingPaymentMethod'
	if jsonDict["type"] == "NetBankingPaymentMethod" {
		// try to unmarshal JSON data into NetBankingPaymentMethod
		err = json.Unmarshal(data, &dst.NetBankingPaymentMethod)
		if err == nil {
			return nil // data stored in dst.NetBankingPaymentMethod, return on the first match
		} else {
			dst.NetBankingPaymentMethod = nil
			return fmt.Errorf("failed to unmarshal PayOrderRequestPaymentMethod as NetBankingPaymentMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PaylaterPaymentMethod'
	if jsonDict["type"] == "PaylaterPaymentMethod" {
		// try to unmarshal JSON data into PaylaterPaymentMethod
		err = json.Unmarshal(data, &dst.PaylaterPaymentMethod)
		if err == nil {
			return nil // data stored in dst.PaylaterPaymentMethod, return on the first match
		} else {
			dst.PaylaterPaymentMethod = nil
			return fmt.Errorf("failed to unmarshal PayOrderRequestPaymentMethod as PaylaterPaymentMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UPIPaymentMethod'
	if jsonDict["type"] == "UPIPaymentMethod" {
		// try to unmarshal JSON data into UPIPaymentMethod
		err = json.Unmarshal(data, &dst.UPIPaymentMethod)
		if err == nil {
			return nil // data stored in dst.UPIPaymentMethod, return on the first match
		} else {
			dst.UPIPaymentMethod = nil
			return fmt.Errorf("failed to unmarshal PayOrderRequestPaymentMethod as UPIPaymentMethod: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PayOrderRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	if src.AppPaymentMethod != nil {
		return json.Marshal(&src.AppPaymentMethod)
	}

	if src.CardEMIPaymentMethod != nil {
		return json.Marshal(&src.CardEMIPaymentMethod)
	}

	if src.CardPaymentMethod != nil {
		return json.Marshal(&src.CardPaymentMethod)
	}

	if src.CardlessEMIPaymentMethod != nil {
		return json.Marshal(&src.CardlessEMIPaymentMethod)
	}

	if src.NetBankingPaymentMethod != nil {
		return json.Marshal(&src.NetBankingPaymentMethod)
	}

	if src.PaylaterPaymentMethod != nil {
		return json.Marshal(&src.PaylaterPaymentMethod)
	}

	if src.UPIPaymentMethod != nil {
		return json.Marshal(&src.UPIPaymentMethod)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PayOrderRequestPaymentMethod) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppPaymentMethod != nil {
		return obj.AppPaymentMethod
	}

	if obj.CardEMIPaymentMethod != nil {
		return obj.CardEMIPaymentMethod
	}

	if obj.CardPaymentMethod != nil {
		return obj.CardPaymentMethod
	}

	if obj.CardlessEMIPaymentMethod != nil {
		return obj.CardlessEMIPaymentMethod
	}

	if obj.NetBankingPaymentMethod != nil {
		return obj.NetBankingPaymentMethod
	}

	if obj.PaylaterPaymentMethod != nil {
		return obj.PaylaterPaymentMethod
	}

	if obj.UPIPaymentMethod != nil {
		return obj.UPIPaymentMethod
	}

	// all schemas are nil
	return nil
}

type NullablePayOrderRequestPaymentMethod struct {
	value *PayOrderRequestPaymentMethod
	isSet bool
}

func (v NullablePayOrderRequestPaymentMethod) Get() *PayOrderRequestPaymentMethod {
	return v.value
}

func (v *NullablePayOrderRequestPaymentMethod) Set(val *PayOrderRequestPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePayOrderRequestPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePayOrderRequestPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayOrderRequestPaymentMethod(val *PayOrderRequestPaymentMethod) *NullablePayOrderRequestPaymentMethod {
	return &NullablePayOrderRequestPaymentMethod{value: val, isSet: true}
}

func (v NullablePayOrderRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayOrderRequestPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


