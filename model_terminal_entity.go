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

// checks if the TerminalEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminalEntity{}

// TerminalEntity Create terminal response object
type TerminalEntity struct {
	AddedOn *string `json:"added_on,omitempty"`
	CfTerminalId *int32 `json:"cf_terminal_id,omitempty"`
	LastUpdatedOn *string `json:"last_updated_on,omitempty"`
	TerminalAddress *string `json:"terminal_address,omitempty"`
	TerminalEmail *string `json:"terminal_email,omitempty"`
	TerminalType *string `json:"terminal_type,omitempty"`
	TeminalId *string `json:"teminal_id,omitempty"`
	TerminalName *string `json:"terminal_name,omitempty"`
	TerminalNote *string `json:"terminal_note,omitempty"`
	TerminalPhoneNo *string `json:"terminal_phone_no,omitempty"`
	TerminalStatus *string `json:"terminal_status,omitempty"`
	TerminalMeta *string `json:"terminal_meta,omitempty"`
}

// NewTerminalEntity instantiates a new TerminalEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalEntity() *TerminalEntity {
	this := TerminalEntity{}
	return &this
}

// NewTerminalEntityWithDefaults instantiates a new TerminalEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalEntityWithDefaults() *TerminalEntity {
	this := TerminalEntity{}
	return &this
}

// GetAddedOn returns the AddedOn field value if set, zero value otherwise.
func (o *TerminalEntity) GetAddedOn() string {
	if o == nil || IsNil(o.AddedOn) {
		var ret string
		return ret
	}
	return *o.AddedOn
}

// GetAddedOnOk returns a tuple with the AddedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalEntity) GetAddedOnOk() (*string, bool) {
	if o == nil || IsNil(o.AddedOn) {
		return nil, false
	}
	return o.AddedOn, true
}

// HasAddedOn returns a boolean if a field has been set.
func (o *TerminalEntity) HasAddedOn() bool {
	if o != nil && !IsNil(o.AddedOn) {
		return true
	}

	return false
}

// SetAddedOn gets a reference to the given string and assigns it to the AddedOn field.
func (o *TerminalEntity) SetAddedOn(v string) {
	o.AddedOn = &v
}

// GetCfTerminalId returns the CfTerminalId field value if set, zero value otherwise.
func (o *TerminalEntity) GetCfTerminalId() int32 {
	if o == nil || IsNil(o.CfTerminalId) {
		var ret int32
		return ret
	}
	return *o.CfTerminalId
}

// GetCfTerminalIdOk returns a tuple with the CfTerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalEntity) GetCfTerminalIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CfTerminalId) {
		return nil, false
	}
	return o.CfTerminalId, true
}

// HasCfTerminalId returns a boolean if a field has been set.
func (o *TerminalEntity) HasCfTerminalId() bool {
	if o != nil && !IsNil(o.CfTerminalId) {
		return true
	}

	return false
}

// SetCfTerminalId gets a reference to the given int32 and assigns it to the CfTerminalId field.
func (o *TerminalEntity) SetCfTerminalId(v int32) {
	o.CfTerminalId = &v
}

// GetLastUpdatedOn returns the LastUpdatedOn field value if set, zero value otherwise.
func (o *TerminalEntity) GetLastUpdatedOn() string {
	if o == nil || IsNil(o.LastUpdatedOn) {
		var ret string
		return ret
	}
	return *o.LastUpdatedOn
}

// GetLastUpdatedOnOk returns a tuple with the LastUpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalEntity) GetLastUpdatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedOn) {
		return nil, false
	}
	return o.LastUpdatedOn, true
}

// HasLastUpdatedOn returns a boolean if a field has been set.
func (o *TerminalEntity) HasLastUpdatedOn() bool {
	if o != nil && !IsNil(o.LastUpdatedOn) {
		return true
	}

	return false
}

// SetLastUpdatedOn gets a reference to the given string and assigns it to the LastUpdatedOn field.
func (o *TerminalEntity) SetLastUpdatedOn(v string) {
	o.LastUpdatedOn = &v
}

// GetTerminalAddress returns the TerminalAddress field value if set, zero value otherwise.
func (o *TerminalEntity) GetTerminalAddress() string {
	if o == nil || IsNil(o.TerminalAddress) {
		var ret string
		return ret
	}
	return *o.TerminalAddress
}

// GetTerminalAddressOk returns a tuple with the TerminalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalEntity) GetTerminalAddressOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalAddress) {
		return nil, false
	}
	return o.TerminalAddress, true
}

