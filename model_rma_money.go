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
)

// checks if the RmaMoney type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaMoney{}

// RmaMoney struct for RmaMoney
type RmaMoney struct {
	// The whole units of the amount. For example if `currencyCode` is `\"USD\"`, then 1 unit is one US dollar.
	Units *string `json:"units,omitempty"`
	// Number of micro (10^-6) units of the amount. The value must be between -999,999 and +999,999 inclusive. If `units` is positive, `micros` must be positive or zero. If `units` is zero, `micros` can be positive, zero, or negative. If `units` is negative, `micros` must be negative or zero. For example $-1.75 is represented as `units`=-1 and `micros`=-750,000.
	Micros *int32 `json:"micros,omitempty"`
}

// NewRmaMoney instantiates a new RmaMoney object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaMoney() *RmaMoney {
	this := RmaMoney{}
	return &this
}

// NewRmaMoneyWithDefaults instantiates a new RmaMoney object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaMoneyWithDefaults() *RmaMoney {
	this := RmaMoney{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *RmaMoney) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaMoney) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *RmaMoney) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *RmaMoney) SetUnits(v string) {
	o.Units = &v
}

// GetMicros returns the Micros field value if set, zero value otherwise.
func (o *RmaMoney) GetMicros() int32 {
	if o == nil || IsNil(o.Micros) {
		var ret int32
		return ret
	}
	return *o.Micros
}

// GetMicrosOk returns a tuple with the Micros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaMoney) GetMicrosOk() (*int32, bool) {
	if o == nil || IsNil(o.Micros) {
		return nil, false
	}
	return o.Micros, true
}

// HasMicros returns a boolean if a field has been set.
func (o *RmaMoney) HasMicros() bool {
	if o != nil && !IsNil(o.Micros) {
		return true
	}

	return false
}

// SetMicros gets a reference to the given int32 and assigns it to the Micros field.
func (o *RmaMoney) SetMicros(v int32) {
	o.Micros = &v
}

func (o RmaMoney) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaMoney) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if !IsNil(o.Micros) {
		toSerialize["micros"] = o.Micros
	}
	return toSerialize, nil
}

type NullableRmaMoney struct {
	value *RmaMoney
	isSet bool
}

func (v NullableRmaMoney) Get() *RmaMoney {
	return v.value
}

func (v *NullableRmaMoney) Set(val *RmaMoney) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaMoney) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaMoney) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaMoney(val *RmaMoney) *NullableRmaMoney {
	return &NullableRmaMoney{value: val, isSet: true}
}

func (v NullableRmaMoney) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaMoney) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

