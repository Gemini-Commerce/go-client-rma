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

// checks if the RmaCancelReturnRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaCancelReturnRequest{}

// RmaCancelReturnRequest struct for RmaCancelReturnRequest
type RmaCancelReturnRequest struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
}

type _RmaCancelReturnRequest RmaCancelReturnRequest

// NewRmaCancelReturnRequest instantiates a new RmaCancelReturnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaCancelReturnRequest(tenantId string, id string) *RmaCancelReturnRequest {
	this := RmaCancelReturnRequest{}
	this.TenantId = tenantId
	this.Id = id
	return &this
}

// NewRmaCancelReturnRequestWithDefaults instantiates a new RmaCancelReturnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaCancelReturnRequestWithDefaults() *RmaCancelReturnRequest {
	this := RmaCancelReturnRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *RmaCancelReturnRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *RmaCancelReturnRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *RmaCancelReturnRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *RmaCancelReturnRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RmaCancelReturnRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RmaCancelReturnRequest) SetId(v string) {
	o.Id = v
}

func (o RmaCancelReturnRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaCancelReturnRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *RmaCancelReturnRequest) UnmarshalJSON(data []byte) (err error) {
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

	varRmaCancelReturnRequest := _RmaCancelReturnRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRmaCancelReturnRequest)

	if err != nil {
		return err
	}

	*o = RmaCancelReturnRequest(varRmaCancelReturnRequest)

	return err
}

type NullableRmaCancelReturnRequest struct {
	value *RmaCancelReturnRequest
	isSet bool
}

func (v NullableRmaCancelReturnRequest) Get() *RmaCancelReturnRequest {
	return v.value
}

func (v *NullableRmaCancelReturnRequest) Set(val *RmaCancelReturnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaCancelReturnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaCancelReturnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaCancelReturnRequest(val *RmaCancelReturnRequest) *NullableRmaCancelReturnRequest {
	return &NullableRmaCancelReturnRequest{value: val, isSet: true}
}

func (v NullableRmaCancelReturnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaCancelReturnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

