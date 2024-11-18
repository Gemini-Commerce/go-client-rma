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

// checks if the RmaSkipReturnStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaSkipReturnStatusRequest{}

// RmaSkipReturnStatusRequest struct for RmaSkipReturnStatusRequest
type RmaSkipReturnStatusRequest struct {
	TenantId             string `json:"tenantId"`
	Id                   string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _RmaSkipReturnStatusRequest RmaSkipReturnStatusRequest

// NewRmaSkipReturnStatusRequest instantiates a new RmaSkipReturnStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaSkipReturnStatusRequest(tenantId string, id string) *RmaSkipReturnStatusRequest {
	this := RmaSkipReturnStatusRequest{}
	this.TenantId = tenantId
	this.Id = id
	return &this
}

// NewRmaSkipReturnStatusRequestWithDefaults instantiates a new RmaSkipReturnStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaSkipReturnStatusRequestWithDefaults() *RmaSkipReturnStatusRequest {
	this := RmaSkipReturnStatusRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *RmaSkipReturnStatusRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *RmaSkipReturnStatusRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *RmaSkipReturnStatusRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *RmaSkipReturnStatusRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RmaSkipReturnStatusRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RmaSkipReturnStatusRequest) SetId(v string) {
	o.Id = v
}

func (o RmaSkipReturnStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaSkipReturnStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaSkipReturnStatusRequest) UnmarshalJSON(data []byte) (err error) {
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
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRmaSkipReturnStatusRequest := _RmaSkipReturnStatusRequest{}

	err = json.Unmarshal(data, &varRmaSkipReturnStatusRequest)

	if err != nil {
		return err
	}

	*o = RmaSkipReturnStatusRequest(varRmaSkipReturnStatusRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaSkipReturnStatusRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *RmaSkipReturnStatusRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableRmaSkipReturnStatusRequest struct {
	value *RmaSkipReturnStatusRequest
	isSet bool
}

func (v NullableRmaSkipReturnStatusRequest) Get() *RmaSkipReturnStatusRequest {
	return v.value
}

func (v *NullableRmaSkipReturnStatusRequest) Set(val *RmaSkipReturnStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaSkipReturnStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaSkipReturnStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaSkipReturnStatusRequest(val *RmaSkipReturnStatusRequest) *NullableRmaSkipReturnStatusRequest {
	return &NullableRmaSkipReturnStatusRequest{value: val, isSet: true}
}

func (v NullableRmaSkipReturnStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaSkipReturnStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
