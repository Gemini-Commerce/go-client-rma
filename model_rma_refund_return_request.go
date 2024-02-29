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
	"bytes"
	"fmt"
)

// checks if the RmaRefundReturnRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaRefundReturnRequest{}

// RmaRefundReturnRequest struct for RmaRefundReturnRequest
type RmaRefundReturnRequest struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
}

type _RmaRefundReturnRequest RmaRefundReturnRequest

// NewRmaRefundReturnRequest instantiates a new RmaRefundReturnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaRefundReturnRequest(tenantId string, id string) *RmaRefundReturnRequest {
	this := RmaRefundReturnRequest{}
	this.TenantId = tenantId
	this.Id = id
	return &this
}

// NewRmaRefundReturnRequestWithDefaults instantiates a new RmaRefundReturnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaRefundReturnRequestWithDefaults() *RmaRefundReturnRequest {
	this := RmaRefundReturnRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *RmaRefundReturnRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *RmaRefundReturnRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *RmaRefundReturnRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *RmaRefundReturnRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RmaRefundReturnRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RmaRefundReturnRequest) SetId(v string) {
	o.Id = v
}

func (o RmaRefundReturnRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaRefundReturnRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *RmaRefundReturnRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRmaRefundReturnRequest := _RmaRefundReturnRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRmaRefundReturnRequest)

	if err != nil {
		return err
	}

	*o = RmaRefundReturnRequest(varRmaRefundReturnRequest)

	return err
}

type NullableRmaRefundReturnRequest struct {
	value *RmaRefundReturnRequest
	isSet bool
}

func (v NullableRmaRefundReturnRequest) Get() *RmaRefundReturnRequest {
	return v.value
}

func (v *NullableRmaRefundReturnRequest) Set(val *RmaRefundReturnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaRefundReturnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaRefundReturnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaRefundReturnRequest(val *RmaRefundReturnRequest) *NullableRmaRefundReturnRequest {
	return &NullableRmaRefundReturnRequest{value: val, isSet: true}
}

func (v NullableRmaRefundReturnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaRefundReturnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

