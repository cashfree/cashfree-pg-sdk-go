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
	"fmt"
)

// PaymentEntityPaymentMethod - struct for PaymentEntityPaymentMethod
type PaymentEntityPaymentMethod struct {
	PaymentMethodAppInPaymentsEntity *PaymentMethodAppInPaymentsEntity
	PaymentMethodCardEMIInPaymentsEntity *PaymentMethodCardEMIInPaymentsEntity
	PaymentMethodCardInPaymentsEntity *PaymentMethodCardInPaymentsEntity
	PaymentMethodCardlessEMIInPaymentsEntity *PaymentMethodCardlessEMIInPaymentsEntity
	PaymentMethodNetBankingInPaymentsEntity *PaymentMethodNetBankingInPaymentsEntity
	PaymentMethodPaylaterInPaymentsEntity *PaymentMethodPaylaterInPaymentsEntity
	PaymentMethodUPIInPaymentsEntity *PaymentMethodUPIInPaymentsEntity
}

// PaymentMethodAppInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodAppInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodAppInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodAppInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodAppInPaymentsEntity: v,
	}
}

// PaymentMethodCardEMIInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodCardEMIInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodCardEMIInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodCardEMIInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodCardEMIInPaymentsEntity: v,
	}
}

// PaymentMethodCardInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodCardInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodCardInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodCardInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodCardInPaymentsEntity: v,
	}
}

// PaymentMethodCardlessEMIInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodCardlessEMIInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodCardlessEMIInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodCardlessEMIInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodCardlessEMIInPaymentsEntity: v,
	}
}

// PaymentMethodNetBankingInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodNetBankingInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodNetBankingInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodNetBankingInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodNetBankingInPaymentsEntity: v,
	}
}

// PaymentMethodPaylaterInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodPaylaterInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodPaylaterInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodPaylaterInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodPaylaterInPaymentsEntity: v,
	}
}

