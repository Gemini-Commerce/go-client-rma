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

// RmaSortOrder the model 'RmaSortOrder'
type RmaSortOrder string

// List of rmaSortOrder
const (
	RMASORTORDER_DESC RmaSortOrder = "DESC"
	RMASORTORDER_ASC RmaSortOrder = "ASC"
)

// All allowed values of RmaSortOrder enum
var AllowedRmaSortOrderEnumValues = []RmaSortOrder{
	"DESC",
	"ASC",
}

func (v *RmaSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RmaSortOrder(value)
	for _, existing := range AllowedRmaSortOrderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RmaSortOrder", value)
}

// NewRmaSortOrderFromValue returns a pointer to a valid RmaSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRmaSortOrderFromValue(v string) (*RmaSortOrder, error) {
	ev := RmaSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RmaSortOrder: valid values are %v", v, AllowedRmaSortOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RmaSortOrder) IsValid() bool {
	for _, existing := range AllowedRmaSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to rmaSortOrder value
func (v RmaSortOrder) Ptr() *RmaSortOrder {
	return &v
}

type NullableRmaSortOrder struct {
	value *RmaSortOrder
	isSet bool
}

func (v NullableRmaSortOrder) Get() *RmaSortOrder {
	return v.value
}

func (v *NullableRmaSortOrder) Set(val *RmaSortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaSortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaSortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaSortOrder(val *RmaSortOrder) *NullableRmaSortOrder {
	return &NullableRmaSortOrder{value: val, isSet: true}
}

func (v NullableRmaSortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaSortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

