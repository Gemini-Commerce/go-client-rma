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

// checks if the RmaListReturnsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaListReturnsResponse{}

// RmaListReturnsResponse struct for RmaListReturnsResponse
type RmaListReturnsResponse struct {
	Returns              []RmaReturnResponse `json:"returns,omitempty"`
	NextPageToken        *string             `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RmaListReturnsResponse RmaListReturnsResponse

// NewRmaListReturnsResponse instantiates a new RmaListReturnsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaListReturnsResponse() *RmaListReturnsResponse {
	this := RmaListReturnsResponse{}
	return &this
}

// NewRmaListReturnsResponseWithDefaults instantiates a new RmaListReturnsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaListReturnsResponseWithDefaults() *RmaListReturnsResponse {
	this := RmaListReturnsResponse{}
	return &this
}

// GetReturns returns the Returns field value if set, zero value otherwise.
func (o *RmaListReturnsResponse) GetReturns() []RmaReturnResponse {
	if o == nil || IsNil(o.Returns) {
		var ret []RmaReturnResponse
		return ret
	}
	return o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListReturnsResponse) GetReturnsOk() ([]RmaReturnResponse, bool) {
	if o == nil || IsNil(o.Returns) {
		return nil, false
	}
	return o.Returns, true
}

// HasReturns returns a boolean if a field has been set.
func (o *RmaListReturnsResponse) HasReturns() bool {
	if o != nil && !IsNil(o.Returns) {
		return true
	}

	return false
}

// SetReturns gets a reference to the given []RmaReturnResponse and assigns it to the Returns field.
func (o *RmaListReturnsResponse) SetReturns(v []RmaReturnResponse) {
	o.Returns = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *RmaListReturnsResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListReturnsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *RmaListReturnsResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *RmaListReturnsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o RmaListReturnsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaListReturnsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Returns) {
		toSerialize["returns"] = o.Returns
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaListReturnsResponse) UnmarshalJSON(data []byte) (err error) {
	varRmaListReturnsResponse := _RmaListReturnsResponse{}

	err = json.Unmarshal(data, &varRmaListReturnsResponse)

	if err != nil {
		return err
	}

	*o = RmaListReturnsResponse(varRmaListReturnsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "returns")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaListReturnsResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *RmaListReturnsResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableRmaListReturnsResponse struct {
	value *RmaListReturnsResponse
	isSet bool
}

func (v NullableRmaListReturnsResponse) Get() *RmaListReturnsResponse {
	return v.value
}

func (v *NullableRmaListReturnsResponse) Set(val *RmaListReturnsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaListReturnsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaListReturnsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaListReturnsResponse(val *RmaListReturnsResponse) *NullableRmaListReturnsResponse {
	return &NullableRmaListReturnsResponse{value: val, isSet: true}
}

func (v NullableRmaListReturnsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaListReturnsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
