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
)

// checks if the TerminalDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminalDetails{}

// TerminalDetails Use this if you are creating an order for cashfree's softPOS
type TerminalDetails struct {
	// date time at which terminal is added
	AddedOn *string `json:"added_on,omitempty"`
	// cashfree terminal id
	CfTerminalId *int32 `json:"cf_terminal_id,omitempty"`
	// last instant when this terminal was updated
	LastUpdatedOn *string `json:"last_updated_on,omitempty"`
	// location of terminal
	TerminalAddress *string `json:"terminal_address,omitempty"`
	// terminal id for merchant reference
	TerminalId string `json:"terminal_id"`
	// name of terminal/agent/storefront
	TerminalName *string `json:"terminal_name,omitempty"`
	// note given by merchant while creating the terminal
	TerminalNote *string `json:"terminal_note,omitempty"`
	// mobile num of the terminal/agent/storefront
	TerminalPhoneNo string `json:"terminal_phone_no"`
	// status of terminal active/inactive
	TerminalStatus *string `json:"terminal_status,omitempty"`
	// To identify the type of terminal product in use, in this case it is SPOS.
	TerminalType string `json:"terminal_type"`
}

// NewTerminalDetails instantiates a new TerminalDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalDetails(terminalId string, terminalPhoneNo string, terminalType string) *TerminalDetails {
	this := TerminalDetails{}
	this.TerminalId = terminalId
	this.TerminalPhoneNo = terminalPhoneNo
	this.TerminalType = terminalType
	return &this
}

// NewTerminalDetailsWithDefaults instantiates a new TerminalDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalDetailsWithDefaults() *TerminalDetails {
	this := TerminalDetails{}
	return &this
}

// GetAddedOn returns the AddedOn field value if set, zero value otherwise.
func (o *TerminalDetails) GetAddedOn() string {
	if o == nil || IsNil(o.AddedOn) {
		var ret string
		return ret
	}
	return *o.AddedOn
}

// GetAddedOnOk returns a tuple with the AddedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalDetails) GetAddedOnOk() (*string, bool) {
	if o == nil || IsNil(o.AddedOn) {
		return nil, false
	}
	return o.AddedOn, true
}

// HasAddedOn returns a boolean if a field has been set.
func (o *TerminalDetails) HasAddedOn() bool {
	if o != nil && !IsNil(o.AddedOn) {
		return true
	}

	return false
}

// SetAddedOn gets a reference to the given string and assigns it to the AddedOn field.
func (o *TerminalDetails) SetAddedOn(v string) {
	o.AddedOn = &v
}

// GetCfTerminalId returns the CfTerminalId field value if set, zero value otherwise.
func (o *TerminalDetails) GetCfTerminalId() int32 {
	if o == nil || IsNil(o.CfTerminalId) {
		var ret int32
		return ret
	}
	return *o.CfTerminalId
}

// GetCfTerminalIdOk returns a tuple with the CfTerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalDetails) GetCfTerminalIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CfTerminalId) {
		return nil, false
	}
	return o.CfTerminalId, true
}

// HasCfTerminalId returns a boolean if a field has been set.
func (o *TerminalDetails) HasCfTerminalId() bool {
	if o != nil && !IsNil(o.CfTerminalId) {
		return true
	}

	return false
}

// SetCfTerminalId gets a reference to the given int32 and assigns it to the CfTerminalId field.
func (o *TerminalDetails) SetCfTerminalId(v int32) {
	o.CfTerminalId = &v
}

// GetLastUpdatedOn returns the LastUpdatedOn field value if set, zero value otherwise.
func (o *TerminalDetails) GetLastUpdatedOn() string {
	if o == nil || IsNil(o.LastUpdatedOn) {
		var ret string
		return ret
	}
	return *o.LastUpdatedOn
}

// GetLastUpdatedOnOk returns a tuple with the LastUpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalDetails) GetLastUpdatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedOn) {
		return nil, false
	}
	return o.LastUpdatedOn, true
}

// HasLastUpdatedOn returns a boolean if a field has been set.
func (o *TerminalDetails) HasLastUpdatedOn() bool {
	if o != nil && !IsNil(o.LastUpdatedOn) {
		return true
	}

	return false
}

// SetLastUpdatedOn gets a reference to the given string and assigns it to the LastUpdatedOn field.
func (o *TerminalDetails) SetLastUpdatedOn(v string) {
	o.LastUpdatedOn = &v
}

// GetTerminalAddress returns the TerminalAddress field value if set, zero value otherwise.
func (o *TerminalDetails) GetTerminalAddress() string {
	if o == nil || IsNil(o.TerminalAddress) {
		var ret string
		return ret
	}
	return *o.TerminalAddress
}

// GetTerminalAddressOk returns a tuple with the TerminalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalDetails) GetTerminalAddressOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalAddress) {
		return nil, false
	}
	return o.TerminalAddress, true
}

// HasTerminalAddress returns a boolean if a field has been set.
func (o *TerminalDetails) HasTerminalAddress() bool {
	if o != nil && !IsNil(o.TerminalAddress) {
		return true
	}

	return false
}

// SetTerminalAddress gets a reference to the given string and assigns it to the TerminalAddress field.
func (o *TerminalDetails) SetTerminalAddress(v string) {
	o.TerminalAddress = &v
}

