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

// checks if the RmaListNotesByReturnIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaListNotesByReturnIdRequest{}

// RmaListNotesByReturnIdRequest struct for RmaListNotesByReturnIdRequest
type RmaListNotesByReturnIdRequest struct {
	TenantId string `json:"tenantId"`
	ReturnId string `json:"returnId"`
	PageSize *int64 `json:"pageSize,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
	Sorts []RmaListNotesByReturnIdRequestSort `json:"sorts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RmaListNotesByReturnIdRequest RmaListNotesByReturnIdRequest

// NewRmaListNotesByReturnIdRequest instantiates a new RmaListNotesByReturnIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaListNotesByReturnIdRequest(tenantId string, returnId string) *RmaListNotesByReturnIdRequest {
	this := RmaListNotesByReturnIdRequest{}
	this.TenantId = tenantId
	this.ReturnId = returnId
	return &this
}

// NewRmaListNotesByReturnIdRequestWithDefaults instantiates a new RmaListNotesByReturnIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaListNotesByReturnIdRequestWithDefaults() *RmaListNotesByReturnIdRequest {
	this := RmaListNotesByReturnIdRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *RmaListNotesByReturnIdRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *RmaListNotesByReturnIdRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *RmaListNotesByReturnIdRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetReturnId returns the ReturnId field value
func (o *RmaListNotesByReturnIdRequest) GetReturnId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnId
}

// GetReturnIdOk returns a tuple with the ReturnId field value
// and a boolean to check if the value has been set.
func (o *RmaListNotesByReturnIdRequest) GetReturnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnId, true
}

// SetReturnId sets field value
func (o *RmaListNotesByReturnIdRequest) SetReturnId(v string) {
	o.ReturnId = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *RmaListNotesByReturnIdRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListNotesByReturnIdRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// &#39;Has&#39;PageSize returns a boolean if a field has been set.
func (o *RmaListNotesByReturnIdRequest) &#39;Has&#39;PageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *RmaListNotesByReturnIdRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *RmaListNotesByReturnIdRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListNotesByReturnIdRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// &#39;Has&#39;PageToken returns a boolean if a field has been set.
func (o *RmaListNotesByReturnIdRequest) &#39;Has&#39;PageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *RmaListNotesByReturnIdRequest) SetPageToken(v string) {
	o.PageToken = &v
}

// GetSorts returns the Sorts field value if set, zero value otherwise.
func (o *RmaListNotesByReturnIdRequest) GetSorts() []RmaListNotesByReturnIdRequestSort {
	if o == nil || IsNil(o.Sorts) {
		var ret []RmaListNotesByReturnIdRequestSort
		return ret
	}
	return o.Sorts
}

// GetSortsOk returns a tuple with the Sorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListNotesByReturnIdRequest) GetSortsOk() ([]RmaListNotesByReturnIdRequestSort, bool) {
	if o == nil || IsNil(o.Sorts) {
		return nil, false
	}
	return o.Sorts, true
}

// &#39;Has&#39;Sorts returns a boolean if a field has been set.
func (o *RmaListNotesByReturnIdRequest) &#39;Has&#39;Sorts() bool {
	if o != nil && !IsNil(o.Sorts) {
		return true
	}

	return false
}

// SetSorts gets a reference to the given []RmaListNotesByReturnIdRequestSort and assigns it to the Sorts field.
func (o *RmaListNotesByReturnIdRequest) SetSorts(v []RmaListNotesByReturnIdRequestSort) {
	o.Sorts = v
}

func (o RmaListNotesByReturnIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaListNotesByReturnIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["returnId"] = o.ReturnId
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	if !IsNil(o.Sorts) {
		toSerialize["sorts"] = o.Sorts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaListNotesByReturnIdRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"returnId",
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

	varRmaListNotesByReturnIdRequest := _RmaListNotesByReturnIdRequest{}

	err = json.Unmarshal(data, &varRmaListNotesByReturnIdRequest)

	if err != nil {
		return err
	}

	*o = RmaListNotesByReturnIdRequest(varRmaListNotesByReturnIdRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "returnId")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "pageToken")
		delete(additionalProperties, "sorts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaListNotesByReturnIdRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *RmaListNotesByReturnIdRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableRmaListNotesByReturnIdRequest struct {
	value *RmaListNotesByReturnIdRequest
	isSet bool
}

func (v NullableRmaListNotesByReturnIdRequest) Get() *RmaListNotesByReturnIdRequest {
	return v.value
}

func (v *NullableRmaListNotesByReturnIdRequest) Set(val *RmaListNotesByReturnIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaListNotesByReturnIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaListNotesByReturnIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaListNotesByReturnIdRequest(val *RmaListNotesByReturnIdRequest) *NullableRmaListNotesByReturnIdRequest {
	return &NullableRmaListNotesByReturnIdRequest{value: val, isSet: true}
}

func (v NullableRmaListNotesByReturnIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaListNotesByReturnIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