// HasTerminalAddress returns a boolean if a field has been set.
func (o *TerminalEntity) HasTerminalAddress() bool {
	if o != nil && !IsNil(o.TerminalAddress) {
		return true
	}

	return false
}

// SetTerminalAddress gets a reference to the given string and assigns it to the TerminalAddress field.
func (o *TerminalEntity) SetTerminalAddress(v string) {
	o.TerminalAddress = &v
}

// GetTerminalEmail returns the TerminalEmail field value if set, zero value otherwise.
func (o *TerminalEntity) GetTerminalEmail() string {
	if o == nil || IsNil(o.TerminalEmail) {
		var ret string
		return ret
	}
	return *o.TerminalEmail
}

// GetTerminalEmailOk returns a tuple with the TerminalEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalEntity) GetTerminalEmailOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalEmail) {
		return nil, false
	}
	return o.TerminalEmail, true
}

// HasTerminalEmail returns a boolean if a field has been set.
func (o *TerminalEntity) HasTerminalEmail() bool {
	if o != nil && !IsNil(o.TerminalEmail) {
		return true
	}

	return false
}

// SetTerminalEmail gets a reference to the given string and assigns it to the TerminalEmail field.
func (o *TerminalEntity) SetTerminalEmail(v string) {
	o.TerminalEmail = &v
}

// GetTerminalType returns the TerminalType field value if set, zero value otherwise.
func (o *TerminalEntity) GetTerminalType() string {
	if o == nil || IsNil(o.TerminalType) {
		var ret string
		return ret
	}
	return *o.TerminalType
}

// GetTerminalTypeOk returns a tuple with the TerminalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalEntity) GetTerminalTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalType) {
		return nil, false
	}
	return o.TerminalType, true
}

// HasTerminalType returns a boolean if a field has been set.
func (o *TerminalEntity) HasTerminalType() bool {
	if o != nil && !IsNil(o.TerminalType) {
		return true
	}

	return false
}

// SetTerminalType gets a reference to the given string and assigns it to the TerminalType field.
func (o *TerminalEntity) SetTerminalType(v string) {
	o.TerminalType = &v
}

// GetTeminalId returns the TeminalId field value if set, zero value otherwise.
func (o *TerminalEntity) GetTeminalId() string {
	if o == nil || IsNil(o.TeminalId) {
		var ret string
		return ret
	}
	return *o.TeminalId
}

// GetTeminalIdOk returns a tuple with the TeminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalEntity) GetTeminalIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeminalId) {
		return nil, false
	}
	return o.TeminalId, true
}

// HasTeminalId returns a boolean if a field has been set.
func (o *TerminalEntity) HasTeminalId() bool {
	if o != nil && !IsNil(o.TeminalId) {
		return true
	}

	return false
}

// SetTeminalId gets a reference to the given string and assigns it to the TeminalId field.
func (o *TerminalEntity) SetTeminalId(v string) {
	o.TeminalId = &v
}

// GetTerminalName returns the TerminalName field value if set, zero value otherwise.
func (o *TerminalEntity) GetTerminalName() string {
	if o == nil || IsNil(o.TerminalName) {
		var ret string
		return ret
	}
	return *o.TerminalName
}

// GetTerminalNameOk returns a tuple with the TerminalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalEntity) GetTerminalNameOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalName) {
		return nil, false
	}
	return o.TerminalName, true
}

// HasTerminalName returns a boolean if a field has been set.
func (o *TerminalEntity) HasTerminalName() bool {
	if o != nil && !IsNil(o.TerminalName) {
		return true
	}

	return false
}

// SetTerminalName gets a reference to the given string and assigns it to the TerminalName field.
func (o *TerminalEntity) SetTerminalName(v string) {
	o.TerminalName = &v
}

// GetTerminalNote returns the TerminalNote field value if set, zero value otherwise.
func (o *TerminalEntity) GetTerminalNote() string {
	if o == nil || IsNil(o.TerminalNote) {
		var ret string
		return ret
	}
	return *o.TerminalNote
}

// GetTerminalNoteOk returns a tuple with the TerminalNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalEntity) GetTerminalNoteOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalNote) {
		return nil, false
	}
	return o.TerminalNote, true
}

// HasTerminalNote returns a boolean if a field has been set.
func (o *TerminalEntity) HasTerminalNote() bool {
	if o != nil && !IsNil(o.TerminalNote) {
		return true
	}

	return false
}

