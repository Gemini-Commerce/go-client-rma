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

// checks if the RmaApproveReturnRequestItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaApproveReturnRequestItem{}

// RmaApproveReturnRequestItem struct for RmaApproveReturnRequestItem
type RmaApproveReturnRequestItem struct {
	Grn string `json:"grn"`
	Quantity string `json:"quantity"`
	AdditionalProperties map[string]interface{}
}

type _RmaApproveReturnRequestItem RmaApproveReturnRequestItem

// NewRmaApproveReturnRequestItem instantiates a new RmaApproveReturnRequestItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaApproveReturnRequestItem(grn string, quantity string) *RmaApproveReturnRequestItem {
	this := RmaApproveReturnRequestItem{}
	this.Grn = grn
	this.Quantity = quantity
	return &this
}

// NewRmaApproveReturnRequestItemWithDefaults instantiates a new RmaApproveReturnRequestItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaApproveReturnRequestItemWithDefaults() *RmaApproveReturnRequestItem {
	this := RmaApproveReturnRequestItem{}
	return &this
}

// GetGrn returns the Grn field value
func (o *RmaApproveReturnRequestItem) GetGrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Grn
}

// GetGrnOk returns a tuple with the Grn field value
// and a boolean to check if the value has been set.
func (o *RmaApproveReturnRequestItem) GetGrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grn, true
}

// SetGrn sets field value
func (o *RmaApproveReturnRequestItem) SetGrn(v string) {
	o.Grn = v
}

// GetQuantity returns the Quantity field value
func (o *RmaApproveReturnRequestItem) GetQuantity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *RmaApproveReturnRequestItem) GetQuantityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *RmaApproveReturnRequestItem) SetQuantity(v string) {
	o.Quantity = v
}

func (o RmaApproveReturnRequestItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaApproveReturnRequestItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grn"] = o.Grn
	toSerialize["quantity"] = o.Quantity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaApproveReturnRequestItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grn",
		"quantity",
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

	varRmaApproveReturnRequestItem := _RmaApproveReturnRequestItem{}

	err = json.Unmarshal(data, &varRmaApproveReturnRequestItem)

	if err != nil {
		return err
	}

	*o = RmaApproveReturnRequestItem(varRmaApproveReturnRequestItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grn")
		delete(additionalProperties, "quantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaApproveReturnRequestItem) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *RmaApproveReturnRequestItem) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableRmaApproveReturnRequestItem struct {
	value *RmaApproveReturnRequestItem
	isSet bool
}

func (v NullableRmaApproveReturnRequestItem) Get() *RmaApproveReturnRequestItem {
	return v.value
}

func (v *NullableRmaApproveReturnRequestItem) Set(val *RmaApproveReturnRequestItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaApproveReturnRequestItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaApproveReturnRequestItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaApproveReturnRequestItem(val *RmaApproveReturnRequestItem) *NullableRmaApproveReturnRequestItem {
	return &NullableRmaApproveReturnRequestItem{value: val, isSet: true}
}

func (v NullableRmaApproveReturnRequestItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaApproveReturnRequestItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