// PaymentMethodUPIInPaymentsEntityAsPaymentEntityPaymentMethod is a convenience function that returns PaymentMethodUPIInPaymentsEntity wrapped in PaymentEntityPaymentMethod
func PaymentMethodUPIInPaymentsEntityAsPaymentEntityPaymentMethod(v *PaymentMethodUPIInPaymentsEntity) PaymentEntityPaymentMethod {
	return PaymentEntityPaymentMethod{
		PaymentMethodUPIInPaymentsEntity: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PaymentEntityPaymentMethod) UnmarshalJSON(data []byte) error {
		var err error





	match := 0


	// try to unmarshal data into PaymentMethodAppInPaymentsEntity

	err = json.Unmarshal(data, &dst.PaymentMethodAppInPaymentsEntity)

	if err == nil {

		jsonPaymentMethodAppInPaymentsEntity, _ := json.Marshal(dst.PaymentMethodAppInPaymentsEntity)

		if strings.Contains(string(jsonPaymentMethodAppInPaymentsEntity), "{}") || strings.Contains(string(jsonPaymentMethodAppInPaymentsEntity), "null") { // empty struct

			dst.PaymentMethodAppInPaymentsEntity = nil

		} else {

			match++

		}

	} else {

		dst.PaymentMethodAppInPaymentsEntity = nil

	}


	// try to unmarshal data into PaymentMethodCardEMIInPaymentsEntity

	err = json.Unmarshal(data, &dst.PaymentMethodCardEMIInPaymentsEntity)

	if err == nil {

		jsonPaymentMethodCardEMIInPaymentsEntity, _ := json.Marshal(dst.PaymentMethodCardEMIInPaymentsEntity)

		if strings.Contains(string(jsonPaymentMethodCardEMIInPaymentsEntity), "{}") || strings.Contains(string(jsonPaymentMethodCardEMIInPaymentsEntity), "null") { // empty struct

			dst.PaymentMethodCardEMIInPaymentsEntity = nil

		} else {

			match++

		}

	} else {

		dst.PaymentMethodCardEMIInPaymentsEntity = nil

	}


	// try to unmarshal data into PaymentMethodCardInPaymentsEntity

	err = json.Unmarshal(data, &dst.PaymentMethodCardInPaymentsEntity)

	if err == nil {

		jsonPaymentMethodCardInPaymentsEntity, _ := json.Marshal(dst.PaymentMethodCardInPaymentsEntity)

		if strings.Contains(string(jsonPaymentMethodCardInPaymentsEntity), "{}") || strings.Contains(string(jsonPaymentMethodCardInPaymentsEntity), "null") { // empty struct

			dst.PaymentMethodCardInPaymentsEntity = nil

		} else {

			match++

		}

	} else {

		dst.PaymentMethodCardInPaymentsEntity = nil

	}


	// try to unmarshal data into PaymentMethodCardlessEMIInPaymentsEntity

	err = json.Unmarshal(data, &dst.PaymentMethodCardlessEMIInPaymentsEntity)

	if err == nil {

		jsonPaymentMethodCardlessEMIInPaymentsEntity, _ := json.Marshal(dst.PaymentMethodCardlessEMIInPaymentsEntity)

		if strings.Contains(string(jsonPaymentMethodCardlessEMIInPaymentsEntity), "{}") || strings.Contains(string(jsonPaymentMethodCardlessEMIInPaymentsEntity), "null") { // empty struct

			dst.PaymentMethodCardlessEMIInPaymentsEntity = nil

		} else {

			match++

		}

	} else {

		dst.PaymentMethodCardlessEMIInPaymentsEntity = nil

	}


	// try to unmarshal data into PaymentMethodNetBankingInPaymentsEntity

	err = json.Unmarshal(data, &dst.PaymentMethodNetBankingInPaymentsEntity)

	if err == nil {

		jsonPaymentMethodNetBankingInPaymentsEntity, _ := json.Marshal(dst.PaymentMethodNetBankingInPaymentsEntity)

		if strings.Contains(string(jsonPaymentMethodNetBankingInPaymentsEntity), "{}") || strings.Contains(string(jsonPaymentMethodNetBankingInPaymentsEntity), "null") { // empty struct

			dst.PaymentMethodNetBankingInPaymentsEntity = nil

		} else {

			match++

		}

	} else {

		dst.PaymentMethodNetBankingInPaymentsEntity = nil

	}


	// try to unmarshal data into PaymentMethodPaylaterInPaymentsEntity

	err = json.Unmarshal(data, &dst.PaymentMethodPaylaterInPaymentsEntity)

	if err == nil {

		jsonPaymentMethodPaylaterInPaymentsEntity, _ := json.Marshal(dst.PaymentMethodPaylaterInPaymentsEntity)

		if strings.Contains(string(jsonPaymentMethodPaylaterInPaymentsEntity), "{}") || strings.Contains(string(jsonPaymentMethodPaylaterInPaymentsEntity), "null") { // empty struct

			dst.PaymentMethodPaylaterInPaymentsEntity = nil

		} else {

			match++

		}

	} else {

		dst.PaymentMethodPaylaterInPaymentsEntity = nil

	}


	// try to unmarshal data into PaymentMethodUPIInPaymentsEntity

	err = json.Unmarshal(data, &dst.PaymentMethodUPIInPaymentsEntity)

	if err == nil {

		jsonPaymentMethodUPIInPaymentsEntity, _ := json.Marshal(dst.PaymentMethodUPIInPaymentsEntity)

		if strings.Contains(string(jsonPaymentMethodUPIInPaymentsEntity), "{}") || strings.Contains(string(jsonPaymentMethodUPIInPaymentsEntity), "null") { // empty struct

			dst.PaymentMethodUPIInPaymentsEntity = nil

		} else {

			match++

		}

	} else {

		dst.PaymentMethodUPIInPaymentsEntity = nil

	}


	if match > 1 { // more than 1 match

		// reset to nil


		dst.PaymentMethodAppInPaymentsEntity = nil


		dst.PaymentMethodCardEMIInPaymentsEntity = nil


		dst.PaymentMethodCardInPaymentsEntity = nil


		dst.PaymentMethodCardlessEMIInPaymentsEntity = nil


		dst.PaymentMethodNetBankingInPaymentsEntity = nil


		dst.PaymentMethodPaylaterInPaymentsEntity = nil


		dst.PaymentMethodUPIInPaymentsEntity = nil


		return fmt.Errorf("data matches more than one schema in oneOf(PaymentEntityPaymentMethod)")

	} else if match == 1 {

		return nil // exactly one match

	} else { // no match

		return fmt.Errorf("data failed to match schemas in oneOf(PaymentEntityPaymentMethod)")

	}



}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PaymentEntityPaymentMethod) MarshalJSON() ([]byte, error) {
	if src.PaymentMethodAppInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodAppInPaymentsEntity)
	}

	if src.PaymentMethodCardEMIInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodCardEMIInPaymentsEntity)
	}

	if src.PaymentMethodCardInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodCardInPaymentsEntity)
	}

	if src.PaymentMethodCardlessEMIInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodCardlessEMIInPaymentsEntity)
	}

	if src.PaymentMethodNetBankingInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodNetBankingInPaymentsEntity)
	}

	if src.PaymentMethodPaylaterInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodPaylaterInPaymentsEntity)
	}

	if src.PaymentMethodUPIInPaymentsEntity != nil {
		return json.Marshal(&src.PaymentMethodUPIInPaymentsEntity)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PaymentEntityPaymentMethod) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PaymentMethodAppInPaymentsEntity != nil {
		return obj.PaymentMethodAppInPaymentsEntity
	}

	if obj.PaymentMethodCardEMIInPaymentsEntity != nil {
		return obj.PaymentMethodCardEMIInPaymentsEntity
	}

	if obj.PaymentMethodCardInPaymentsEntity != nil {
		return obj.PaymentMethodCardInPaymentsEntity
	}

	if obj.PaymentMethodCardlessEMIInPaymentsEntity != nil {
		return obj.PaymentMethodCardlessEMIInPaymentsEntity
	}

	if obj.PaymentMethodNetBankingInPaymentsEntity != nil {
		return obj.PaymentMethodNetBankingInPaymentsEntity
	}

	if obj.PaymentMethodPaylaterInPaymentsEntity != nil {
		return obj.PaymentMethodPaylaterInPaymentsEntity
	}

	if obj.PaymentMethodUPIInPaymentsEntity != nil {
		return obj.PaymentMethodUPIInPaymentsEntity
	}

	// all schemas are nil
	return nil
}




func cashfreeStringTest() {
	strings.HasPrefix("cf", "cf")
}