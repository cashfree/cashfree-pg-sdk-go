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

// checks if the CreateLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLinkRequest{}

// CreateLinkRequest Request paramenters for link creation
type CreateLinkRequest struct {
	// Unique Identifier (provided by merchant) for the Link. Alphanumeric and only - and _ allowed (50 character limit). Use this for other link-related APIs.
	LinkId string `json:"link_id"`
	// Amount to be collected using this link. Provide upto two decimals for paise.
	LinkAmount float64 `json:"link_amount"`
	// Currency for the payment link. Default is INR. Contact care@cashfree.com to enable new currencies.
	LinkCurrency string `json:"link_currency"`
	// A brief description for which payment must be collected. This is shown to the customer.
	LinkPurpose string `json:"link_purpose"`
	CustomerDetails LinkCustomerDetailsEntity `json:"customer_details"`
	// If \"true\", customer can make partial payments for the link.
	LinkPartialPayments *bool `json:"link_partial_payments,omitempty"`
	// Minimum amount in first installment that needs to be paid by the customer if partial payments are enabled. This should be less than the link_amount.
	LinkMinimumPartialAmount *float64 `json:"link_minimum_partial_amount,omitempty"`
	// Time after which the link expires. Customers will not be able to make the payment beyond the time specified here. You can provide them in a valid ISO 8601 time format. Default is 30 days.
	LinkExpiryTime *string `json:"link_expiry_time,omitempty"`
	LinkNotify *LinkNotifyEntity `json:"link_notify,omitempty"`
	// If \"true\", reminders will be sent to customers for collecting payments.
	LinkAutoReminders *bool `json:"link_auto_reminders,omitempty"`
	// Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs
	LinkNotes *map[string]string `json:"link_notes,omitempty"`
	LinkMeta *LinkMetaEntity `json:"link_meta,omitempty"`
}

// NewCreateLinkRequest instantiates a new CreateLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLinkRequest(linkId string, linkAmount float64, linkCurrency string, linkPurpose string, customerDetails LinkCustomerDetailsEntity) *CreateLinkRequest {
	this := CreateLinkRequest{}
	this.LinkId = linkId
	this.LinkAmount = linkAmount
	this.LinkCurrency = linkCurrency
	this.LinkPurpose = linkPurpose
	this.CustomerDetails = customerDetails
	return &this
}

// NewCreateLinkRequestWithDefaults instantiates a new CreateLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLinkRequestWithDefaults() *CreateLinkRequest {
	this := CreateLinkRequest{}
	return &this
}

// GetLinkId returns the LinkId field value
func (o *CreateLinkRequest) GetLinkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value
// and a boolean to check if the value has been set.
func (o *CreateLinkRequest) GetLinkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkId, true
}

// SetLinkId sets field value
func (o *CreateLinkRequest) SetLinkId(v string) {
	o.LinkId = v
}

// GetLinkAmount returns the LinkAmount field value
func (o *CreateLinkRequest) GetLinkAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.LinkAmount
}

// GetLinkAmountOk returns a tuple with the LinkAmount field value
// and a boolean to check if the value has been set.
func (o *CreateLinkRequest) GetLinkAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkAmount, true
}

// SetLinkAmount sets field value
func (o *CreateLinkRequest) SetLinkAmount(v float64) {
	o.LinkAmount = v
}

// GetLinkCurrency returns the LinkCurrency field value
func (o *CreateLinkRequest) GetLinkCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkCurrency
}

// GetLinkCurrencyOk returns a tuple with the LinkCurrency field value
// and a boolean to check if the value has been set.
func (o *CreateLinkRequest) GetLinkCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkCurrency, true
}

// SetLinkCurrency sets field value
func (o *CreateLinkRequest) SetLinkCurrency(v string) {
	o.LinkCurrency = v
}

// GetLinkPurpose returns the LinkPurpose field value
func (o *CreateLinkRequest) GetLinkPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkPurpose
}

// GetLinkPurposeOk returns a tuple with the LinkPurpose field value
// and a boolean to check if the value has been set.
func (o *CreateLinkRequest) GetLinkPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkPurpose, true
}

// SetLinkPurpose sets field value
func (o *CreateLinkRequest) SetLinkPurpose(v string) {
	o.LinkPurpose = v
}

// GetCustomerDetails returns the CustomerDetails field value
func (o *CreateLinkRequest) GetCustomerDetails() LinkCustomerDetailsEntity {
	if o == nil {
		var ret LinkCustomerDetailsEntity
		return ret
	}

	return o.CustomerDetails
}

// GetCustomerDetailsOk returns a tuple with the CustomerDetails field value
// and a boolean to check if the value has been set.
func (o *CreateLinkRequest) GetCustomerDetailsOk() (*LinkCustomerDetailsEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerDetails, true
}

