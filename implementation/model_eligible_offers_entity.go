/*
New Payment Gateway APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2022-09-01
Contact: nextgenapi@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg_sdk_go

import (
	"encoding/json"
)

// checks if the EligibleOffersEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EligibleOffersEntity{}

// EligibleOffersEntity struct for EligibleOffersEntity
type EligibleOffersEntity struct {
	Eligibility *string `json:"eligibility,omitempty"`
	EntityType *string `json:"entity_type,omitempty"`
	EntityValue *string `json:"entity_value,omitempty"`
	EntityDetails *CFOfferEntity `json:"entity_details,omitempty"`
}

// NewEligibleOffersEntity instantiates a new EligibleOffersEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEligibleOffersEntity() *EligibleOffersEntity {
	this := EligibleOffersEntity{}
	return &this
}

// NewEligibleOffersEntityWithDefaults instantiates a new EligibleOffersEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEligibleOffersEntityWithDefaults() *EligibleOffersEntity {
	this := EligibleOffersEntity{}
	return &this
}

// GetEligibility returns the Eligibility field value if set, zero value otherwise.
func (o *EligibleOffersEntity) GetEligibility() string {
	if o == nil || IsNil(o.Eligibility) {
		var ret string
		return ret
	}
	return *o.Eligibility
}

// GetEligibilityOk returns a tuple with the Eligibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibleOffersEntity) GetEligibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Eligibility) {
		return nil, false
	}
	return o.Eligibility, true
}

// HasEligibility returns a boolean if a field has been set.
func (o *EligibleOffersEntity) HasEligibility() bool {
	if o != nil && !IsNil(o.Eligibility) {
		return true
	}

	return false
}

// SetEligibility gets a reference to the given string and assigns it to the Eligibility field.
func (o *EligibleOffersEntity) SetEligibility(v string) {
	o.Eligibility = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *EligibleOffersEntity) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibleOffersEntity) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *EligibleOffersEntity) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *EligibleOffersEntity) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityValue returns the EntityValue field value if set, zero value otherwise.
func (o *EligibleOffersEntity) GetEntityValue() string {
	if o == nil || IsNil(o.EntityValue) {
		var ret string
		return ret
	}
	return *o.EntityValue
}

// GetEntityValueOk returns a tuple with the EntityValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibleOffersEntity) GetEntityValueOk() (*string, bool) {
	if o == nil || IsNil(o.EntityValue) {
		return nil, false
	}
	return o.EntityValue, true
}

// HasEntityValue returns a boolean if a field has been set.
func (o *EligibleOffersEntity) HasEntityValue() bool {
	if o != nil && !IsNil(o.EntityValue) {
		return true
	}

	return false
}

// SetEntityValue gets a reference to the given string and assigns it to the EntityValue field.
func (o *EligibleOffersEntity) SetEntityValue(v string) {
	o.EntityValue = &v
}

// GetEntityDetails returns the EntityDetails field value if set, zero value otherwise.
func (o *EligibleOffersEntity) GetEntityDetails() CFOfferEntity {
	if o == nil || IsNil(o.EntityDetails) {
		var ret CFOfferEntity
		return ret
	}
	return *o.EntityDetails
}

// GetEntityDetailsOk returns a tuple with the EntityDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibleOffersEntity) GetEntityDetailsOk() (*CFOfferEntity, bool) {
	if o == nil || IsNil(o.EntityDetails) {
		return nil, false
	}
	return o.EntityDetails, true
}

// HasEntityDetails returns a boolean if a field has been set.
func (o *EligibleOffersEntity) HasEntityDetails() bool {
	if o != nil && !IsNil(o.EntityDetails) {
		return true
	}

	return false
}

// SetEntityDetails gets a reference to the given OfferEntity and assigns it to the EntityDetails field.
func (o *EligibleOffersEntity) SetEntityDetails(v CFOfferEntity) {
	o.EntityDetails = &v
}

func (o EligibleOffersEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EligibleOffersEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Eligibility) {
		toSerialize["eligibility"] = o.Eligibility
	}
	if !IsNil(o.EntityType) {
		toSerialize["entity_type"] = o.EntityType
	}
	if !IsNil(o.EntityValue) {
		toSerialize["entity_value"] = o.EntityValue
	}
	if !IsNil(o.EntityDetails) {
		toSerialize["entity_details"] = o.EntityDetails
	}
	return toSerialize, nil
}

type NullableEligibleOffersEntity struct {
	value *EligibleOffersEntity
	isSet bool
}

func (v NullableEligibleOffersEntity) Get() *EligibleOffersEntity {
	return v.value
}

func (v *NullableEligibleOffersEntity) Set(val *EligibleOffersEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableEligibleOffersEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableEligibleOffersEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEligibleOffersEntity(val *EligibleOffersEntity) *NullableEligibleOffersEntity {
	return &NullableEligibleOffersEntity{value: val, isSet: true}
}

func (v NullableEligibleOffersEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEligibleOffersEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


