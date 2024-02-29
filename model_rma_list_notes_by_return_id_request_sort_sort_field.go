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

// RmaListNotesByReturnIdRequestSortSortField the model 'RmaListNotesByReturnIdRequestSortSortField'
type RmaListNotesByReturnIdRequestSortSortField string

// List of rmaListNotesByReturnIdRequestSortSortField
const (
	UNKNOWN RmaListNotesByReturnIdRequestSortSortField = "UNKNOWN"
	DATE RmaListNotesByReturnIdRequestSortSortField = "DATE"
)

// All allowed values of RmaListNotesByReturnIdRequestSortSortField enum
var AllowedRmaListNotesByReturnIdRequestSortSortFieldEnumValues = []RmaListNotesByReturnIdRequestSortSortField{
	"UNKNOWN",
	"DATE",
}

func (v *RmaListNotesByReturnIdRequestSortSortField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RmaListNotesByReturnIdRequestSortSortField(value)
	for _, existing := range AllowedRmaListNotesByReturnIdRequestSortSortFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RmaListNotesByReturnIdRequestSortSortField", value)
}

// NewRmaListNotesByReturnIdRequestSortSortFieldFromValue returns a pointer to a valid RmaListNotesByReturnIdRequestSortSortField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRmaListNotesByReturnIdRequestSortSortFieldFromValue(v string) (*RmaListNotesByReturnIdRequestSortSortField, error) {
	ev := RmaListNotesByReturnIdRequestSortSortField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RmaListNotesByReturnIdRequestSortSortField: valid values are %v", v, AllowedRmaListNotesByReturnIdRequestSortSortFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RmaListNotesByReturnIdRequestSortSortField) IsValid() bool {
	for _, existing := range AllowedRmaListNotesByReturnIdRequestSortSortFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to rmaListNotesByReturnIdRequestSortSortField value
func (v RmaListNotesByReturnIdRequestSortSortField) Ptr() *RmaListNotesByReturnIdRequestSortSortField {
	return &v
}

type NullableRmaListNotesByReturnIdRequestSortSortField struct {
	value *RmaListNotesByReturnIdRequestSortSortField
	isSet bool
}

func (v NullableRmaListNotesByReturnIdRequestSortSortField) Get() *RmaListNotesByReturnIdRequestSortSortField {
	return v.value
}

func (v *NullableRmaListNotesByReturnIdRequestSortSortField) Set(val *RmaListNotesByReturnIdRequestSortSortField) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaListNotesByReturnIdRequestSortSortField) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaListNotesByReturnIdRequestSortSortField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaListNotesByReturnIdRequestSortSortField(val *RmaListNotesByReturnIdRequestSortSortField) *NullableRmaListNotesByReturnIdRequestSortSortField {
	return &NullableRmaListNotesByReturnIdRequestSortSortField{value: val, isSet: true}
}

func (v NullableRmaListNotesByReturnIdRequestSortSortField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaListNotesByReturnIdRequestSortSortField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