// SetCustomerDetails sets field value
func (o *CreateLinkRequest) SetCustomerDetails(v LinkCustomerDetailsEntity) {
	o.CustomerDetails = v
}

// GetLinkPartialPayments returns the LinkPartialPayments field value if set, zero value otherwise.
func (o *CreateLinkRequest) GetLinkPartialPayments() bool {
	if o == nil || IsNil(o.LinkPartialPayments) {
		var ret bool
		return ret
	}
	return *o.LinkPartialPayments
}

// GetLinkPartialPaymentsOk returns a tuple with the LinkPartialPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkRequest) GetLinkPartialPaymentsOk() (*bool, bool) {
	if o == nil || IsNil(o.LinkPartialPayments) {
		return nil, false
	}
	return o.LinkPartialPayments, true
}

// HasLinkPartialPayments returns a boolean if a field has been set.
func (o *CreateLinkRequest) HasLinkPartialPayments() bool {
	if o != nil && !IsNil(o.LinkPartialPayments) {
		return true
	}

	return false
}

// SetLinkPartialPayments gets a reference to the given bool and assigns it to the LinkPartialPayments field.
func (o *CreateLinkRequest) SetLinkPartialPayments(v bool) {
	o.LinkPartialPayments = &v
}

// GetLinkMinimumPartialAmount returns the LinkMinimumPartialAmount field value if set, zero value otherwise.
func (o *CreateLinkRequest) GetLinkMinimumPartialAmount() float64 {
	if o == nil || IsNil(o.LinkMinimumPartialAmount) {
		var ret float64
		return ret
	}
	return *o.LinkMinimumPartialAmount
}

// GetLinkMinimumPartialAmountOk returns a tuple with the LinkMinimumPartialAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkRequest) GetLinkMinimumPartialAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.LinkMinimumPartialAmount) {
		return nil, false
	}
	return o.LinkMinimumPartialAmount, true
}

// HasLinkMinimumPartialAmount returns a boolean if a field has been set.
func (o *CreateLinkRequest) HasLinkMinimumPartialAmount() bool {
	if o != nil && !IsNil(o.LinkMinimumPartialAmount) {
		return true
	}

	return false
}

// SetLinkMinimumPartialAmount gets a reference to the given float64 and assigns it to the LinkMinimumPartialAmount field.
func (o *CreateLinkRequest) SetLinkMinimumPartialAmount(v float64) {
	o.LinkMinimumPartialAmount = &v
}

// GetLinkExpiryTime returns the LinkExpiryTime field value if set, zero value otherwise.
func (o *CreateLinkRequest) GetLinkExpiryTime() string {
	if o == nil || IsNil(o.LinkExpiryTime) {
		var ret string
		return ret
	}
	return *o.LinkExpiryTime
}

// GetLinkExpiryTimeOk returns a tuple with the LinkExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkRequest) GetLinkExpiryTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LinkExpiryTime) {
		return nil, false
	}
	return o.LinkExpiryTime, true
}

// HasLinkExpiryTime returns a boolean if a field has been set.
func (o *CreateLinkRequest) HasLinkExpiryTime() bool {
	if o != nil && !IsNil(o.LinkExpiryTime) {
		return true
	}

	return false
}

// SetLinkExpiryTime gets a reference to the given string and assigns it to the LinkExpiryTime field.
func (o *CreateLinkRequest) SetLinkExpiryTime(v string) {
	o.LinkExpiryTime = &v
}

// GetLinkNotify returns the LinkNotify field value if set, zero value otherwise.
func (o *CreateLinkRequest) GetLinkNotify() LinkNotifyEntity {
	if o == nil || IsNil(o.LinkNotify) {
		var ret LinkNotifyEntity
		return ret
	}
	return *o.LinkNotify
}

// GetLinkNotifyOk returns a tuple with the LinkNotify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkRequest) GetLinkNotifyOk() (*LinkNotifyEntity, bool) {
	if o == nil || IsNil(o.LinkNotify) {
		return nil, false
	}
	return o.LinkNotify, true
}

// HasLinkNotify returns a boolean if a field has been set.
func (o *CreateLinkRequest) HasLinkNotify() bool {
	if o != nil && !IsNil(o.LinkNotify) {
		return true
	}

	return false
}

// SetLinkNotify gets a reference to the given LinkNotifyEntity and assigns it to the LinkNotify field.
func (o *CreateLinkRequest) SetLinkNotify(v LinkNotifyEntity) {
	o.LinkNotify = &v
}

