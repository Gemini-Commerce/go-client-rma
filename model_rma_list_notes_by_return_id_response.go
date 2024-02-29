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

// checks if the RmaListNotesByReturnIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaListNotesByReturnIdResponse{}

// RmaListNotesByReturnIdResponse struct for RmaListNotesByReturnIdResponse
type RmaListNotesByReturnIdResponse struct {
	Notes []RmaNoteResponse `json:"notes,omitempty"`
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// NewRmaListNotesByReturnIdResponse instantiates a new RmaListNotesByReturnIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaListNotesByReturnIdResponse() *RmaListNotesByReturnIdResponse {
	this := RmaListNotesByReturnIdResponse{}
	return &this
}

// NewRmaListNotesByReturnIdResponseWithDefaults instantiates a new RmaListNotesByReturnIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaListNotesByReturnIdResponseWithDefaults() *RmaListNotesByReturnIdResponse {
	this := RmaListNotesByReturnIdResponse{}
	return &this
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *RmaListNotesByReturnIdResponse) GetNotes() []RmaNoteResponse {
	if o == nil || IsNil(o.Notes) {
		var ret []RmaNoteResponse
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListNotesByReturnIdResponse) GetNotesOk() ([]RmaNoteResponse, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *RmaListNotesByReturnIdResponse) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []RmaNoteResponse and assigns it to the Notes field.
func (o *RmaListNotesByReturnIdResponse) SetNotes(v []RmaNoteResponse) {
	o.Notes = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *RmaListNotesByReturnIdResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListNotesByReturnIdResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *RmaListNotesByReturnIdResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *RmaListNotesByReturnIdResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o RmaListNotesByReturnIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaListNotesByReturnIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return toSerialize, nil
}

type NullableRmaListNotesByReturnIdResponse struct {
	value *RmaListNotesByReturnIdResponse
	isSet bool
}

func (v NullableRmaListNotesByReturnIdResponse) Get() *RmaListNotesByReturnIdResponse {
	return v.value
}

func (v *NullableRmaListNotesByReturnIdResponse) Set(val *RmaListNotesByReturnIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaListNotesByReturnIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaListNotesByReturnIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaListNotesByReturnIdResponse(val *RmaListNotesByReturnIdResponse) *NullableRmaListNotesByReturnIdResponse {
	return &NullableRmaListNotesByReturnIdResponse{value: val, isSet: true}
}

func (v NullableRmaListNotesByReturnIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaListNotesByReturnIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