// GetTerminalId returns the TerminalId field value
func (o *TerminalDetails) GetTerminalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerminalId
}

// GetTerminalIdOk returns a tuple with the TerminalId field value
// and a boolean to check if the value has been set.
func (o *TerminalDetails) GetTerminalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminalId, true
}

// SetTerminalId sets field value
func (o *TerminalDetails) SetTerminalId(v string) {
	o.TerminalId = v
}

// GetTerminalName returns the TerminalName field value if set, zero value otherwise.
func (o *TerminalDetails) GetTerminalName() string {
	if o == nil || IsNil(o.TerminalName) {
		var ret string
		return ret
	}
	return *o.TerminalName
}

// GetTerminalNameOk returns a tuple with the TerminalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalDetails) GetTerminalNameOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalName) {
		return nil, false
	}
	return o.TerminalName, true
}

// HasTerminalName returns a boolean if a field has been set.
func (o *TerminalDetails) HasTerminalName() bool {
	if o != nil && !IsNil(o.TerminalName) {
		return true
	}

	return false
}

// SetTerminalName gets a reference to the given string and assigns it to the TerminalName field.
func (o *TerminalDetails) SetTerminalName(v string) {
	o.TerminalName = &v
}

// GetTerminalNote returns the TerminalNote field value if set, zero value otherwise.
func (o *TerminalDetails) GetTerminalNote() string {
	if o == nil || IsNil(o.TerminalNote) {
		var ret string
		return ret
	}
	return *o.TerminalNote
}

// GetTerminalNoteOk returns a tuple with the TerminalNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalDetails) GetTerminalNoteOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalNote) {
		return nil, false
	}
	return o.TerminalNote, true
}

// HasTerminalNote returns a boolean if a field has been set.
func (o *TerminalDetails) HasTerminalNote() bool {
	if o != nil && !IsNil(o.TerminalNote) {
		return true
	}

	return false
}

// SetTerminalNote gets a reference to the given string and assigns it to the TerminalNote field.
func (o *TerminalDetails) SetTerminalNote(v string) {
	o.TerminalNote = &v
}

// GetTerminalPhoneNo returns the TerminalPhoneNo field value
func (o *TerminalDetails) GetTerminalPhoneNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerminalPhoneNo
}

// GetTerminalPhoneNoOk returns a tuple with the TerminalPhoneNo field value
// and a boolean to check if the value has been set.
func (o *TerminalDetails) GetTerminalPhoneNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminalPhoneNo, true
}

// SetTerminalPhoneNo sets field value
func (o *TerminalDetails) SetTerminalPhoneNo(v string) {
	o.TerminalPhoneNo = v
}

// GetTerminalStatus returns the TerminalStatus field value if set, zero value otherwise.
func (o *TerminalDetails) GetTerminalStatus() string {
	if o == nil || IsNil(o.TerminalStatus) {
		var ret string
		return ret
	}
	return *o.TerminalStatus
}

// GetTerminalStatusOk returns a tuple with the TerminalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalDetails) GetTerminalStatusOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalStatus) {
		return nil, false
	}
	return o.TerminalStatus, true
}

// HasTerminalStatus returns a boolean if a field has been set.
func (o *TerminalDetails) HasTerminalStatus() bool {
	if o != nil && !IsNil(o.TerminalStatus) {
		return true
	}

	return false
}

// SetTerminalStatus gets a reference to the given string and assigns it to the TerminalStatus field.
func (o *TerminalDetails) SetTerminalStatus(v string) {
	o.TerminalStatus = &v
}

// GetTerminalType returns the TerminalType field value
func (o *TerminalDetails) GetTerminalType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerminalType
}

// GetTerminalTypeOk returns a tuple with the TerminalType field value
// and a boolean to check if the value has been set.
func (o *TerminalDetails) GetTerminalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminalType, true
}

// SetTerminalType sets field value
func (o *TerminalDetails) SetTerminalType(v string) {
	o.TerminalType = v
}

func (o TerminalDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddedOn) {
		toSerialize["added_on"] = o.AddedOn
	}
	if !IsNil(o.CfTerminalId) {
		toSerialize["cf_terminal_id"] = o.CfTerminalId
	}
	if !IsNil(o.LastUpdatedOn) {
		toSerialize["last_updated_on"] = o.LastUpdatedOn
	}
	if !IsNil(o.TerminalAddress) {
		toSerialize["terminal_address"] = o.TerminalAddress
	}
	toSerialize["terminal_id"] = o.TerminalId
	if !IsNil(o.TerminalName) {
		toSerialize["terminal_name"] = o.TerminalName
	}
	if !IsNil(o.TerminalNote) {
		toSerialize["terminal_note"] = o.TerminalNote
	}
	toSerialize["terminal_phone_no"] = o.TerminalPhoneNo
	if !IsNil(o.TerminalStatus) {
		toSerialize["terminal_status"] = o.TerminalStatus
	}
	toSerialize["terminal_type"] = o.TerminalType
	return toSerialize, nil
}

type NullableTerminalDetails struct {
	value *TerminalDetails
	isSet bool
}

func (v NullableTerminalDetails) Get() *TerminalDetails {
	return v.value
}

func (v *NullableTerminalDetails) Set(val *TerminalDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalDetails(val *TerminalDetails) *NullableTerminalDetails {
	return &NullableTerminalDetails{value: val, isSet: true}
}

func (v NullableTerminalDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


