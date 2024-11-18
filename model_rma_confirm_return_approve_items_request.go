/*
RMA Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rma

import (
	"encoding/json"
	"fmt"
)

// checks if the RmaConfirmReturnApproveItemsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaConfirmReturnApproveItemsRequest{}

// RmaConfirmReturnApproveItemsRequest struct for RmaConfirmReturnApproveItemsRequest
type RmaConfirmReturnApproveItemsRequest struct {
	TenantId             string                                    `json:"tenantId"`
	Id                   string                                    `json:"id"`
	Items                []RmaConfirmReturnApproveItemsRequestItem `json:"items"`
	AdditionalProperties map[string]interface{}
}

type _RmaConfirmReturnApproveItemsRequest RmaConfirmReturnApproveItemsRequest

// NewRmaConfirmReturnApproveItemsRequest instantiates a new RmaConfirmReturnApproveItemsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaConfirmReturnApproveItemsRequest(tenantId string, id string, items []RmaConfirmReturnApproveItemsRequestItem) *RmaConfirmReturnApproveItemsRequest {
	this := RmaConfirmReturnApproveItemsRequest{}
	this.TenantId = tenantId
	this.Id = id
	this.Items = items
	return &this
}

// NewRmaConfirmReturnApproveItemsRequestWithDefaults instantiates a new RmaConfirmReturnApproveItemsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaConfirmReturnApproveItemsRequestWithDefaults() *RmaConfirmReturnApproveItemsRequest {
	this := RmaConfirmReturnApproveItemsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *RmaConfirmReturnApproveItemsRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *RmaConfirmReturnApproveItemsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *RmaConfirmReturnApproveItemsRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *RmaConfirmReturnApproveItemsRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RmaConfirmReturnApproveItemsRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RmaConfirmReturnApproveItemsRequest) SetId(v string) {
	o.Id = v
}

// GetItems returns the Items field value
func (o *RmaConfirmReturnApproveItemsRequest) GetItems() []RmaConfirmReturnApproveItemsRequestItem {
	if o == nil {
		var ret []RmaConfirmReturnApproveItemsRequestItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RmaConfirmReturnApproveItemsRequest) GetItemsOk() ([]RmaConfirmReturnApproveItemsRequestItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *RmaConfirmReturnApproveItemsRequest) SetItems(v []RmaConfirmReturnApproveItemsRequestItem) {
	o.Items = v
}

func (o RmaConfirmReturnApproveItemsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaConfirmReturnApproveItemsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaConfirmReturnApproveItemsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"items",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRmaConfirmReturnApproveItemsRequest := _RmaConfirmReturnApproveItemsRequest{}

	err = json.Unmarshal(data, &varRmaConfirmReturnApproveItemsRequest)

	if err != nil {
		return err
	}

	*o = RmaConfirmReturnApproveItemsRequest(varRmaConfirmReturnApproveItemsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaConfirmReturnApproveItemsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *RmaConfirmReturnApproveItemsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableRmaConfirmReturnApproveItemsRequest struct {
	value *RmaConfirmReturnApproveItemsRequest
	isSet bool
}

func (v NullableRmaConfirmReturnApproveItemsRequest) Get() *RmaConfirmReturnApproveItemsRequest {
	return v.value
}

func (v *NullableRmaConfirmReturnApproveItemsRequest) Set(val *RmaConfirmReturnApproveItemsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaConfirmReturnApproveItemsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaConfirmReturnApproveItemsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaConfirmReturnApproveItemsRequest(val *RmaConfirmReturnApproveItemsRequest) *NullableRmaConfirmReturnApproveItemsRequest {
	return &NullableRmaConfirmReturnApproveItemsRequest{value: val, isSet: true}
}

func (v NullableRmaConfirmReturnApproveItemsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaConfirmReturnApproveItemsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
