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

// checks if the AddressDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressDetails{}

// AddressDetails Address associated with the customer.
type AddressDetails struct {
	// Full Name of the customer associated with the address.
	Name *string `json:"name,omitempty"`
	// First line of the address.
	AddressLineOne *string `json:"address_line_one,omitempty"`
	// Second line of the address.
	AddressLineTwo *string `json:"address_line_two,omitempty"`
	// Country Name.
	Country *string `json:"country,omitempty"`
	// Country Code.
	CountryCode *string `json:"country_code,omitempty"`
	// State Name.
	State *string `json:"state,omitempty"`
	// State Code.
	StateCode *string `json:"state_code,omitempty"`
	// City Name.
	City *string `json:"city,omitempty"`
	// Pin Code/Zip Code.
	PinCode *string `json:"pin_code,omitempty"`
	// Customer Phone Number.
	Phone *string `json:"phone,omitempty"`
	// Cutomer Email Address.
	Email *string `json:"email,omitempty"`
}


func (o AddressDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AddressLineOne) {
		toSerialize["address_line_one"] = o.AddressLineOne
	}
	if !IsNil(o.AddressLineTwo) {
		toSerialize["address_line_two"] = o.AddressLineTwo
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StateCode) {
		toSerialize["state_code"] = o.StateCode
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.PinCode) {
		toSerialize["pin_code"] = o.PinCode
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}



