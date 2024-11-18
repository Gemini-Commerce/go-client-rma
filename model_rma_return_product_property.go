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

// checks if the RmaReturnProductProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaReturnProductProperty{}

// RmaReturnProductProperty struct for RmaReturnProductProperty
type RmaReturnProductProperty struct {
	Quantity             *string `json:"quantity,omitempty"`
	Note                 *string `json:"note,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RmaReturnProductProperty RmaReturnProductProperty

// NewRmaReturnProductProperty instantiates a new RmaReturnProductProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaReturnProductProperty() *RmaReturnProductProperty {
	this := RmaReturnProductProperty{}
	return &this
}

// NewRmaReturnProductPropertyWithDefaults instantiates a new RmaReturnProductProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaReturnProductPropertyWithDefaults() *RmaReturnProductProperty {
	this := RmaReturnProductProperty{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *RmaReturnProductProperty) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnProductProperty) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *RmaReturnProductProperty) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *RmaReturnProductProperty) SetQuantity(v string) {
	o.Quantity = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *RmaReturnProductProperty) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnProductProperty) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *RmaReturnProductProperty) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *RmaReturnProductProperty) SetNote(v string) {
	o.Note = &v
}

func (o RmaReturnProductProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaReturnProductProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaReturnProductProperty) UnmarshalJSON(data []byte) (err error) {
	varRmaReturnProductProperty := _RmaReturnProductProperty{}

	err = json.Unmarshal(data, &varRmaReturnProductProperty)

	if err != nil {
		return err
	}

	*o = RmaReturnProductProperty(varRmaReturnProductProperty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "note")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaReturnProductProperty) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *RmaReturnProductProperty) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableRmaReturnProductProperty struct {
	value *RmaReturnProductProperty
	isSet bool
}

func (v NullableRmaReturnProductProperty) Get() *RmaReturnProductProperty {
	return v.value
}

func (v *NullableRmaReturnProductProperty) Set(val *RmaReturnProductProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaReturnProductProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaReturnProductProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaReturnProductProperty(val *RmaReturnProductProperty) *NullableRmaReturnProductProperty {
	return &NullableRmaReturnProductProperty{value: val, isSet: true}
}

func (v NullableRmaReturnProductProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaReturnProductProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