// SetTerminalNote gets a reference to the given string and assigns it to the TerminalNote field.
func (o *TerminalEntity) SetTerminalNote(v string) {
	o.TerminalNote = &v
}

// GetTerminalPhoneNo returns the TerminalPhoneNo field value if set, zero value otherwise.
func (o *TerminalEntity) GetTerminalPhoneNo() string {
	if o == nil || IsNil(o.TerminalPhoneNo) {
		var ret string
		return ret
	}
	return *o.TerminalPhoneNo
}

// GetTerminalPhoneNoOk returns a tuple with the TerminalPhoneNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalEntity) GetTerminalPhoneNoOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalPhoneNo) {
		return nil, false
	}
	return o.TerminalPhoneNo, true
}

// HasTerminalPhoneNo returns a boolean if a field has been set.
func (o *TerminalEntity) HasTerminalPhoneNo() bool {
	if o != nil && !IsNil(o.TerminalPhoneNo) {
		return true
	}

	return false
}

// SetTerminalPhoneNo gets a reference to the given string and assigns it to the TerminalPhoneNo field.
func (o *TerminalEntity) SetTerminalPhoneNo(v string) {
	o.TerminalPhoneNo = &v
}

// GetTerminalStatus returns the TerminalStatus field value if set, zero value otherwise.
func (o *TerminalEntity) GetTerminalStatus() string {
	if o == nil || IsNil(o.TerminalStatus) {
		var ret string
		return ret
	}
	return *o.TerminalStatus
}

// GetTerminalStatusOk returns a tuple with the TerminalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalEntity) GetTerminalStatusOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalStatus) {
		return nil, false
	}
	return o.TerminalStatus, true
}

// HasTerminalStatus returns a boolean if a field has been set.
func (o *TerminalEntity) HasTerminalStatus() bool {
	if o != nil && !IsNil(o.TerminalStatus) {
		return true
	}

	return false
}

// SetTerminalStatus gets a reference to the given string and assigns it to the TerminalStatus field.
func (o *TerminalEntity) SetTerminalStatus(v string) {
	o.TerminalStatus = &v
}

// GetTerminalMeta returns the TerminalMeta field value if set, zero value otherwise.
func (o *TerminalEntity) GetTerminalMeta() string {
	if o == nil || IsNil(o.TerminalMeta) {
		var ret string
		return ret
	}
	return *o.TerminalMeta
}

// GetTerminalMetaOk returns a tuple with the TerminalMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalEntity) GetTerminalMetaOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalMeta) {
		return nil, false
	}
	return o.TerminalMeta, true
}

// HasTerminalMeta returns a boolean if a field has been set.
func (o *TerminalEntity) HasTerminalMeta() bool {
	if o != nil && !IsNil(o.TerminalMeta) {
		return true
	}

	return false
}

// SetTerminalMeta gets a reference to the given string and assigns it to the TerminalMeta field.
func (o *TerminalEntity) SetTerminalMeta(v string) {
	o.TerminalMeta = &v
}

func (o TerminalEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalEntity) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.TerminalEmail) {
		toSerialize["terminal_email"] = o.TerminalEmail
	}
	if !IsNil(o.TerminalType) {
		toSerialize["terminal_type"] = o.TerminalType
	}
	if !IsNil(o.TeminalId) {
		toSerialize["teminal_id"] = o.TeminalId
	}
	if !IsNil(o.TerminalName) {
		toSerialize["terminal_name"] = o.TerminalName
	}
	if !IsNil(o.TerminalNote) {
		toSerialize["terminal_note"] = o.TerminalNote
	}
	if !IsNil(o.TerminalPhoneNo) {
		toSerialize["terminal_phone_no"] = o.TerminalPhoneNo
	}
	if !IsNil(o.TerminalStatus) {
		toSerialize["terminal_status"] = o.TerminalStatus
	}
	if !IsNil(o.TerminalMeta) {
		toSerialize["terminal_meta"] = o.TerminalMeta
	}
	return toSerialize, nil
}

type NullableTerminalEntity struct {
	value *TerminalEntity
	isSet bool
}

func (v NullableTerminalEntity) Get() *TerminalEntity {
	return v.value
}

func (v *NullableTerminalEntity) Set(val *TerminalEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalEntity(val *TerminalEntity) *NullableTerminalEntity {
	return &NullableTerminalEntity{value: val, isSet: true}
}

func (v NullableTerminalEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


