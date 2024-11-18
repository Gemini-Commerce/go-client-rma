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

// ProtobufNullValue `NullValue` is a singleton enumeration to represent the null value for the `Value` type union.   The JSON representation for `NullValue` is JSON `null`.   - NULL_VALUE: Null value.
type ProtobufNullValue string

// List of protobufNullValue
const (
	PROTOBUFNULLVALUE_NULL_VALUE ProtobufNullValue = "NULL_VALUE"
)

// All allowed values of ProtobufNullValue enum
var AllowedProtobufNullValueEnumValues = []ProtobufNullValue{
	"NULL_VALUE",
}

func (v *ProtobufNullValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProtobufNullValue(value)
	for _, existing := range AllowedProtobufNullValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProtobufNullValue", value)
}

// NewProtobufNullValueFromValue returns a pointer to a valid ProtobufNullValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProtobufNullValueFromValue(v string) (*ProtobufNullValue, error) {
	ev := ProtobufNullValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProtobufNullValue: valid values are %v", v, AllowedProtobufNullValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProtobufNullValue) IsValid() bool {
	for _, existing := range AllowedProtobufNullValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to protobufNullValue value
func (v ProtobufNullValue) Ptr() *ProtobufNullValue {
	return &v
}

type NullableProtobufNullValue struct {
	value *ProtobufNullValue
	isSet bool
}

func (v NullableProtobufNullValue) Get() *ProtobufNullValue {
	return v.value
}

func (v *NullableProtobufNullValue) Set(val *ProtobufNullValue) {
	v.value = val
	v.isSet = true
}

func (v NullableProtobufNullValue) IsSet() bool {
	return v.isSet
}

func (v *NullableProtobufNullValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtobufNullValue(val *ProtobufNullValue) *NullableProtobufNullValue {
	return &NullableProtobufNullValue{value: val, isSet: true}
}

func (v NullableProtobufNullValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtobufNullValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
