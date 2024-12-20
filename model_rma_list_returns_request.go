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

// checks if the RmaListReturnsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaListReturnsRequest{}

// RmaListReturnsRequest struct for RmaListReturnsRequest
type RmaListReturnsRequest struct {
	TenantId             string                      `json:"tenantId"`
	PageSize             *int64                      `json:"pageSize,omitempty"`
	PageToken            *string                     `json:"pageToken,omitempty"`
	Sorts                []RmaListReturnsRequestSort `json:"sorts,omitempty"`
	FilterMask           *string                     `json:"filterMask,omitempty"`
	Filter               *ListReturnsRequestFilter   `json:"filter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RmaListReturnsRequest RmaListReturnsRequest

// NewRmaListReturnsRequest instantiates a new RmaListReturnsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaListReturnsRequest(tenantId string) *RmaListReturnsRequest {
	this := RmaListReturnsRequest{}
	this.TenantId = tenantId
	return &this
}

// NewRmaListReturnsRequestWithDefaults instantiates a new RmaListReturnsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaListReturnsRequestWithDefaults() *RmaListReturnsRequest {
	this := RmaListReturnsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *RmaListReturnsRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *RmaListReturnsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *RmaListReturnsRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *RmaListReturnsRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListReturnsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *RmaListReturnsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *RmaListReturnsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *RmaListReturnsRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListReturnsRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *RmaListReturnsRequest) HasPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *RmaListReturnsRequest) SetPageToken(v string) {
	o.PageToken = &v
}

// GetSorts returns the Sorts field value if set, zero value otherwise.
func (o *RmaListReturnsRequest) GetSorts() []RmaListReturnsRequestSort {
	if o == nil || IsNil(o.Sorts) {
		var ret []RmaListReturnsRequestSort
		return ret
	}
	return o.Sorts
}

// GetSortsOk returns a tuple with the Sorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListReturnsRequest) GetSortsOk() ([]RmaListReturnsRequestSort, bool) {
	if o == nil || IsNil(o.Sorts) {
		return nil, false
	}
	return o.Sorts, true
}

// HasSorts returns a boolean if a field has been set.
func (o *RmaListReturnsRequest) HasSorts() bool {
	if o != nil && !IsNil(o.Sorts) {
		return true
	}

	return false
}

// SetSorts gets a reference to the given []RmaListReturnsRequestSort and assigns it to the Sorts field.
func (o *RmaListReturnsRequest) SetSorts(v []RmaListReturnsRequestSort) {
	o.Sorts = v
}

// GetFilterMask returns the FilterMask field value if set, zero value otherwise.
func (o *RmaListReturnsRequest) GetFilterMask() string {
	if o == nil || IsNil(o.FilterMask) {
		var ret string
		return ret
	}
	return *o.FilterMask
}

// GetFilterMaskOk returns a tuple with the FilterMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListReturnsRequest) GetFilterMaskOk() (*string, bool) {
	if o == nil || IsNil(o.FilterMask) {
		return nil, false
	}
	return o.FilterMask, true
}

// HasFilterMask returns a boolean if a field has been set.
func (o *RmaListReturnsRequest) HasFilterMask() bool {
	if o != nil && !IsNil(o.FilterMask) {
		return true
	}

	return false
}

// SetFilterMask gets a reference to the given string and assigns it to the FilterMask field.
func (o *RmaListReturnsRequest) SetFilterMask(v string) {
	o.FilterMask = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RmaListReturnsRequest) GetFilter() ListReturnsRequestFilter {
	if o == nil || IsNil(o.Filter) {
		var ret ListReturnsRequestFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListReturnsRequest) GetFilterOk() (*ListReturnsRequestFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RmaListReturnsRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given ListReturnsRequestFilter and assigns it to the Filter field.
func (o *RmaListReturnsRequest) SetFilter(v ListReturnsRequestFilter) {
	o.Filter = &v
}

func (o RmaListReturnsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaListReturnsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	if !IsNil(o.Sorts) {
		toSerialize["sorts"] = o.Sorts
	}
	if !IsNil(o.FilterMask) {
		toSerialize["filterMask"] = o.FilterMask
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaListReturnsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
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

	varRmaListReturnsRequest := _RmaListReturnsRequest{}

	err = json.Unmarshal(data, &varRmaListReturnsRequest)

	if err != nil {
		return err
	}

	*o = RmaListReturnsRequest(varRmaListReturnsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "pageToken")
		delete(additionalProperties, "sorts")
		delete(additionalProperties, "filterMask")
		delete(additionalProperties, "filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaListReturnsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *RmaListReturnsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableRmaListReturnsRequest struct {
	value *RmaListReturnsRequest
	isSet bool
}

func (v NullableRmaListReturnsRequest) Get() *RmaListReturnsRequest {
	return v.value
}

func (v *NullableRmaListReturnsRequest) Set(val *RmaListReturnsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaListReturnsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaListReturnsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaListReturnsRequest(val *RmaListReturnsRequest) *NullableRmaListReturnsRequest {
	return &NullableRmaListReturnsRequest{value: val, isSet: true}
}

func (v NullableRmaListReturnsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaListReturnsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