// GetLinkAutoReminders returns the LinkAutoReminders field value if set, zero value otherwise.
func (o *CreateLinkRequest) GetLinkAutoReminders() bool {
	if o == nil || IsNil(o.LinkAutoReminders) {
		var ret bool
		return ret
	}
	return *o.LinkAutoReminders
}

// GetLinkAutoRemindersOk returns a tuple with the LinkAutoReminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkRequest) GetLinkAutoRemindersOk() (*bool, bool) {
	if o == nil || IsNil(o.LinkAutoReminders) {
		return nil, false
	}
	return o.LinkAutoReminders, true
}

// HasLinkAutoReminders returns a boolean if a field has been set.
func (o *CreateLinkRequest) HasLinkAutoReminders() bool {
	if o != nil && !IsNil(o.LinkAutoReminders) {
		return true
	}

	return false
}

// SetLinkAutoReminders gets a reference to the given bool and assigns it to the LinkAutoReminders field.
func (o *CreateLinkRequest) SetLinkAutoReminders(v bool) {
	o.LinkAutoReminders = &v
}

// GetLinkNotes returns the LinkNotes field value if set, zero value otherwise.
func (o *CreateLinkRequest) GetLinkNotes() map[string]string {
	if o == nil || IsNil(o.LinkNotes) {
		var ret map[string]string
		return ret
	}
	return *o.LinkNotes
}

// GetLinkNotesOk returns a tuple with the LinkNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkRequest) GetLinkNotesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.LinkNotes) {
		return nil, false
	}
	return o.LinkNotes, true
}

// HasLinkNotes returns a boolean if a field has been set.
func (o *CreateLinkRequest) HasLinkNotes() bool {
	if o != nil && !IsNil(o.LinkNotes) {
		return true
	}

	return false
}

// SetLinkNotes gets a reference to the given map[string]string and assigns it to the LinkNotes field.
func (o *CreateLinkRequest) SetLinkNotes(v map[string]string) {
	o.LinkNotes = &v
}

// GetLinkMeta returns the LinkMeta field value if set, zero value otherwise.
func (o *CreateLinkRequest) GetLinkMeta() LinkMetaEntity {
	if o == nil || IsNil(o.LinkMeta) {
		var ret LinkMetaEntity
		return ret
	}
	return *o.LinkMeta
}

// GetLinkMetaOk returns a tuple with the LinkMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkRequest) GetLinkMetaOk() (*LinkMetaEntity, bool) {
	if o == nil || IsNil(o.LinkMeta) {
		return nil, false
	}
	return o.LinkMeta, true
}

// HasLinkMeta returns a boolean if a field has been set.
func (o *CreateLinkRequest) HasLinkMeta() bool {
	if o != nil && !IsNil(o.LinkMeta) {
		return true
	}

	return false
}

// SetLinkMeta gets a reference to the given LinkMetaEntity and assigns it to the LinkMeta field.
func (o *CreateLinkRequest) SetLinkMeta(v LinkMetaEntity) {
	o.LinkMeta = &v
}

func (o CreateLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["link_id"] = o.LinkId
	toSerialize["link_amount"] = o.LinkAmount
	toSerialize["link_currency"] = o.LinkCurrency
	toSerialize["link_purpose"] = o.LinkPurpose
	toSerialize["customer_details"] = o.CustomerDetails
	if !IsNil(o.LinkPartialPayments) {
		toSerialize["link_partial_payments"] = o.LinkPartialPayments
	}
	if !IsNil(o.LinkMinimumPartialAmount) {
		toSerialize["link_minimum_partial_amount"] = o.LinkMinimumPartialAmount
	}
	if !IsNil(o.LinkExpiryTime) {
		toSerialize["link_expiry_time"] = o.LinkExpiryTime
	}
	if !IsNil(o.LinkNotify) {
		toSerialize["link_notify"] = o.LinkNotify
	}
	if !IsNil(o.LinkAutoReminders) {
		toSerialize["link_auto_reminders"] = o.LinkAutoReminders
	}
	if !IsNil(o.LinkNotes) {
		toSerialize["link_notes"] = o.LinkNotes
	}
	if !IsNil(o.LinkMeta) {
		toSerialize["link_meta"] = o.LinkMeta
	}
	return toSerialize, nil
}

type NullableCreateLinkRequest struct {
	value *CreateLinkRequest
	isSet bool
}

func (v NullableCreateLinkRequest) Get() *CreateLinkRequest {
	return v.value
}

func (v *NullableCreateLinkRequest) Set(val *CreateLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLinkRequest(val *CreateLinkRequest) *NullableCreateLinkRequest {
	return &NullableCreateLinkRequest{value: val, isSet: true}
}

func (v NullableCreateLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


